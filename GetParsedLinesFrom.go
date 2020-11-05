package strhelper

import (
	"regexp"

	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

// Gets regex wrapper for all the regex given
func GetParsedLinesFrom(content *string, regexps ...*regexp.Regexp) []*strhelpercore.RegExResultWrapper {
	results := make([]*strhelpercore.RegExResultWrapper, 0, len(regexps))

	for index, regex := range regexps {
		wrapper := strhelpercore.NewRegExResultWrapper(index, content, regex)
		results = append(results, &wrapper)
	}

	return results
}
