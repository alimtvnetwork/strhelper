package strs

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

func IsEmpty(lines *[]string) bool {
	return lines == nil || *lines == nil || len(*lines) == 0
}

func IndexOf(lines *[]string, findingString *string, startsAtIndex int, isCaseSensitive bool) int {
	if startsAtIndex <= constants.InvalidNotFoundCase {
		panic("Start index cannot be negative.")
	}

	if IsEmpty(lines) {
		return constants.InvalidNotFoundCase
	}

	length := len(*lines)

	if !isCaseSensitive {
		// insensitive
		return IndexOfForCaseInsensitive(lines, findingString, startsAtIndex)
	}

	for i := startsAtIndex; i < length; i++ {
		var currentOne = (*lines)[i]
		if currentOne == *findingString {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}

func IndexOfForCaseInsensitive(lines *[]string, findingString *string, startsAtIndex int) int {
	if startsAtIndex <= constants.InvalidNotFoundCase {
		panic("Start index cannot be negative.")
	}

	if IsEmpty(lines) {
		return constants.InvalidNotFoundCase
	}

	length := len(*lines)
	findingStringToLower := strings.ToLower(*findingString)

	for i := startsAtIndex; i < length; i++ {
		var currentOne = strings.ToLower((*lines)[i])
		if currentOne == findingStringToLower {
			return i
		}
	}

	return constants.InvalidNotFoundCase
}

func IsExists(lines *[]string, findingString *string, isCaseSensitive bool) bool {
	return IndexOf(lines, findingString, 0, isCaseSensitive) > constants.InvalidNotFoundCase
}

// Refers to non empty array !(lines == nil || *lines == nil || len(*lines) == 0)
func HasAnyItems(lines *[]string) bool {
	return !(lines == nil || *lines == nil || len(*lines) == 0)
}

func DoesntExist(lines *[]string, findingString *string, isCaseSensitive bool) bool {
	return !IsExists(lines, findingString, isCaseSensitive)
}

func GetNonEmptyStrings(lines *[]string, isTrimSpace bool) *[]string {
	newLines := make([]string, 0, len(*lines))

	if IsEmpty(lines) {
		return &newLines
	}

	for _, line := range *lines {
		line2 := line

		if isTrimSpace {
			line2 = strings.TrimSpace(line2)
		}

		if line == constants.EmptyString || len(line) == 0 {
			continue
		}

		newLines = append(newLines, line2)
	}

	return &newLines
}
