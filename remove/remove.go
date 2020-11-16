package remove

import (
	"unicode"

	"gitlab.com/evatix-go/strhelper/chars"
	"gitlab.com/evatix-go/strhelper/internal/pkg/panichelper"
	"gitlab.com/evatix-go/strhelper/replace"
	"gitlab.com/evatix-go/strhelper/strconst"
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

func GetPtr(str, removeStr *string, startsAt, count int, isCaseSensitive bool) string {
	return replace.ReplacePtr(
		str,
		removeStr,
		strconst.EmptyStringPtr,
		startsAt,
		count,
		isCaseSensitive)
}

func Get(str, removeStr string, startsAt, count int, isCaseSensitive bool) string {
	return replace.ReplacePtr(
		&str,
		&removeStr,
		strconst.EmptyStringPtr,
		startsAt,
		count,
		isCaseSensitive)
}

func UsingRequest(request *strhelpercore.RemoveRequest) string {
	if request == nil {
		panic("Remove request cannot be nil.")
	}

	replaceRequest := request.ToReplaceRequest()

	return replace.UsingReplaceRequest(replaceRequest)
}

func All(str, removeStr string) string {
	return GetPtr(&str, &removeStr, 0, -1, true)
}

func AllPtr(str, removeStr *string) string {
	return GetPtr(str, removeStr, 0, -1, true)
}

func AllWithCase(str, removeStr string, isCaseSensitive bool) string {
	return GetPtr(&str, &removeStr, 0, -1, isCaseSensitive)
}

func AllWithCasePtr(str, removeStr *string, isCaseSensitive bool) string {
	return GetPtr(str, removeStr, 0, -1, isCaseSensitive)
}

// Returns empty string if str is nil or empty.
func Whitespaces(str string, startsAt int) string {
	return WhitespacesPtr(&str, startsAt)
}

// Returns empty string if str is nil or empty.
func WhitespacesPtr(str *string, startsAt int) string {
	if str == nil || len(*str) == 0 {
		return strconst.EmptyString
	}

	length := len(*str)

	if startsAt < 0 || length-1 < startsAt {
		panichelper.StartAtIndexFailed(startsAt, length)
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
func NewLines(str string, startsAt int) string {
	return NewLinesPtr(&str, startsAt)
}

// Returns empty string if str is nil or empty.
// FormFeed \f is also marked as newline here and will be removed from string if has any.
func NewLinesPtr(str *string, startsAt int) string {
	if str == nil || len(*str) == 0 {
		return strconst.EmptyString
	}

	length := len(*str)

	if startsAt < 0 || length-1 < startsAt {
		panichelper.StartAtIndexFailed(startsAt, length)
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
func CommaPtr(
	str *string,
	startsAt int,
) string {
	return CharactersPtr(
		str,
		startsAt,
		&commaRemoveAscIIArray,
		true)
}

// Returns empty string if str is nil or empty.
// remove only those comma from the string.
func HyphenPtr(
	str *string,
	startsAt int,
) string {
	return CharactersPtr(
		str,
		startsAt,
		&hyphenRemoveAscIIArray,
		true)
}

// Returns empty string if str is nil or empty.
// remove only those characters which are given
//
// Limited to ASCII only
func CharactersPtr(
	str *string,
	startsAt int,
	removingCharacters *[256]uint8,
	isCaseSensitive bool,
) string {
	if str == nil || len(*str) == 0 {
		return strconst.EmptyString
	}

	length := len(*str)

	if startsAt < 0 || length-1 < startsAt {
		panichelper.StartAtIndexFailed(startsAt, length)
	}

	removingCharactersCount := chars.CountAscIICharsPtr(
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
