package strs

import (
	"regexp"
)

func GetLinesByRegexSubstringMatches(
	lines *[]string,
	regexp *regexp.Regexp,
) *map[int]*[]string {
	if lines == nil {
		return &map[int]*[]string{}
	}

	length := len(*lines)

	if length == 0 {
		return &map[int]*[]string{}
	}

	mappedItems := make(map[int]*[]string, length)

	for index, line := range *lines {
		matches := regexp.FindStringSubmatch(line)

		if matches != nil && len(matches) > 0 {
			mappedItems[index] = &matches
		}
	}

	return &mappedItems
}
