package sw

import (
	"regexp"
	"strings"

	"gitlab.com/evatix-go/strhelper"
	"gitlab.com/evatix-go/strhelper/charhelper"
	"gitlab.com/evatix-go/strhelper/concat"
	"gitlab.com/evatix-go/strhelper/constants"
)

// TODO : https://gitlab.com/evatix-go/strhelper/-/issues/3
type StringWrapper string

func (stringWrapper *StringWrapper) Value() string {
	return string(*stringWrapper)
}

func (stringWrapper *StringWrapper) ValuePtr() *string {
	value := stringWrapper.Value()

	return &value
}

func (stringWrapper *StringWrapper) Length() int {
	return len(stringWrapper.Value())
}

func (stringWrapper *StringWrapper) LengthWithoutWhitespaces() int {
	return len(strings.TrimSpace(stringWrapper.Value()))
}

func (stringWrapper *StringWrapper) IsEquals(s string, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return s == stringWrapper.Value()
	}

	// insensitive
	lower := strings.ToLower(s)

	return lower == stringWrapper.ToLower()
}

func (stringWrapper *StringWrapper) IsEqualsPtr(s *string) bool {
	if s == nil {
		return false
	}

	return *s == stringWrapper.Value()
}

// returns len(s) == 0 || s == ""
func (stringWrapper *StringWrapper) IsEmpty() bool {
	return (*stringWrapper).Length() == 0 || string(*stringWrapper) == constants.EmptyString
}

// IsNull(s) || IsEmpty(s)
func (stringWrapper *StringWrapper) IsNullOrEmpty() bool {
	return (*stringWrapper).IsEmpty()
}

// IsNullOrEmpty(s) || len(strings.TrimSpace(s)) == 0
func (stringWrapper *StringWrapper) IsNullOrWhitespace() bool {
	return (*stringWrapper).IsNullOrEmpty() || strhelper.IsBlankPtr(stringWrapper.ValuePtr())
}

func (stringWrapper *StringWrapper) TrimSpace() string {
	return strings.TrimSpace((*stringWrapper).Value())
}

func (stringWrapper *StringWrapper) Trim(cutSet string) string {
	return strings.Trim((*stringWrapper).Value(), cutSet)
}

func (stringWrapper *StringWrapper) TrimLeft(cutSet string) string {
	return strings.TrimLeft((*stringWrapper).Value(), cutSet)
}

func (stringWrapper *StringWrapper) TrimRight(cutSet string) string {
	return strings.TrimRight((*stringWrapper).Value(), cutSet)
}

// get uint8 array
func (stringWrapper *StringWrapper) ToUInt8s() []uint8 {
	return []uint8(stringWrapper.Value())
}

// get bytes array
func (stringWrapper *StringWrapper) ToBytes() []byte {
	return []byte(stringWrapper.Value())
}

// get rune array
func (stringWrapper *StringWrapper) ToRunes() []rune {
	return []rune(stringWrapper.Value())
}

// returns true if IsNullOrWhitespace(s)
func (stringWrapper *StringWrapper) IsBlank() bool {
	return (*stringWrapper).IsNullOrWhitespace()
}

// Has at least one character other than space or whitespace
func (stringWrapper *StringWrapper) HasCharacter() bool {
	return !(*stringWrapper).IsNullOrWhitespace()
}

// Has at least one character other than space or whitespace
func (stringWrapper *StringWrapper) IsDefined() bool {
	return !(*stringWrapper).IsNullOrWhitespace()
}

func (stringWrapper *StringWrapper) ToLower() string {
	return strings.ToLower(stringWrapper.Value())
}

func (stringWrapper *StringWrapper) ToUpper() string {
	return strings.ToUpper(stringWrapper.Value())
}

func (stringWrapper *StringWrapper) ToLowerWrapper() StringWrapper {
	return StringWrapper(strings.ToLower(stringWrapper.Value()))
}

// Returns strings to upper case as StringWrapper
func (stringWrapper *StringWrapper) ToUpperWrapper() StringWrapper {
	return StringWrapper(strings.ToUpper(stringWrapper.Value()))
}

// Returns character at the given index, if not exist then panic.
func (stringWrapper *StringWrapper) At(index int) uint8 {
	return stringWrapper.Value()[index]
}

// Create regular expression from current string.
// Recommendation, do not create regular expressions inside a function call, keep it on top of the file as variables
func (stringWrapper *StringWrapper) CreateRegularExpression() *regexp.Regexp {
	return regexp.MustCompile(stringWrapper.Value())
}

// Same as Value()
func (stringWrapper *StringWrapper) String() string {
	return string(*stringWrapper)
}

// returns -1 if the index is not present in strings length.
// or else returns the character value from that index
func (stringWrapper *StringWrapper) GetSafeIndexAt(index int) int16 {
	if !stringWrapper.HasIndex(index) {
		return constants.InvalidNotFoundCase
	}

	return int16(stringWrapper.Value()[index])
}

func (stringWrapper *StringWrapper) IsEqualAtIndex(
	index int,
	char uint8,
	isCaseSensitive bool,
) bool {
	valueAt := stringWrapper.Value()[index]

	if isCaseSensitive {
		return valueAt == char
	}

	return charhelper.IsMatchCaseInsensitive(valueAt, char)
}

func (stringWrapper *StringWrapper) HasIndex(index int) bool {
	return (*stringWrapper).Length()-1 >= index
}

func (stringWrapper *StringWrapper) Builder(additionalGrowLength int) strings.Builder {
	builder := strings.Builder{}
	length := stringWrapper.Length() + additionalGrowLength
	builder.Grow(length)

	return builder
}

func (stringWrapper *StringWrapper) AppendLines(isSkipOnEmpty bool, contents ...string) StringWrapper {
	return stringWrapper.concat(constants.NewLine, isSkipOnEmpty, &contents)
}

func (stringWrapper *StringWrapper) Concat(contents ...string) StringWrapper {
	return stringWrapper.concat(
		constants.EmptyString,
		false, // must add everything
		&contents)
}

func (stringWrapper *StringWrapper) ConcatWithSeparator(
	separator string, isSkipOnEmpty bool, contents ...string,
) StringWrapper {
	return stringWrapper.concat(
		separator,
		isSkipOnEmpty,
		&contents)
}

// combine current wrapper strings + all given ones with given separator and returns as wrapper
func (stringWrapper *StringWrapper) concat(separator string, isSkipOnEmpty bool, contents *[]string) StringWrapper {
	combinedResult := concat.StringsArrayWithSeparator(
		stringWrapper.ValuePtr(),
		&separator,
		isSkipOnEmpty,
		contents)

	return StringWrapper(combinedResult)
}

func (stringWrapper *StringWrapper) ReplaceWrapper(
	searchingWrapper,
	replacingWrapper *StringWrapper,
	isCaseSensitive bool,
	startsAt int,
	replaceCount int,
) StringWrapper {
	panic("Not Implemented")
}

func (stringWrapper *StringWrapper) Replace(
	search,
	newText string,
	isCaseSensitive bool,
	startsAt int,
	replaceCount int,
) string {
	panic("Not Implemented")
}

func (stringWrapper *StringWrapper) ReplaceAll(
	search,
	newText string,
	isCaseSensitive bool,
	startsAt int,
) string {
	panic("Not Implemented")
}

func (stringWrapper *StringWrapper) LastIndexOf(
	search string,
	isCaseSensitive bool,
	startsLastIndexAt int,
) int {
	panic("Not Implemented")
}

func (stringWrapper *StringWrapper) IsStartsWith(
	search string,
	isCaseSensitive bool,
	startsAt int,
) bool {
	return strhelper.IsStartsWithPtr(
		stringWrapper.ValuePtr(),
		&search,
		startsAt,
		isCaseSensitive)
}

func (stringWrapper *StringWrapper) IsEndsWith(
	endsWith string,
	isCaseSensitive bool,
	startsAt int,
) bool {
	// TODO : Move to pointer implementation later
	return strhelper.IsEndsWith(
		stringWrapper.Value(),
		endsWith,
		startsAt,
		isCaseSensitive)
}

// TODO : https://gitlab.com/evatix-go/strhelper/-/issues/5
func (stringWrapper *StringWrapper) PadLeftWithSpace(width int) string {
	panic("Not Implemented : https://gitlab.com/evatix-go/strhelper/-/issues/5")
}

// TODO : https://gitlab.com/evatix-go/strhelper/-/issues/5
func (stringWrapper *StringWrapper) PadRightWithSpace(width int) string {
	panic("Not Implemented : https://gitlab.com/evatix-go/strhelper/-/issues/5")
}

// TODO : https://gitlab.com/evatix-go/strhelper/-/issues/5
func (stringWrapper *StringWrapper) PadLeft(width int, char uint8) string {
	panic("Not Implemented : https://gitlab.com/evatix-go/strhelper/-/issues/5")
}

// TODO : https://gitlab.com/evatix-go/strhelper/-/issues/5
func (stringWrapper *StringWrapper) PadRight(width int, char uint8) string {
	panic("Not Implemented : https://gitlab.com/evatix-go/strhelper/-/issues/5")
}

// Code Copied from Reference: https://bit.ly/35ZGJHc
// Has Longest common suffix, returns -1 if doesn't found.
// If any is nil or has "" empty string then it return -1
// TODO : https://gitlab.com/evatix-go/strhelper/-/issues/2
func (stringWrapper *StringWrapper) IndexOfLongestCommonSuffix(
	b string,
	startsAt int,
	isCaseSensitive bool,
) int {
	panic("Not implemented : https://gitlab.com/evatix-go/strhelper/-/issues/2")

	bWrapper := StringWrapper(b)
	if stringWrapper.IsNullOrEmpty() || bWrapper.IsNullOrEmpty() {
		return constants.InvalidNotFoundCase
	}

	lenA := stringWrapper.Length()
	lenB := bWrapper.Length()

	if !isCaseSensitive {
		// both needs to be in same case
	}

	i := startsAt
	for ; i < lenA && i < lenB; i++ {
		if stringWrapper.At(lenA-1-i) != b[lenB-1-i] {
			return i
		}
	}

	return i
}
