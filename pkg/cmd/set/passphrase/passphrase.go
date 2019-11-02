/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package passphrase

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/spf13/cobra"

	"github.com/vchain-us/vcn/internal/assert"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/cmd/internal/cli"
	"github.com/vchain-us/vcn/pkg/store"
)

// NewCommand returns the cobra command for `vcn set notarization-password`
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "notarization-password",
		Aliases: []string{"passphrase"},
		Short:   "Change the notarization password for the current user",
		Long:    `This command allows you to set a custom notarization password.`,
		RunE:    runPasshphrase,
		Args:    cobra.NoArgs,
	}

	return cmd
}

func runPasshphrase(cmd *cobra.Command, args []string) error {

	cmd.SilenceUsage = true

	// User
	if err := assert.UserLogin(); err != nil {
		return err
	}
	u := api.NewUser(store.Config().CurrentContext)
	fmt.Printf("User:	%s\n", u.Email())

	secret, id, offline, err := u.Secret()
	if err != nil {
		return err
	}
	if offline {
		return fmt.Errorf("offline secret is not supported by the current vcn version")
	}
	fmt.Printf("SignerID:	%s\n", id)

	pass, err := cli.ProvidePasswordWithMessage("Enter your current notarization password: ")
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(secret)
	key, err := keystore.DecryptKey(buf.Bytes(), pass)
	if err != nil {
		if err.Error() == "could not decrypt key with given passphrase" {
			return fmt.Errorf("incorrect password")
		}
		return err
	}

	pass, err = cli.PromptPassphrase()
	if err != nil {
		return err
	}

	b, err := keystore.EncryptKey(key, pass, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return err
	}

	err = u.UploadSecret(bytes.NewReader(b), pass)
	if err != nil {
		return err
	}
	fmt.Println("You have successfully updated the password.")
	return nil
}
