/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package bundle

import (
	"io"
	"sort"

	// See https://github.com/opencontainers/go-digest#usage
	_ "crypto/sha256"
	_ "crypto/sha512"

	digest "github.com/opencontainers/go-digest"
)

// Descriptor describes the disposition of targeted content.
type Descriptor struct {
	// Digest is the digest of the targeted content.
	Digest digest.Digest `json:"digest"`

	// Size specifies the size in bytes of the targeted content.
	Size uint64 `json:"size"`

	// Paths specifies the relative locations of the targeted content.
	Paths []string `json:"paths"`
}

func (d *Descriptor) sortUnique() {
	tmp := make(map[string]bool, len(d.Paths))
	for _, p := range d.Paths {
		tmp[p] = true
	}
	d.Paths = make([]string, len(tmp))
	i := 0
	for p := range tmp {
		d.Paths[i] = p
		i++
	}
	sort.Strings(d.Paths)
}

// NewDescriptor returns a new *Descriptor for the provided path and src.
func NewDescriptor(path string, src io.Reader) (*Descriptor, error) {
	digester := digest.SHA256.Digester()
	size, err := io.Copy(digester.Hash(), src)
	if err != nil {
		return nil, err
	}

	return &Descriptor{
		Paths:  []string{path},
		Digest: digester.Digest(),
		Size:   uint64(size),
	}, nil
}
