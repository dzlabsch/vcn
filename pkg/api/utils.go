/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package api

import (
	"fmt"

	"github.com/dghubble/sling"
	"github.com/sirupsen/logrus"
	"github.com/vchain-us/vcn/internal/errors"
	"github.com/vchain-us/vcn/internal/logs"
	"github.com/vchain-us/vcn/pkg/meta"
)

func logger() *logrus.Logger {
	return logs.LOG
}

func makeError(msg string, fields logrus.Fields) error {
	err := fmt.Errorf(msg)
	logger().WithFields(fields).Error(err)
	return err
}

func makeFatal(msg string, fields logrus.Fields) error {
	err := fmt.Errorf(msg)
	logger().WithFields(fields).Fatal(err)
	return err
}

func makeAuthRequiredError() error {
	return makeError(errors.AuthRequired, nil)
}

func contains(xs []string, x string) bool {
	for _, a := range xs {
		if a == x {
			return true
		}
	}
	return false
}

func newSling(token string) (s *sling.Sling) {
	s = sling.New()
	s.Add("User-Agent", meta.UserAgent())
	if token != "" {
		s = s.Add("Authorization", "Bearer "+token)
	}
	return s
}
