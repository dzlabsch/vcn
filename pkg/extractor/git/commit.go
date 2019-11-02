/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package git

import (
	"crypto/sha256"
	"encoding/hex"
	"io"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func lastCommit(repo *git.Repository) (*object.Commit, error) {
	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	return repo.CommitObject(ref.Hash())
}

func digestCommit(c object.Commit) (hash string, size uint64, err error) {
	o := &plumbing.MemoryObject{}
	c.Encode(o)

	reader, err := o.Reader()
	if err != nil {
		return
	}
	defer reader.Close()

	h := sha256.New()
	n, err := io.Copy(h, reader)
	if err != nil {
		return
	}
	return hex.EncodeToString(h.Sum(nil)), uint64(n), nil
}
