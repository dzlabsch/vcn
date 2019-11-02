/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package store

import (
	"os"
)

// FilePerm holds permission bits that are used for all files that store creates.
const FilePerm os.FileMode = 0600

// DirPerm holds permission bits that are used for all directories that store creates.
const DirPerm os.FileMode = 0700

// DefaultDirName is the name of the store working directory.
const DefaultDirName = ".vcn"

const configFilename = "config.json"

const defaultSecretFile = "secret.json"
