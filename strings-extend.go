package strhelper

import (
	"regexp"
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

// IsEmpty(s)
func IsNullOrEmpty(s string) bool {
	return s == constants.EmptyString || len(s) == 0
}

// IsNullOrEmpty(s) || len(strings.TrimSpace(s)) == 0
func IsNullOrWhitespace(s string) bool {
	return s == constants.EmptyString || len(s) == 0 || len(strings.TrimSpace(s)) == 0
}

// returns len(s) == 0 || s == constants.EmptyString
func IsEmpty(s string) bool {
	return s == constants.EmptyString || len(s) == 0
}

// returns true if IsNullOrWhitespace(s)
func IsBlank(s string) bool {
	return s == constants.EmptyString || len(s) == 0 || len(strings.TrimSpace(s)) == 0
}

// Has at least one character any
func HasCharacter(s string) bool {
	return !(s == constants.EmptyString || len(s) == 0 || len(s) == 0)
}

// Has at least one character other than space or whitespace
func HasCharacterWithoutSpaces(s string) bool {
	return !(s == constants.EmptyString || len(s) == 0 || len(strings.TrimSpace(s)) == 0)
}

// Has at least one character other than space or whitespace
func IsDefinedWithCharsWithoutSpaces(s string) bool {
	return !(s == constants.EmptyString || len(s) == 0 || len(strings.TrimSpace(s)) == 0)
}

func IndexOf(s, findingString string, startsAt int, isCaseSensitive bool) int {
	if isCaseSensitive {
		return strings.Index(s, findingString)
	}

	return IndexOfLongestCommonSuffix(s, findingString, startsAt, isCaseSensitive)
}

// Copied from golang library (reference : https://bit.ly/3oPHGdy).
// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func JoinPtrExceptEmpty(elements *[]string, sep *string) string {
	elementsLength := len(*elements)

	switch elementsLength {
	case 0:
		return constants.EmptyString
	case 1:
		return (*elements)[0]
	}

	n := len(*sep) * (elementsLength - 1)
	for i := 0; i < elementsLength; i++ {
		n += len((*elements)[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString((*elements)[0])
	restOfTheElementsExceptFirst := (*elements)[1:]
	for _, s := range restOfTheElementsExceptFirst {
		if s == constants.EmptyString || len(s) == 0 {
			continue
		}

		b.WriteString(*sep)
		b.WriteString(s)
	}

	return b.String()
}

// Copied from golang library (reference : https://bit.ly/3oPHGdy).
// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func JoinPtrExceptFor(filterSkipMap *map[string]byte, elements *[]string, sep *string) string {
	elementsLength := len(*elements)
	if elementsLength == 0 {
		return constants.EmptyString
	}

	n := len(*sep) * (elementsLength - 1)
	for i := 0; i < elementsLength; i++ {
		n += len((*elements)[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString((*elements)[0])
	restOfTheElementsExceptFirst := (*elements)[1:]
	for _, s := range restOfTheElementsExceptFirst {
		_, has := (*filterSkipMap)[s]
		if has {
			continue
		}

		b.WriteString(*sep)
		b.WriteString(s)
	}

	return b.String()
}

func GetNotImplementedPanicMessage(url string) string {
	return constants.NotImplemented + " : [TODO] Will be solved at (" + url + ")"
}

// Copied from golang library (reference : https://bit.ly/3oPHGdy).
// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func JoinPtr(elements *[]string, sep *string) string {
	elementsLength := len(*elements)

	switch elementsLength {
	case 0:
		return constants.EmptyString
	case 1:
		return (*elements)[0]
	}

	n := len(*sep) * (elementsLength - 1)
	for i := 0; i < elementsLength; i++ {
		n += len((*elements)[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString((*elements)[0])
	restOfTheElementsExceptFirst := (*elements)[1:]
	for _, s := range restOfTheElementsExceptFirst {
		b.WriteString(*sep)
		b.WriteString(s)
	}

	return b.String()
}

// String join using Unix New Line operating system newline (For windows it is \r\n and for unix it is \n)
func GetContentFormLines(lines []string) string {
	return JoinPtr(&lines, constants.NewLinePtr)
}

// String join using Unix New Line "\n"
func GetContentFormLinesUnix(lines []string) string {
	return JoinPtr(&lines, constants.NewLineUnixPtr)
}

// Gets new line by os specific new line
func GetLines(content string) []string {
	return strings.Split(content, constants.NewLine)
}

// Gets new line by \n
func GetLinesUnix(content string) []string {
	return strings.Split(content, constants.NewLineUnix)
}

// Gets new line by \n
func GetParsedLinesFrom(content string, regexps ...*regexp.Regexp) []*strhelpercore.RegExResultWrapper {
	results := make([]*strhelpercore.RegExResultWrapper, 0, len(regexps))

	for index, regex := range regexps {
		wrapper := strhelpercore.NewRegExResultWrapper(index, &content, regex)
		results = append(results, &wrapper)
	}

	return results
}

// Code Copied from Reference: https://bit.ly/35ZGJHc
// Has Longest common suffix, returns -1 if doesn't found.
// If any is nil or has constants.EmptyString empty string then it return -1
// TODO : https://gitlab.com/evatix-go/strhelper/-/issues/2
func IndexOfLongestCommonSuffix(
	a, b string,
	startsAt int,
	isCaseSensitive bool,
) int {
	panic(constants.NotImplemented + " https://gitlab.com/evatix-go/strhelper/-/issues/2")
	if IsNullOrEmpty(a) || IsNullOrEmpty(b) {
		return constants.InvalidNotFoundCase
	}

	lenA := len(a)
	lenB := len(b)

	if isCaseSensitive {
		// both needs to be in same case
		a = strings.ToLower(a)
		b = strings.ToLower(b)
	}

	i := startsAt
	for ; i < lenA && i < lenB; i++ {
		if a[lenA-1-i] != b[lenB-1-i] {
			return i
		}
	}

	return i
}

func IsExists(
	s, findingString string,
	isCaseSensitive bool,
) bool {
	return IndexOf(s, findingString, 0, isCaseSensitive) > constants.InvalidNotFoundCase
}

func DoesntExist(
	s, findingString string,
	isCaseSensitive bool,
) bool {
	return !IsExists(s, findingString, isCaseSensitive)
}

func ToUInt8Array(string string) []uint8 {
	return []uint8(string)
}

func ToBytesArray(string string) []byte {
	return []byte(string)
}

func ToUInt8ArrayPtr(string string) *[]uint8 {
	val := []uint8(string)

	return &val
}

func ToBytesArrayPtr(string string) *[]byte {
	val := []byte(string)

	return &val
}

// returns true if len(str) >= length
func HasLength(str string, length int) bool {
	return len(str) >= length
}

// returns true if len(str)-1 >= index
func HasIndex(str string, index int) bool {
	return len(str)-1 >= index
}

// language integrated ones will be faster str[startAtIndex:endsAtIndex]
// Under the hood this method usages that functionality from language
// In case length is larger than given index, then it will fix to the length end
// panics if startsAtIndex < 0
// if endsAtIndex-startsAtIndex <= 0 then returns EmptyString(constants.EmptyString)
func SafeSubstringAtIndex(
	str string,
	startsAtIndex, endsAtIndex int,
) string {
	if startsAtIndex < 0 {
		message := "Substring Index cannot have negative startsAtIndex : " + string(startsAtIndex)

		panic(message)
	}

	if len(str)-1 < endsAtIndex {
		endsAtIndex = len(str) - 1
	}

	if endsAtIndex-startsAtIndex <= 0 {
		return constants.EmptyString
	}

	return str[startsAtIndex:endsAtIndex]
}

// language integrated ones will be faster str[startAtIndex:endsAtIndex]
// Under the hood this method usages that functionality from language
// panics if startsAtIndex < 0
func SubstringAtIndex(
	str string,
	startsAtIndex, endsAtIndex int,
) string {
	if startsAtIndex < 0 {
		message := "Substring Index cannot have negative startsAtIndex : " + string(startsAtIndex)

		panic(message)
	}

	return str[startsAtIndex:endsAtIndex]
}

// Use language integrated one that will be faster. eg. str[startsAtIndex : startsAtIndex+length]
// If want to use by function then pointer one would be more efficient than this one.
func SubstringByLength(str string, startsAtIndex, length int) string {
	return str[startsAtIndex : startsAtIndex+length]
}

// startsAtLastIndex = 0 meaning starts from last position, giving 2 meaning len(wholeText)-2
// If EmptyString(constants.EmptyString) is given for search and if the startsAt less than the length of the wholeText then it returns true.
func IsEndsWith(
	wholeText, endsWithSearch string,
	startsAtLastIndex int,
	isCaseSensitive bool,
) bool {
	if IsEmpty(endsWithSearch) {
		return wholeText == constants.EmptyString && startsAtLastIndex == 0 || len(wholeText)-1 >= startsAtLastIndex
	}

	if IsEmpty(wholeText) {
		return endsWithSearch == constants.EmptyString && startsAtLastIndex == 0
	}

	textLength := len(wholeText) - startsAtLastIndex
	searchLength := len(endsWithSearch)
	if searchLength > textLength {
		return false
	}

	startingIndex := textLength - searchLength
	if startingIndex > textLength {
		return false
	}

	substringFromText := wholeText[startingIndex:textLength]

	if isCaseSensitive {
		return substringFromText == endsWithSearch
	}

	lowerCaseSubstring := strings.ToLower(substringFromText)
	lowerCaseEndSearchText := strings.ToLower(endsWithSearch)

	// insensitive
	return lowerCaseSubstring == lowerCaseEndSearchText
}

// startsAt = 0 meaning starts from last position, giving 2 meaning start index += 2
// If EmptyString(constants.EmptyString) is given for search and if the startsAt less than the length of the wholeText then it returns true.
func IsStartsWith(
	wholeText, startsWith string,
	startsAt int,
	isCaseSensitive bool,
) bool {
	if IsEmpty(startsWith) {
		return wholeText == constants.EmptyString && startsAt == 0 || len(wholeText)-1 >= startsAt
	}

	if IsEmpty(wholeText) {
		return startsWith == constants.EmptyString && startsAt == 0
	}

	textLength := len(wholeText) - startsAt
	searchLength := len(startsWith)
	if searchLength > textLength {
		return false
	}

	endingLength := startsAt + searchLength
	substringFromText := wholeText[startsAt:endingLength]

	if isCaseSensitive {
		return substringFromText == startsWith
	}

	lowerCaseSubstring := strings.ToLower(substringFromText)
	lowerCaseSearchText := strings.ToLower(startsWith)

	// insensitive
	return lowerCaseSubstring == lowerCaseSearchText
}
