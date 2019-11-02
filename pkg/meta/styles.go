/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package meta

import (
	"github.com/fatih/color"
)

// StatusColor returns color.Attribute(s) for the given status
func StatusColor(status Status) (color.Attribute, color.Attribute) {
	switch status {
	case StatusTrusted:
		return StyleSuccess()
	case StatusUnknown:
		return StyleWarning()
	default:
		return StyleError()
	}
}

// StyleAffordance returns color.Attribute(s) for affordance messages
func StyleAffordance() (color.Attribute, color.Attribute) {
	return color.FgHiBlue, color.Bold
}

// StyleError returns color.Attribute(s) for error messages
func StyleError() (color.Attribute, color.Attribute) {
	return color.FgRed, color.Bold
}

// StyleWarning returns color.Attribute(s) for warning messages
func StyleWarning() (color.Attribute, color.Attribute) {
	return color.FgYellow, color.Bold
}

// StyleSuccess returns color.Attribute(s) for success messages
func StyleSuccess() (color.Attribute, color.Attribute) {
	return color.FgGreen, color.Bold
}
