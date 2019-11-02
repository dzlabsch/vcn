/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package login

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/vchain-us/vcn/pkg/meta"
	"github.com/vchain-us/vcn/pkg/store"
)

const testUserEnv = "VCN_TEST_USER"
const testPassEnv = "VCN_TEST_PASS"

const testPassphrase = "dummy"

// leogr (todo): decouple integration tests
func TestLoginByEnv(t *testing.T) {

	user := os.Getenv(testUserEnv)
	password := os.Getenv(testPassEnv)

	if user == "" || password == "" {
		t.Skip(
			fmt.Sprintf(
				"Please set %s and %s environment variables to run this test.",
				testUserEnv,
				testPassEnv,
			),
		)
		return
	}

	tdir, err := ioutil.TempDir("", "vcn-testing")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tdir) // clean up

	// Setup temporary env
	store.SetDir(tdir)

	os.Setenv(meta.VcnUserEnv, user)
	os.Setenv(meta.VcnPasswordEnv, password)

	Execute()
}
