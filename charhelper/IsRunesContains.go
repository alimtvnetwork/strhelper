package charhelper

// runes nil results false regardless
// Or else returns true if searchingFor contains in runes
func IsRunesContains(runes *[]rune, searchingFor rune) bool {
	if runes == nil || len(*runes) == 0 {
		return false
	}

	for _, currentRune := range *runes {
		if currentRune == searchingFor {
			return true
		}
	}

	return false
}
