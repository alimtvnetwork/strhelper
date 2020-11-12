package sw

import (
	"regexp"
	"strings"

	"gitlab.com/evatix-go/strhelper"
	"gitlab.com/evatix-go/strhelper/charhelper"
	"gitlab.com/evatix-go/strhelper/concat"
	"gitlab.com/evatix-go/strhelper/constants"
)

type StringWrapper string

func (stringWrapper *StringWrapper) Value() string {
	return string(*stringWrapper)
}

func (stringWrapper *StringWrapper) ValuePtr() *string {
	value := string(*stringWrapper)

	return &value
}

func (stringWrapper *StringWrapper) Length() int {
	return len(stringWrapper.Value())
}

// Too slow, if you want IsEmptySpace use isEmptyOrWhitespace.
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

// Returns true based on text compare case sensitive.
func (stringWrapper *StringWrapper) IsSensitiveEquals(s *string) bool {
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

// IsNullOrEmpty(s) || strhelper.IsBlankPtr(stringWrapper.ValuePtr())
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

// get uint8 array ptr
func (stringWrapper *StringWrapper) ToUInt8sPtr() *[]uint8 {
	values := []uint8(stringWrapper.Value())

	return &values
}

// get bytes array
func (stringWrapper *StringWrapper) ToBytes() []byte {
	return []byte(stringWrapper.Value())
}

// get bytes array ptr
func (stringWrapper *StringWrapper) ToBytesPtr() *[]byte {
	values := []byte(stringWrapper.Value())

	return &values
}

// get rune array
func (stringWrapper *StringWrapper) ToRunes() []rune {
	return []rune(stringWrapper.Value())
}

// get rune array ptr
func (stringWrapper *StringWrapper) ToRunesPtr() *[]rune {
	runes := []rune(stringWrapper.Value())

	return &runes
}

// get lowercase rune array ptr
func (stringWrapper *StringWrapper) ToLowerRunesPtr() *[]rune {
	runes := []rune(stringWrapper.ToLower())

	return &runes
}

// get uppercase rune array ptr
func (stringWrapper *StringWrapper) ToUpperRunesPtr() *[]rune {
	runes := []rune(stringWrapper.ToUpper())

	return &runes
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

// returns -1 if the index is not present in strings lengthInBytes.
// or else returns the character value from that index
// performance should be very slow, use direct access of str.
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

// Returns a new string builder contains text of stringWrapper and has a
// growth = stringWrapper.lengthInBytes + additionalGrowLength
func (stringWrapper *StringWrapper) Builder(additionalGrowLength int) strings.Builder {
	builder := strings.Builder{}
	currentString := stringWrapper.Value()
	length := len(currentString) + additionalGrowLength
	builder.Grow(length)
	builder.WriteString(currentString)
	return builder
}

// returns -1 if the index is not present in strings lengthInBytes.
// or else returns the character value from that index
// performance should be very slow, use direct access of str.
func (stringWrapper *StringWrapper) GetSafeRuneIndexAt(index int) rune {
	if !stringWrapper.HasIndex(index) {
		return constants.InvalidNotFoundCase
	}

	return stringWrapper.ToRunes()[index]
}

// Returns a new string builder contains text of stringWrapper + str and has a
// growth = stringWrapper.lengthInBytes + additionalGrowLength + len(str)
func (stringWrapper *StringWrapper) BuilderWithStr(str *string, additionalGrowLength int) strings.Builder {
	builder := strings.Builder{}
	currentString := stringWrapper.Value()
	length := len(currentString) + len(*str) + additionalGrowLength
	builder.Grow(length)
	builder.WriteString(currentString)
	builder.WriteString(*str)

	return builder
}

// Better to use slice or builder for appending lines in a loop.
func (stringWrapper *StringWrapper) AppendLines(isSkipOnEmpty bool, contents ...string) StringWrapper {
	return stringWrapper.concat(constants.NewLine, isSkipOnEmpty, &contents)
}

// Better to use slice or builder for appending or concatenating lines in a loop.
func (stringWrapper *StringWrapper) Concat(contents ...string) StringWrapper {
	return stringWrapper.concat(
		constants.EmptyString,
		false, // must add everything
		&contents)
}

// Better to use slice or builder for appending lines.
func (stringWrapper *StringWrapper) ConcatWrappers(
	separator string,
	isSkipOnEmpty bool,
	stringWrappers ...StringWrapper,
) StringWrapper {
	strArray := make([]string, len(stringWrappers))

	for i, wrapper := range stringWrappers {
		strArray[i] = wrapper.Value()
	}

	return stringWrapper.concat(
		separator,
		isSkipOnEmpty,
		&strArray)
}

// Better to use slice or builder for appending lines.
func (stringWrapper *StringWrapper) ConcatWithSeparator(
	separator string,
	isSkipOnEmpty bool,
	contents ...string,
) StringWrapper {
	return stringWrapper.concat(
		separator,
		isSkipOnEmpty,
		&contents)
}

// combine current wrapper strings + all given ones with given separator and returns as wrapper
func (stringWrapper *StringWrapper) concat(
	separator string,
	isSkipOnEmpty bool,
	contents *[]string,
) StringWrapper {
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
	replacedText := strhelper.ReplacePtr(
		stringWrapper.ValuePtr(),
		searchingWrapper.ValuePtr(),
		replacingWrapper.ValuePtr(),
		startsAt,
		replaceCount,
		isCaseSensitive,
	)

	return StringWrapper(replacedText)
}

// For better performance use Ptr version.
func (stringWrapper *StringWrapper) Replace(
	search,
	replaceText string,
	isCaseSensitive bool,
	startsAt int,
	replaceCount int,
) string {
	return strhelper.ReplacePtr(
		stringWrapper.ValuePtr(),
		&search,
		&replaceText,
		startsAt,
		replaceCount,
		isCaseSensitive,
	)
}

func (stringWrapper *StringWrapper) ReplacePtr(
	search,
	replaceText *string,
	startsAt int,
	replaceCount int,
	isCaseSensitive bool,
) string {
	return strhelper.ReplacePtr(
		stringWrapper.ValuePtr(),
		search,
		replaceText,
		startsAt,
		replaceCount,
		isCaseSensitive,
	)
}

func (stringWrapper *StringWrapper) ReplaceAll(
	search,
	replaceText string,
	isCaseSensitive bool,
	startsAt int,
) string {
	return strhelper.ReplacePtr(
		stringWrapper.ValuePtr(),
		&search,
		&replaceText,
		startsAt,
		-1,
		isCaseSensitive,
	)
}

func (stringWrapper *StringWrapper) LastIndexOf(
	search string,
	lastStartIndexReducedBy int,
	isCaseSensitive bool,
) int {
	return strhelper.LastIndexOfPtr(
		stringWrapper.ValuePtr(),
		&search,
		lastStartIndexReducedBy,
		isCaseSensitive)
}

func (stringWrapper *StringWrapper) LastIndexOfPtr(
	search *string,
	lastStartIndexReducedBy int,
	isCaseSensitive bool,
) int {
	return strhelper.LastIndexOfPtr(
		stringWrapper.ValuePtr(),
		search,
		lastStartIndexReducedBy,
		isCaseSensitive)
}

// For better performance use strhelper.IsStartsWithPtr
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

// Use direct strhelper.IsEndsWithPtr will be faster
func (stringWrapper *StringWrapper) IsEndsWith(
	endsWith string,
	isCaseSensitive bool,
	startsAt int,
) bool {
	return strhelper.IsEndsWithPtr(
		stringWrapper.ValuePtr(),
		&endsWith,
		startsAt,
		isCaseSensitive)
}

// Use direct strhelper.IsEndsWithPtr will be faster
func (stringWrapper *StringWrapper) IsEndsWithPtr(
	endsWith *string,
	isCaseSensitive bool,
	startsAt int,
) bool {
	return strhelper.IsEndsWithPtr(
		stringWrapper.ValuePtr(),
		endsWith,
		startsAt,
		isCaseSensitive)
}

func (stringWrapper *StringWrapper) PadLeftWithSpace(width int) string {
	return strhelper.PadSpaceLeft(stringWrapper.ValuePtr(), width)
}

func (stringWrapper *StringWrapper) PadRightWithSpace(width int) string {
	return strhelper.PadSpaceRight(stringWrapper.ValuePtr(), width)
}

func (stringWrapper *StringWrapper) PadLeft(width int, padding string) string {
	return strhelper.PadLeft(stringWrapper.ValuePtr(), &padding, width)
}

func (stringWrapper *StringWrapper) PadRight(width int, padding string) string {
	return strhelper.PadRight(stringWrapper.ValuePtr(), &padding, width)
}

func (stringWrapper *StringWrapper) Pad(width int, padding string, isLeft, isRight bool) string {
	return strhelper.Pad(stringWrapper.ValuePtr(), &padding, width, isLeft, isRight)
}
