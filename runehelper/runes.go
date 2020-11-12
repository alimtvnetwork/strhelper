package runehelper

import (
	"gitlab.com/evatix-go/strhelper/constants"
)

func IsMatch(rune1 rune, rune2 rune, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return rune1 == rune2
	}

	// Insensitive case
	return ToLowerRune(rune1) == ToLowerRune(rune2)
}

func IsMatchCaseSensitive(rune1 rune, rune2 rune) bool {
	return rune1 == rune2
}

func IsMatchCaseInsensitive(rune1 rune, rune2 rune) bool {
	return ToLowerRune(rune1) == ToLowerRune(rune2)
}

func IsUpperCasePtr(r *rune) bool {
	return *r >= constants.UpperCaseA &&
		*r <= constants.UpperCaseZ
}

func IsUpperCase(r rune) bool {
	return r >= constants.UpperCaseA &&
		r <= constants.UpperCaseZ
}

func ToLowerRune(r rune) rune {
	if r >= constants.UpperCaseA &&
		r <= constants.UpperCaseZ {
		lowerCaseRune := rune(uint8(r) + constants.LowerCase)

		return lowerCaseRune
	}

	return r
}

func ToRuneArray(string *string) []rune {
	return []rune(*string)
}

func ToRuneArrayPtr(string *string) *[]rune {
	val := []rune(*string)

	return &val
}

// Creates new rune array of lowercase runes
// if inputs == nil then returns nil
func ToLowerRunes(inputs *[]rune) *[]rune {
	if inputs == nil {
		return nil
	}

	newRunes := make([]rune, len(*inputs))

	for index, rune := range *inputs {
		if rune >= constants.UpperCaseA && rune <= constants.UpperCaseZ {
			// in upper, make lower
			rune = rune + constants.LowerCase
		}

		newRunes[index] = rune
	}

	return &newRunes
}

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

func GetSafeRunesIndexAtBy(runes *[]rune, index int) rune {
	if !(len(*runes)-1 >= index) {
		return constants.InvalidNotFoundCase
	}

	return (*runes)[index]
}
