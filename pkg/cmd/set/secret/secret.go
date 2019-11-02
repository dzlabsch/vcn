/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package secret

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/cmd/internal/cli"
	"github.com/vchain-us/vcn/pkg/mnemonic"
	"github.com/vchain-us/vcn/pkg/store"
)

// NewCommand returns the cobra command for `vcn set secret`
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secret",
		Short: "Import the user's secret",
		Long: `Import a secret for the current user from a given mnemonic code and securely store it into the local vcn installation,
if successful, any pre-stored secret will be overwritten.
This feature is for advanced user only.
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			fmt.Println("Please, provide your mnemonic code in order to recover your secret.")
			return Execute()
		},
		Args: cobra.NoArgs,
	}

	return cmd
}

// Execute recover secret action
func Execute() error {

	u := api.NewUser(store.Config().CurrentContext)
	hasAuth, err := u.IsAuthenticated()
	if err != nil {
		return err
	}
	if !hasAuth {
		return fmt.Errorf("you need to be logged in, please use <vcn login>")
	}

	userCfg := u.Config()
	code, err := cli.PromptMnemonic()
	if err != nil {
		return err
	}

	privKey, err := mnemonic.ToECDSA(code)
	if err != nil {
		return err
	}

	passphrase, err := cli.PromptPassphrase()
	if err != nil {
		return err
	}

	if err := userCfg.ImportSecret(*privKey, passphrase); err != nil {
		return err
	}

	fmt.Println("Secret successfully imported.")
	fmt.Println("Secret Storage:\t", userCfg.KeyStore)
	fmt.Println("SignerID:\t", userCfg.SignerIDFromSecret())
	return nil
}
