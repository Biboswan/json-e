package prattparser

import (
	"fmt"
	"regexp"
)

func mustBeNonCaptureRegex(r string) {
	if regexp.MustCompile(r).NumSubexp() > 0 {
		panic(fmt.Sprintf("regular expression '%s' should be non-capturing", r))
	}
}

func stringsContains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
