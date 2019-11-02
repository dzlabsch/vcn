/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package dir

import (
	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/bundle"
)

const (
	manifestFile = ".vcn.manifest.json"
)

// Metadata extracts dir related info from a.
func Metadata(a api.Artifact) (manifest *bundle.Manifest, path string) {
	if a.Kind != Scheme {
		return
	}

	// Get manifest
	m := a.Metadata[ManifestKey]
	if m != nil {
		if mm, ok := m.(*bundle.Manifest); ok {
			manifest = mm
		}
	}

	// Get path
	p := a.Metadata[PathKey]
	if p != nil {
		if pp, ok := p.(string); ok {
			path = pp
		}
	}

	return
}

// RemoveMetadata removes dir related info from a.
func RemoveMetadata(a *api.Artifact) {
	if a == nil || a.Kind != Scheme {
		return
	}
	delete(a.Metadata, ManifestKey)
	delete(a.Metadata, PathKey)
}
