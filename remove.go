package strhelper

import (
	"unicode"

	"gitlab.com/evatix-go/strhelper/charhelper"
	"gitlab.com/evatix-go/strhelper/constants"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

var (
	asciiSpaceArray       = whitespace.GetAscIISpaceArray()
	asciiNewLinesArray    = whitespace.GetAscIINewLinesArray()
	commaRemoveAscIIArray = [256]uint8{
		',': 1,
	}
	hyphenRemoveAscIIArray = [256]uint8{
		'-': 1,
	}
)

func RemovePtr(str, removeStr *string, startsAt, count int, isCaseSensitive bool) string {
	return ReplacePtr(
		str,
		removeStr,
		constants.EmptyStringPtr,
		startsAt,
		count,
		isCaseSensitive)
}

func Remove(str, removeStr string, startsAt, count int, isCaseSensitive bool) string {
	return ReplacePtr(
		&str,
		&removeStr,
		constants.EmptyStringPtr,
		startsAt,
		count,
		isCaseSensitive)
}

func RemoveUsingRequest(request *strhelpercore.RemoveRequest) string {
	if request == nil {
		panic("Remove request cannot be nil.")
	}

	replaceRequest := request.ToReplaceRequest()

	return ReplaceUsingReplaceRequest(replaceRequest)
}

func RemoveAll(str, removeStr string) string {
	return RemovePtr(&str, &removeStr, 0, -1, true)
}

func RemoveAllPtr(str, removeStr *string) string {
	return RemovePtr(str, removeStr, 0, -1, true)
}

func RemoveAllWithCase(str, removeStr string, isCaseSensitive bool) string {
	return RemovePtr(&str, &removeStr, 0, -1, isCaseSensitive)
}

func RemoveAllWithCasePtr(str, removeStr *string, isCaseSensitive bool) string {
	return RemovePtr(str, removeStr, 0, -1, isCaseSensitive)
}

// Returns empty string if str is nil or empty.
func RemoveWhitespaces(str string, startsAt int) string {
	return RemoveWhitespacesPtr(&str, startsAt)
}

// Returns empty string if str is nil or empty.
func RemoveWhitespacesPtr(str *string, startsAt int) string {
	if str == nil || len(*str) == 0 {
		return constants.EmptyString
	}

	length := len(*str)

	if startsAt < 0 || length-1 < startsAt {
		startAtIndexFailed(startsAt, length)
	}

	chars := make([]byte, length-whitespace.AllWhitespaceCount(str, startsAt))

	for i := 0; i < startsAt; i++ {
		// copy as is
		chars[i] = (*str)[i]
	}

	wordIndex := startsAt
	for ; startsAt < length; startsAt++ {
		char := (*str)[startsAt]
		if !(asciiSpaceArray[char] == 1 || unicode.IsSpace(rune(char))) {
			chars[wordIndex] = (*str)[startsAt]
			wordIndex++
		}
	}

	return string(chars)
}

// Returns empty string if str is nil or empty.
// FormFeed \f is also marked as newline here and will be removed from string if has any.
func RemoveNewLines(str string, startsAt int) string {
	return RemoveNewLinesPtr(&str, startsAt)
}

// Returns empty string if str is nil or empty.
// FormFeed \f is also marked as newline here and will be removed from string if has any.
func RemoveNewLinesPtr(str *string, startsAt int) string {
	if str == nil || len(*str) == 0 {
		return constants.EmptyString
	}

	length := len(*str)

	if startsAt < 0 || length-1 < startsAt {
		startAtIndexFailed(startsAt, length)
	}

	chars := make([]byte, length-whitespace.AllNewLinesCount(str, startsAt))

	for i := 0; i < startsAt; i++ {
		// copy as is
		chars[i] = (*str)[i]
	}

	wordIndex := startsAt
	for ; startsAt < length; startsAt++ {
		char := (*str)[startsAt]
		if !(asciiNewLinesArray[char] == 1) {
			chars[wordIndex] = (*str)[startsAt]
			wordIndex++
		}
	}

	return string(chars)
}

// Returns empty string if str is nil or empty.
// remove only those comma from the string.
func RemoveCommaPtr(
	str *string,
	startsAt int,
) string {
	return RemoveCharactersPtr(
		str,
		startsAt,
		&commaRemoveAscIIArray,
		true)
}

// Returns empty string if str is nil or empty.
// remove only those comma from the string.
func RemoveHyphenPtr(
	str *string,
	startsAt int,
) string {
	return RemoveCharactersPtr(
		str,
		startsAt,
		&hyphenRemoveAscIIArray,
		true)
}

// Returns empty string if str is nil or empty.
// remove only those characters which are given
//
// Limited to ASCII only
func RemoveCharactersPtr(
	str *string,
	startsAt int,
	removingCharacters *[256]uint8,
	isCaseSensitive bool,
) string {
	if str == nil || len(*str) == 0 {
		return constants.EmptyString
	}

	length := len(*str)

	if startsAt < 0 || length-1 < startsAt {
		startAtIndexFailed(startsAt, length)
	}

	removingCharactersCount := charhelper.CountAscIICharsPtr(
		str,
		removingCharacters,
		startsAt,
		isCaseSensitive)

	if removingCharactersCount == 0 {
		return *str
	}

	chars := make([]byte, length-removingCharactersCount)

	for i := 0; i < startsAt; i++ {
		// copy as is
		chars[i] = (*str)[i]
	}

	wordIndex := startsAt
	for ; startsAt < length; startsAt++ {
		char := (*str)[startsAt]
		if !(removingCharacters[char] == 1) {
			chars[wordIndex] = (*str)[startsAt]
			wordIndex++
		}
	}

	return string(chars)
}
