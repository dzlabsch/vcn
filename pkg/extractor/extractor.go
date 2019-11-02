/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package extractor

import (
	"fmt"

	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/uri"
)

var extractors = map[string]Extractor{}

// Extractor extract an api.Artifact referenced by the given uri.URI.
type Extractor func(*uri.URI, ...Option) (*api.Artifact, error)

// Register the Extractor e for the given scheme
func Register(scheme string, e Extractor) {
	extractors[scheme] = e
}

// Schemes returns the list of registered schemes.
func Schemes() []string {
	schemes := make([]string, len(extractors))
	i := 0
	for scheme := range extractors {
		schemes[i] = scheme
		i++
	}
	return schemes
}

// Extract returns an api.Artifact for the given rawURI.
func Extract(rawURI string, options ...Option) (*api.Artifact, error) {
	u, err := uri.Parse(rawURI)
	if err != nil {
		return nil, err
	}

	if e, ok := extractors[u.Scheme]; ok {
		return e(u, options...)
	}
	return nil, fmt.Errorf("%s scheme not yet supported", u.Scheme)
}
