/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package set

import (
	"github.com/vchain-us/vcn/pkg/cmd/set/passphrase"

	"github.com/spf13/cobra"
)

// NewCommand returns the cobra command for `vcn set`
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set specific features options",
		Long:  ``,
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(passphrase.NewCommand())
	// todo(leogr): re-enable when offline secret support is ready
	// cmd.AddCommand(secret.NewCommand())

	return cmd
}
