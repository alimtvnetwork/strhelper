package stringsearch

import (
	"regexp"

	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/defaultcapacity"
)

// GetElementsByRegExMatch returns the lines which matches with the regex given
func GetElementsByRegExMatch(
	contentsLines []string,
	regexp *regexp.Regexp,
) []string {
	if len(contentsLines) == 0 {
		return []string{}
	}

	slice := stringslice.MakeDefault(
		defaultcapacity.OfSearch(len(contentsLines)))

	for _, currentLine := range contentsLines {
		if regexp.MatchString(currentLine) {
			slice = append(slice, currentLine)
		}
	}

	return slice
}
