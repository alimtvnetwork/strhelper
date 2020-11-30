package whitespace

func AllNewLinesCount(s *string, startsAt int) int {
	if s == nil || len(*s) == 0 {
		return 0
	}

	allRunes := []rune(*s)

	return AllNewLinesCountOfRunes(&allRunes, startsAt)
}
