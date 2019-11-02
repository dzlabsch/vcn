/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package store

import (
	"io/ioutil"
	"path/filepath"
	"time"
)

const verCheckFilename = ".versioncheck"

// VersionCheckTime returns the time of latest version check, if any.
func VersionCheckTime() *time.Time {
	content, err := ioutil.ReadFile(filepath.Join(dir, verCheckFilename))
	if err != nil {
		return nil
	}
	t, err := time.Parse(time.RFC3339, string(content))
	if err != nil {
		return nil
	}
	return &t
}

// SetVersionCheckTime set the latest version check to now.
func SetVersionCheckTime() {
	if err := ensureDir(dir); err != nil {
		return
	}
	t := time.Now().Format(time.RFC3339)
	ioutil.WriteFile(filepath.Join(dir, verCheckFilename), []byte(t), FilePerm)
}
