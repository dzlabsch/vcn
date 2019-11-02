/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package sign

import (
	"github.com/spf13/cobra"
	"github.com/vchain-us/vcn/pkg/meta"
)

// NewUnsupportCommand returns the cobra command for `vcn unsupport`
func NewUnsupportCommand() *cobra.Command {
	cmd := NewCommand()
	cmd.Use = "unsupport"
	cmd.Aliases = []string{"us"}
	cmd.Short = "Unsupport an asset"
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return runSignWithState(cmd, args, meta.StatusUnsupported)
	}
	cmd.Long = `
Change an asset's status so it is equal to UNSUPPORTED.

Unsupport command calculates the SHA-256 hash of a digital asset 
(file, directory, container's image). 
The hash (not the asset) and the desired status of UNSUPPORTED are then 
cryptographically signed by the signer's secret (private key). 
Next, these signed objects are sent to the blockchain where the signer’s
trust level and a timestamp are added. 
When complete, a new blockchain entry is created that binds the asset’s
signed hash, signed status, level, and timestamp together. 

Note that your asset will not be uploaded but processed locally.

Assets are referenced by passed ARG(s) with unsupport command only accepting 
1 ARG at a time. 

` + helpMsgFooter
	return cmd
}
