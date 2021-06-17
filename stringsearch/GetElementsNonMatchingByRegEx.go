package stringsearch

import (
	"regexp"

	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/core/defaultcapacity"
)

// GetElementsNonMatchingByRegEx returns the lines which doesn't meet with regex requirements
func GetElementsNonMatchingByRegEx(
	contentsLines []string,
	regexp *regexp.Regexp,
) []string {
	if len(contentsLines) == 0 {
		return []string{}
	}

	slice := stringslice.MakeDefault(
		defaultcapacity.OfSearch(len(contentsLines)))

	for _, currentLine := range contentsLines {
		if !regexp.MatchString(currentLine) {
			slice = append(slice, currentLine)
		}
	}

	return slice
}
