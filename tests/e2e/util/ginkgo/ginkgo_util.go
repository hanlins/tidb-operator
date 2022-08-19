// Copyright 2021 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package ginkgo

import (
	"regexp"
	"strings"

	ginkgo "github.com/onsi/ginkgo/v2"
)

// ItWhenFocus run It only when It is focused
func ItWhenFocus(text string, body interface{}, timeout ...float64) bool {
	skip := true

	ginkgoconfig, _ := ginkgo.GinkgoConfiguration()
	focusString := strings.Join(ginkgoconfig.FocusStrings, "|")
	filter := regexp.MustCompile(focusString)

	if focusString != "" && filter.MatchString(text) {
		skip = false
	}

	if skip {
		return ginkgo.PIt(text, body)
	}
	return ginkgo.It(text, body)
}

// ContextWhenFocus run Context only when Context is focused
func ContextWhenFocus(text string, body func()) bool {
	skip := true

	ginkgoconfig, _ := ginkgo.GinkgoConfiguration()
	focusString := strings.Join(ginkgoconfig.FocusStrings, "|")
	filter := regexp.MustCompile(focusString)

	if focusString != "" && filter.MatchString(text) {
		skip = false
	}

	if skip {
		return ginkgo.PContext(text, body)
	}
	return ginkgo.Context(text, body)
}
