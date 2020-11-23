package strhelper

import (
	"regexp"

	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

// NewRegExResultWrapper returns new regex wrappers for all the regex given
func NewRegExResultWrapper(content *string, regexps ...*regexp.Regexp) []*strhelpercore.RegExResultWrapper {
	results := make([]*strhelpercore.RegExResultWrapper, 0, len(regexps))

	for index, regex := range regexps {
		wrapper := strhelpercore.NewRegExResultWrapper(index, content, regex)
		results = append(results, wrapper)
	}

	return results
}
