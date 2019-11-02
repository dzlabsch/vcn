/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package types

type Error struct {
	Error string `json:"error"`
}

func (e *Error) String() string {
	if e != nil {
		return e.Error
	}
	return ""
}

func NewError(err error) *Error {
	return &Error{
		Error: err.Error(),
	}
}
