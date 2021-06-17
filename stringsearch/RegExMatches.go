package stringsearch

import "regexp"

// RegExMatches returns the lines which meet with regex requirements
//
// Alias for GetElementsByRegExMatch
func RegExMatches(
	regexp *regexp.Regexp,
	lines ...string,
) []string {
	if len(lines) == 0 {
		return []string{}
	}

	return GetElementsNonMatchingByRegEx(lines, regexp)
}
