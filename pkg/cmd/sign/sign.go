/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package sign

import (
	"fmt"
	"io"
	"strings"

	"github.com/vchain-us/vcn/pkg/extractor/dir"

	"github.com/fatih/color"

	"github.com/caarlos0/spin"
	"github.com/spf13/cobra"
	"github.com/vchain-us/vcn/internal/assert"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/cmd/internal/cli"
	"github.com/vchain-us/vcn/pkg/cmd/internal/types"
	"github.com/vchain-us/vcn/pkg/extractor"
	"github.com/vchain-us/vcn/pkg/meta"
	"github.com/vchain-us/vcn/pkg/store"
)

const longDescFooter = `

VCN_NOTARIZATION_PASSWORD env var can be used to pass the
required notarization password in a non-interactive environment.
`

const helpMsgFooter = `
ARG must be one of:
  <file>
  file://<file>
  dir://<directory>
  git://<repository>
  docker://<image>
  podman://<image>
`

// NewCommand returns the cobra command for `vcn sign`
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "notarize",
		Aliases: []string{"n", "sign", "s"},
		Short:   "Notarize an asset onto the blockchain",
		Long: `
Notarize an asset onto the blockchain.

Notarization calculates the SHA-256 hash of a digital asset 
(file, directory, container's image). 
The hash (not the asset) and the desired status of TRUSTED are then 
cryptographically signed by the signer's secret (private key). 
Next, these signed objects are sent to the blockchain where the signer’s
trust level and a timestamp are added. 
When complete, a new blockchain entry is created that binds the asset’s
signed hash, signed status, level, and timestamp together. 

Note that your asset will not be uploaded but processed locally.

Assets are referenced by passed ARG with notarization only accepting 
1 ARG at a time.
` + helpMsgFooter,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSignWithState(cmd, args, meta.StatusTrusted)
		},
		Args: noArgsWhenHash,
	}

	cmd.Flags().VarP(make(mapOpts), "attr", "a", "add user defined attributes (repeat --attr for multiple entries)")
	cmd.Flags().StringP("name", "n", "", "set the asset name")
	cmd.Flags().BoolP("public", "p", false, "when notarized as public, the asset name and metadata will be visible to everyone")
	cmd.Flags().String("hash", "", "specify the hash instead of using an asset, if set no ARG(s) can be used")
	cmd.Flags().Bool("no-ignore-file", false, "if set, .vcnignore will be not written inside the targeted dir")
	cmd.SetUsageTemplate(
		strings.Replace(cmd.UsageTemplate(), "{{.UseLine}}", "{{.UseLine}} ARG", 1),
	)

	return cmd
}

func runSignWithState(cmd *cobra.Command, args []string, state meta.Status) error {

	// default extractors options
	extractorOptions := []extractor.Option{}

	noIgnoreFile, err := cmd.Flags().GetBool("no-ignore-file")
	if err != nil {
		return err
	}
	if !noIgnoreFile {
		extractorOptions = append(extractorOptions, dir.WithIgnoreFileInit())
	}

	var hash string
	if hashFlag := cmd.Flags().Lookup("hash"); hashFlag != nil {
		var err error
		hash, err = cmd.Flags().GetString("hash")
		if err != nil {
			return err
		}
	}

	public, err := cmd.Flags().GetBool("public")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	metadata := cmd.Flags().Lookup("attr").Value.(mapOpts).StringToInterface()

	cmd.SilenceUsage = true

	// User
	if err := assert.UserLogin(); err != nil {
		return err
	}
	u := api.NewUser(store.Config().CurrentContext)
	if u == nil {
		return fmt.Errorf("cannot load the current user")
	}

	// Make the artifact to be signed
	var a *api.Artifact
	if hash != "" {
		hash = strings.ToLower(hash)
		// Load existing artifact, if any, otherwise use an empty artifact
		if ar, err := u.LoadArtifact(hash); err == nil && ar != nil {
			a = ar.Artifact()
		} else {
			if name == "" {
				return fmt.Errorf("please set an asset name, by using --name")
			}
			a = &api.Artifact{Hash: hash}
		}
	} else {
		// Extract artifact from arg
		a, err = extractor.Extract(args[0], extractorOptions...)
		if err != nil {
			return err
		}
	}

	if a == nil {
		return fmt.Errorf("unable to process the input asset provided")
	}

	// Override the asset's name, if provided by --name
	if name != "" {
		a.Name = name
	}

	// Copy user provided custom attributes
	a.Metadata.SetValues(metadata)

	return sign(*u, *a, state, meta.VisibilityForFlag(public), output)
}

func sign(u api.User, a api.Artifact, state meta.Status, visibility meta.Visibility, output string) error {

	if output == "" {
		color.Set(meta.StyleAffordance())
		fmt.Println("Your asset will not be uploaded but processed locally.")
		color.Unset()
		fmt.Println()
		fmt.Println("Signer:\t" + u.Email())
	}

	hook := newHook(&a)

	s := spin.New("%s Notarization in progress...")
	s.Set(spin.Spin1)

	var verification *api.BlockchainVerification
	var err error

	for i := 1; true; i++ {
		var passphrase string
		var interactive bool
		passphrase, interactive, err = cli.ProvidePassphrase()
		if err != nil {
			return err
		}

		if output == "" {
			s.Start()
		}

		var keyin io.Reader
		var offline bool
		keyin, _, offline, err = u.Secret()
		if err != nil {
			return err
		}
		if offline {
			return fmt.Errorf("offline secret is not supported by the current vcn version")
		}

		verification, err = u.Sign(
			a,
			api.SignWithStatus(state),
			api.SignWithVisibility(visibility),
			api.SignWithKey(keyin, passphrase),
		)

		if err != nil && i >= 3 {
			s.Stop()
			return fmt.Errorf("too many failed attempts: %s", err)
		}

		if interactive && err == api.WrongPassphraseErr {
			s.Stop()
			fmt.Printf("\nError: %s, please try again\n\n", err.Error())
			continue
		}
		break
	}

	// todo(ameingast/leogr): remove reduntat event - need backend improvement
	api.TrackPublisher(&u, meta.VcnSignEvent)
	api.TrackSign(&u, a.Hash, a.Name, state)

	if output == "" {
		s.Stop()
	}
	if err != nil {
		return err
	}

	err = hook.finalize(verification)
	if err != nil {
		return err
	}

	if output == "" {
		fmt.Println()
	}

	artifact, err := api.LoadArtifact(&u, a.Hash, verification.MetaHash())
	if err != nil {
		return err
	}

	cli.Print(output, types.NewResult(&a, artifact, verification))
	return nil
}
