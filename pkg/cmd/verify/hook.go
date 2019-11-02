/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package verify

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/vchain-us/vcn/pkg/bundle"

	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/extractor/dir"
)

type hook struct {
	a       api.Artifact
	rawDiff bool
}

func newHook(cmd *cobra.Command, a *api.Artifact) *hook {
	if a != nil {
		h := hook{
			a: a.Copy(),
		}
		h.rawDiff, _ = cmd.Flags().GetBool("raw-diff")
		dir.RemoveMetadata(a)
		return &h
	}
	return nil
}

func (h *hook) finalize(v *api.BlockchainVerification, output string) error {
	if h != nil && output == "" {
		manifest, path := dir.Metadata(h.a)
		if manifest != nil && path != "" {
			oldManifest, err := bundle.ReadManifest(filepath.Join(path, bundle.ManifestFilename))
			if err != nil {
				fmt.Printf("Diff is unavailable because '%s' is missing or invalid.\n\n", bundle.ManifestFilename)
				return nil // ignore missing or bad manifest
			}
			// check old manifest integrity
			oldDigest, err := oldManifest.Digest()
			if err != nil {
				fmt.Printf("Diff is unavailable because '%s' is invalid.\n\n", bundle.ManifestFilename)
				return nil // ignore bad manifest
			}
			v, err := api.Verify(oldDigest.Encoded())
			if err != nil {
				return err
			}
			if v != nil && !v.Unknown() {
				var report string
				var equal bool
				var err error
				if h.rawDiff {
					report, equal, err = manifest.Diff(*oldManifest)
				} else {
					report, equal, err = manifest.DiffByPath(*oldManifest)
				}
				if err != nil {
					return err
				}
				if !equal {
					fmt.Printf("Diff since %s\n\n%s\n\n", v.Date(), report)
				}
			} else {
				fmt.Printf("Diff is unavailable because '%s' has been tampered.\n\n", bundle.ManifestFilename)
			}
		}
	}
	return nil
}
