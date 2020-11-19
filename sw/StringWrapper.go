// String Wrapper, no caching, non optimized
package sw

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"gitlab.com/evatix-go/strhelper/chars"
	"gitlab.com/evatix-go/strhelper/concat"
	"gitlab.com/evatix-go/strhelper/index"
	"gitlab.com/evatix-go/strhelper/isstr"
	"gitlab.com/evatix-go/strhelper/lines"
	"gitlab.com/evatix-go/strhelper/padding"
	"gitlab.com/evatix-go/strhelper/remove"
	"gitlab.com/evatix-go/strhelper/replace"
	"gitlab.com/evatix-go/strhelper/strconst"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

// no caching
type StringWrapper struct {
	content    *string
	runeLength *int
	length     int
}

func New(str string) *StringWrapper {
	stringWrapper := StringWrapper{
		content:    &str,
		runeLength: nil,
		length:     len(str),
	}

	return &stringWrapper
}

func NewPtr(str *string) *StringWrapper {
	stringWrapper := StringWrapper{
		content:    str,
		runeLength: nil,
		length:     len(*str),
	}

	return &stringWrapper
}

// use swp, it is optimized for performance.
//
// Recommendation: Use ValuePtr instead of Value()
func (stringWrapper *StringWrapper) Value() string {
	return *(*stringWrapper).content
}

// use swp, it is optimized for performance.
func (stringWrapper *StringWrapper) ValuePtr() *string {
	return (*stringWrapper).content
}

// use swp, it is optimized for performance.
//
// Recommendation: Use ValuePtr / StringPtr() instead of Value() or String()
func (stringWrapper *StringWrapper) String() string {
	return *(*stringWrapper).content
}

// use swp, it is optimized for performance.
func (stringWrapper *StringWrapper) StringPtr() *string {
	return (*stringWrapper).content
}

// use swp, it is optimized for performance.
//
// This verion is not performance optimized.
func (stringWrapper *StringWrapper) LengthInBytes() int {
	return stringWrapper.length
}

// use swp, it is optimized for performance.
//
// Returns utf8.RuneCountInString(stringWrapper.Value())
//
// If don't care about unicode then use len(str) which is LengthInBytes
func (stringWrapper *StringWrapper) Length() int {
	if stringWrapper.runeLength == nil {
		allRunes := stringWrapper.ToRunesPtr()
		runesLength := len(*allRunes)
		stringWrapper.runeLength = &runesLength
	}

	return *stringWrapper.runeLength
}

func (stringWrapper *StringWrapper) HasLengthOf(length int) bool {
	return (*stringWrapper).Length() >= length
}

// Extremely slow, return rune count of after trimming the string, utf8.RuneCountInString(strings.TrimSpace(stringWrapper.Value())).
//
// If you want IsEmptySpace use isEmptyOrWhitespace.
//
// If don't care about unicode use len(str) which is LengthInBytes
func (stringWrapper *StringWrapper) LengthWithoutWhitespaces() int {
	return utf8.RuneCountInString(strings.TrimSpace(stringWrapper.Value()))
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
	return (*stringWrapper).Length() == 0 || *(*stringWrapper).content == strconst.EmptyString
}

// IsNull(s) || IsEmpty(s)
func (stringWrapper *StringWrapper) IsNullOrEmpty() bool {
	return (*stringWrapper).IsEmpty()
}

// IsNullOrEmpty(s) || strhelper.IsBlankPtr(stringWrapper.ValuePtr())
func (stringWrapper *StringWrapper) IsNullOrWhitespace() bool {
	return (*stringWrapper).IsNullOrEmpty() || isstr.BlankPtr(stringWrapper.ValuePtr())
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

func (stringWrapper *StringWrapper) ToLowerWrapper() *StringWrapper {
	return New(strings.ToLower(stringWrapper.Value()))
}

// Returns strings to upper case as StringWrapper
func (stringWrapper *StringWrapper) ToUpperWrapper() *StringWrapper {
	return New(strings.ToUpper(stringWrapper.Value()))
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

// returns -1 if the index is not present in strings lengthInBytes.
// or else returns the character value from that index
// performance should be very slow, use direct access of str.
func (stringWrapper *StringWrapper) GetSafeIndexAt(index int) int16 {
	if !stringWrapper.HasIndex(index) {
		return strconst.InvalidNotFoundCase
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

	return chars.IsMatchCaseInsensitive(valueAt, char)
}

func (stringWrapper *StringWrapper) HasIndex(index int) bool {
	return (*stringWrapper).Length()-1 >= index
}

func (stringWrapper *StringWrapper) RuneLength() int {
	return utf8.RuneCountInString((*stringWrapper).Value())
}

// (*stringWrapper).Length()-1 >= index
func (stringWrapper *StringWrapper) HasRuneIndex(index int) bool {
	return (*stringWrapper).RuneLength()-1 >= index
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
		return strconst.InvalidNotFoundCase
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
func (stringWrapper *StringWrapper) AppendLines(isSkipOnEmpty bool, contents ...string) *StringWrapper {
	return stringWrapper.Concatenates(strconst.NewLine, isSkipOnEmpty, &contents)
}

// Add the contents before the content of StringWrapper.Value()
func (stringWrapper *StringWrapper) Prepend(contents ...string) *StringWrapper {
	return stringWrapper.Prepends(
		strconst.EmptyString,
		false, // must add everything
		&contents)
}

// Line is the separator for add the content before the content of StringWrapper.Value()
func (stringWrapper *StringWrapper) PrependLines(contents ...string) *StringWrapper {
	return stringWrapper.Prepends(
		strconst.NewLine,
		false, // must add everything
		&contents)
}

// Better to use slice or builder for appending or concatenating lines in a loop.
//
// StringWrapper.ValuePtr() + contents joined.
func (stringWrapper *StringWrapper) Concat(contents ...string) *StringWrapper {
	return stringWrapper.Concatenates(
		strconst.EmptyString,
		false, // must add everything
		&contents)
}

// GetLines splitted by newline of os
//
// Windows (`\r\n`), unix (`\n`) - darwin/macos/linux
func (stringWrapper *StringWrapper) GetLines() *[]string {
	lines := strings.Split(*stringWrapper.content, strconst.NewLine)

	return &lines
}

// GetLines splitted by newline using unix Split `\n`
func (stringWrapper *StringWrapper) GetLinesUnix() *[]string {
	lines := strings.Split(*stringWrapper.content, strconst.NewLineUnix)

	return &lines
}

// Loops through all the rune characters
//
// Slower than direct access
func (stringWrapper *StringWrapper) LoopRunes(
	simpleLoopProcessor func(args *strhelpercore.StringWrapperRuneLoopArgs) *string,
) *[]*string {
	runes := *stringWrapper.ToRunesPtr()
	newStrings := make([]*string, len(runes))

	args := strhelpercore.StringWrapperRuneLoopArgs{
		Content: stringWrapper.content,
		Runes:   &runes,
		Index:   0,
		Rune:    0,
	}

	for args.Index, args.Rune = range runes {
		newStrings[args.Index] = simpleLoopProcessor(&args)
	}

	return &newStrings
}

// Loops through all the rune characters
//
// Slower than direct access
func (stringWrapper *StringWrapper) LoopRunesToGetAnys(
	loopProcessorInterface func(args *strhelpercore.StringWrapperRuneLoopArgs) *interface{},
) *[]*interface{} {
	runes := *stringWrapper.ToRunesPtr()
	results := make([]*interface{}, len(runes))

	args := strhelpercore.StringWrapperRuneLoopArgs{
		Content: stringWrapper.content,
		Runes:   &runes,
		Index:   0,
		Rune:    0,
	}

	for args.Index, args.Rune = range runes {
		results[args.Index] = loopProcessorInterface(&args)
	}

	return &results
}

// Loops through all the rune characters
//
// Slower than direct access
func (stringWrapper *StringWrapper) LoopRunesToGetRunes(
	loopProcessorRune func(args *strhelpercore.StringWrapperRuneLoopArgs) rune,
) *[]rune {
	runes := *stringWrapper.ToRunesPtr()
	newRunes := make([]rune, len(runes))

	args := strhelpercore.StringWrapperRuneLoopArgs{
		Content: stringWrapper.content,
		Runes:   &runes,
		Index:   0,
		Rune:    0,
	}

	for args.Index, args.Rune = range runes {
		newRunes[args.Index] = loopProcessorRune(&args)
	}

	return &newRunes
}

// Loops through all the lines.
func (stringWrapper *StringWrapper) LoopLinesToStringArray(
	lineProcessor strhelpercore.LineProcessor,
) *[]*string {
	allLines := stringWrapper.GetLines()

	return lines.Process(
		stringWrapper.content,
		allLines,
		&lineProcessor)
}

// Loops through all the lines.
func (stringWrapper *StringWrapper) LoopUnixLinesToStringArray(
	lineProcessor strhelpercore.LineProcessor,
) *[]*string {
	allLines := stringWrapper.GetLinesUnix()

	return lines.Process(
		stringWrapper.content,
		allLines,
		&lineProcessor)
}

// Loops through all the lines.
func (stringWrapper *StringWrapper) LoopParallelUnixLinesToStringArray(
	lineProcessor strhelpercore.LineProcessor,
) *[]*string {
	allLines := stringWrapper.GetLinesUnix()

	return lines.ProcessAsync(
		stringWrapper.content,
		allLines,
		&lineProcessor)
}

// Better to use slice or builder for appending lines.
//
// StringWrapper.ValuePtr() + contents joined.
func (stringWrapper *StringWrapper) ConcatWrappers(
	separator string,
	isSkipOnEmpty bool,
	stringWrappers ...StringWrapper,
) *StringWrapper {
	strArray := make([]string, len(stringWrappers))

	for i, wrapper := range stringWrappers {
		strArray[i] = wrapper.Value()
	}

	return stringWrapper.Concatenates(
		separator,
		isSkipOnEmpty,
		&strArray)
}

// Better to use slice or builder for appending lines.
//
// StringWrapper.ValuePtr() + contents joined with separator given.
func (stringWrapper *StringWrapper) ConcatWithSeparator(
	separator string,
	isSkipOnEmpty bool,
	contents ...string,
) *StringWrapper {
	return stringWrapper.Concatenates(
		separator,
		isSkipOnEmpty,
		&contents)
}

// combine current wrapper strings + all given ones with given separator and returns as wrapper
//
// StringWrapper.ValuePtr() + contents joined with separator given.
func (stringWrapper *StringWrapper) Concatenates(
	separator string,
	isSkipOnEmpty bool,
	contents *[]string,
) *StringWrapper {
	combinedResult := concat.StringsArrayWithSeparator(
		stringWrapper.content,
		&separator,
		isSkipOnEmpty,
		contents)

	return NewPtr(&combinedResult)
}

// all given strings with given separator + StringWrapper.Value() content and returns as a wrapper
//
// @isSkipEmptyOrNil:
//  - Skip nil or empty string in elements. (not the whitespace)
//  - If final string compiled string from contents is a whitespace then ignored.
//
// @separator:
//  - used to concat each strings / elements.
//
// @Returns:
//  - @isSkipEmptyOrNil false , (contents joined with separator) + separator + StringWrapper.Value()
//  - @isSkipEmptyOrNil true ,
//    - if not empty or whitespace (StringWrapper.Value()) + allContents join with separator (skips any with nil or "")
//    - if not empty or whitespace (allContents join with separator(skips any with nil or "")) then returns StringWrapper.Value()
//    - if both are not empty and combined @contents is not whitespace then (all @contents combined with separator (skips any with nil or "")) + separator + @StringWrapper.Value()
func (stringWrapper *StringWrapper) Prepends(
	separator string,
	isSkipOnEmpty bool,
	contents *[]string,
) *StringWrapper {
	combinedResult := concat.PrependArrayWithSeparator(
		stringWrapper.content,
		&separator,
		isSkipOnEmpty,
		contents)

	return NewPtr(&combinedResult)
}

func (stringWrapper *StringWrapper) ReplaceWrapper(
	searchingWrapper,
	replacingWrapper *StringWrapper,
	isCaseSensitive bool,
	startsAt int,
	replaceCount int,
) *StringWrapper {
	replacedText := replace.ReplacePtr(
		stringWrapper.content,
		searchingWrapper.ValuePtr(),
		replacingWrapper.ValuePtr(),
		startsAt,
		replaceCount,
		isCaseSensitive,
	)

	return NewPtr(&replacedText)
}

// For better performance use Ptr version.
func (stringWrapper *StringWrapper) Replace(
	search,
	replaceText string,
	isCaseSensitive bool,
	startsAt int,
	replaceCount int,
) string {
	return replace.ReplacePtr(
		stringWrapper.content,
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
	return replace.ReplacePtr(
		stringWrapper.content,
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
	return replace.ReplacePtr(
		stringWrapper.content,
		&search,
		&replaceText,
		startsAt,
		-1,
		isCaseSensitive,
	)
}

func (stringWrapper *StringWrapper) Remove(
	removeString *string,
	isCaseSensitive bool,
	startsAt int,
	count int,
) string {
	return remove.GetPtr(
		stringWrapper.content,
		removeString,
		startsAt,
		count,
		isCaseSensitive,
	)
}

func (stringWrapper *StringWrapper) RemoveAll(
	removeString *string,
	isCaseSensitive bool,
	startsAt int,
) string {
	return remove.GetPtr(
		stringWrapper.content,
		removeString,
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
	return index.OfLastPtr(
		stringWrapper.content,
		&search,
		lastStartIndexReducedBy,
		isCaseSensitive)
}

func (stringWrapper *StringWrapper) LastIndexOfPtr(
	search *string,
	lastStartIndexReducedBy int,
	isCaseSensitive bool,
) int {
	return index.OfLastPtr(
		stringWrapper.content,
		search,
		lastStartIndexReducedBy,
		isCaseSensitive)
}

// For better performance use isstr.IsStartsWithPtr
func (stringWrapper *StringWrapper) IsStartsWith(
	search string,
	isCaseSensitive bool,
	startsAt int,
) bool {
	return isstr.StartsWithPtr(
		stringWrapper.content,
		&search,
		startsAt,
		isCaseSensitive)
}

// Use direct isstr.EndsWithPtr will be faster
func (stringWrapper *StringWrapper) IsEndsWith(
	endsWith string,
	isCaseSensitive bool,
	startsAt int,
) bool {
	return isstr.EndsWithPtr(
		stringWrapper.content,
		&endsWith,
		startsAt,
		isCaseSensitive)
}

// Use direct isstr.EndsWithPtr will be faster
func (stringWrapper *StringWrapper) IsEndsWithPtr(
	endsWith *string,
	isCaseSensitive bool,
	startsAt int,
) bool {
	return isstr.EndsWithPtr(
		stringWrapper.content,
		endsWith,
		startsAt,
		isCaseSensitive)
}

// Returns true if the search text contains any where in the text after the start index.
//
// Use direct isstr.ContainsPtr will be faster
func (stringWrapper *StringWrapper) IsContains(
	search *string,
	isCaseSensitive bool,
	startsAt int,
) bool {
	return index.OfPtr(
		stringWrapper.content,
		search,
		startsAt,
		isCaseSensitive) > strconst.InvalidNotFoundCase
}

// Returns true if the search text contains any where in the text after the start index.
//
// Use direct isstr.ContainsPtr will be faster
func (stringWrapper *StringWrapper) Has(search *string) bool {
	return index.OfPtr(
		stringWrapper.content,
		search,
		0,
		true) > strconst.InvalidNotFoundCase
}

func (stringWrapper *StringWrapper) PadLeftWithSpace(width int) string {
	return padding.SpaceLeft(stringWrapper.content, width)
}

func (stringWrapper *StringWrapper) PadRightWithSpace(width int) string {
	return padding.SpaceRight(stringWrapper.content, width)
}

func (stringWrapper *StringWrapper) PadLeft(width int, paddingStr string) string {
	return padding.Left(stringWrapper.content, &paddingStr, width)
}

func (stringWrapper *StringWrapper) PadRight(width int, paddingStr string) string {
	return padding.Right(stringWrapper.content, &paddingStr, width)
}

func (stringWrapper *StringWrapper) Pad(width int, paddingStr string, isLeft, isRight bool) string {
	return padding.Pad(stringWrapper.content, &paddingStr, width, isLeft, isRight)
}
