/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package logout

import (
	"fmt"

	"github.com/vchain-us/vcn/pkg/store"

	"github.com/spf13/cobra"
)

// NewCommand returns the cobra command for `vcn logout`
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout the current user",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			output, err := cmd.Flags().GetString("output")
			if err != nil {
				return err
			}

			if err := Execute(); err != nil {
				return err
			}
			if output == "" {
				fmt.Println("Logout successful.")
			}
			return nil
		},
		Args: cobra.NoArgs,
	}

	return cmd
}

// Execute logout action
func Execute() error {
	store.Config().ClearContext()
	if err := store.SaveConfig(); err != nil {
		return err
	}
	return nil
}
