package whitespace

import "gitlab.com/evatix-go/strhelper/internal/pkg/panichelper"

// AllAsciiWhitespacesCount returns the whitespace (excluding unicode whitespaces) counts
func AllAsciiWhitespacesCount(s *string, startsAt int) int {
	if s == nil || len(*s) == 0 {
		return 0
	}

	length := len(*s)

	if startsAt < 0 || length-1 < startsAt {
		panichelper.StartAtIndexFailed(startsAt, length)
	}

	spacesFound := 0
	var char uint8
	for ; startsAt < length; startsAt++ {
		char = (*s)[startsAt]
		if asciiSpaces[char] == 1 {
			spacesFound++
		}
	}

	return spacesFound
}
