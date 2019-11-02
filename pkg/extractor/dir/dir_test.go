/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package dir

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vchain-us/vcn/pkg/uri"
)

func TestArtifact(t *testing.T) {

	tmpDir, err := ioutil.TempDir("", "TempDir")
	if err != nil {
		t.Fatal(err)
	}

	tmpFile := filepath.Join(tmpDir, "file")
	err = ioutil.WriteFile(tmpFile, nil, 0644)
	if err != nil {
		t.Fatal(err)
	}
	// dir - OK
	u, _ := uri.Parse("dir://" + tmpDir)
	a, err := Artifact(u)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, "dir", a.Kind)
	assert.Equal(t, filepath.Base(tmpDir), a.Name)
	assert.NotEmpty(t, a.Hash)

	// dir (no schema) - OK
	u, _ = uri.Parse(tmpDir)
	a, err = Artifact(u)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, "dir", a.Kind)
	assert.Equal(t, filepath.Base(tmpDir), a.Name)
	assert.NotEmpty(t, a.Hash)

	// wrong schema - SKIP (no error)
	u, _ = uri.Parse("file://" + tmpDir)
	a, err = Artifact(u)
	assert.NoError(t, err)
	assert.Nil(t, a)

	// not a dir - ERROR
	u, _ = uri.Parse("dir://" + tmpFile)
	a, err = Artifact(u)
	assert.Error(t, err)
	assert.Nil(t, a)

	// not existing dir - ERROR
	u, _ = uri.Parse("dir://" + tmpDir + "/not-existing")
	a, err = Artifact(u)
	assert.Error(t, err)
	assert.Nil(t, a)
}
