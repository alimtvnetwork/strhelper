package whitespace

import "gitlab.com/evatix-go/strhelper/internal/panichelper"

// Returns the whitespace (excluding unicode whitespaces) indexes
//
// Returns nil if s is nil or empty
func GetAsciiWhitespaceIndexes(s *string, startsAt int) *[]int {
	if s == nil || len(*s) == 0 {
		return nil
	}

	length := len(*s)

	if startsAt < 0 || length-1 < startsAt {
		panichelper.StartAtIndexFailed(startsAt, length)
	}

	indexes := make([]int, 0, length/2)
	hasFoundAny := false

	var uin8 uint8
	for ; startsAt < length; startsAt++ {
		uin8 = (*s)[startsAt]
		if asciiSpaces[uin8] == 1 {
			indexes = append(indexes, startsAt)
			hasFoundAny = true
		}
	}

	if !hasFoundAny {
		return nil
	}

	return &indexes
}
