package strs

import (
	"regexp"

	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/defaultcapacity"
)

func GetLinesByRegexMatches(
	lines *[]string,
	regexp *regexp.Regexp,
) *[]string {
	if lines == nil {
		return core.EmptyStringsPtr()
	}

	length := len(*lines)

	if length == 0 {
		return core.EmptyStringsPtr()
	}

	capacity := defaultcapacity.OfSearch(length)
	list := make([]string, 0, capacity)

	for _, line := range *lines {
		if regexp.MatchString(line) {
			list = append(list, line)
		}
	}

	return &list
}
