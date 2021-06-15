package whitespace

// AllWhitespacesCount Returns the whitespace (including unicode whitespaces) counts
//
// Note: Since unicode is in the calculation, it requires str to be in ([]rune)
// format which requires more memory and cost.
func AllWhitespacesCount(s *string, startsAt int) int {
	if s == nil || len(*s) == 0 {
		return 0
	}

	allRunes := []rune(*s)

	return AllWhitespacesCountOfRunes(
		&allRunes,
		startsAt)
}
