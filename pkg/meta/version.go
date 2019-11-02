/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package meta

var version = ""

var gitCommit = ""
var gitBranch = ""

// Version returns the current CodeNotary vcn version string
func Version() string {
	return version
}

// GitRevision returns the current CodeNotary vcn git revision string
func GitRevision() string {
	rev := gitCommit
	if gitBranch != "" {
		rev += " (" + gitBranch + ")"
	}
	return rev
}
