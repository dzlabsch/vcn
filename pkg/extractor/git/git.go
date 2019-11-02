/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package git

import (
	"path/filepath"
	"strings"

	git "gopkg.in/src-d/go-git.v4"

	"github.com/vchain-us/vcn/pkg/api"
	"github.com/vchain-us/vcn/pkg/extractor"
	"github.com/vchain-us/vcn/pkg/uri"
)

// Scheme for git
const Scheme = "git"

// Artifact returns a git *api.Artifact from a given u
func Artifact(u *uri.URI, options ...extractor.Option) (*api.Artifact, error) {

	if u.Scheme != Scheme {
		return nil, nil
	}

	path := strings.TrimPrefix(u.Opaque, "//")
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	commit, err := lastCommit(repo)
	if err != nil {
		return nil, err
	}

	hash, size, err := digestCommit(*commit)
	if err != nil {
		return nil, err
	}

	// Metadata container
	m := api.Metadata{
		Scheme: map[string]interface{}{
			"Commit": commit.Hash.String(),
			"Tree":   commit.TreeHash.String(),
			"Parents": func() []string {
				res := make([]string, len(commit.ParentHashes))
				for i, h := range commit.ParentHashes {
					res[i] = h.String()
				}
				return res
			}(),
			"Author":       commit.Author,
			"Committer":    commit.Committer,
			"Message":      commit.Message,
			"PGPSignature": commit.PGPSignature,
		},
	}

	name := filepath.Base(path)
	if remotes, err := repo.Remotes(); err == nil && len(remotes) > 0 {
		urls := remotes[0].Config().URLs
		if len(urls) > 0 {
			name = urls[0]
		}
	}
	name += "@" + commit.Hash.String()[:7]

	return &api.Artifact{
		Kind:     Scheme,
		Hash:     hash,
		Size:     size,
		Name:     name,
		Metadata: m,
	}, nil
}
