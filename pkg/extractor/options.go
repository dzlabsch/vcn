/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package extractor

// Option is a functional option for extractors.
type Option func(interface{}) error

// Options is a slice of Option.
type Options []Option

// Apply interates over Options and calls each functional option with a given opts.
func (o Options) Apply(opts interface{}) error {
	for _, f := range o {
		if err := f(opts); err != nil {
			return err
		}
	}

	return nil
}
