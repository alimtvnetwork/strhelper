// Refers to StringWrapperPointer (swp)
//
// Thread safety NOT guaranteed
package swp

import (
	"regexp"
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/chars"
	"gitlab.com/evatix-go/strhelper/concat"
	"gitlab.com/evatix-go/strhelper/index"
	"gitlab.com/evatix-go/strhelper/internal/isinternal"
	"gitlab.com/evatix-go/strhelper/isstr"
	"gitlab.com/evatix-go/strhelper/lines"
	padding2 "gitlab.com/evatix-go/strhelper/padding"
	"gitlab.com/evatix-go/strhelper/remove"
	"gitlab.com/evatix-go/strhelper/replace"
	"gitlab.com/evatix-go/strhelper/reverse"
	"gitlab.com/evatix-go/strhelper/splits"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

// Refers to StringWrapperPointer (swp)
//
// Thread safety NOT guaranteed
//
// Pointer based string wrapper, with cached values
type StringWrapper struct {
	content             *string
	lowerContent        *string
	upperContent        *string
	trimmedSpaceContent *string
	isNullOrEmpty       *bool
	isEmptyOrWhitespace *bool
	uint8s              *[]uint8
	bytes               *[]byte
	runes               *[]rune
	lines               *[]string
	linesUnix           *[]string
	lowerRunes          *[]rune
	upperRunes          *[]rune

	// This represents actual characters length
	runesLength *int
	// len(string) not the actual character size
	lengthInBytes int
}

func (stringWrapper *StringWrapper) Value() *string {
	return stringWrapper.content
}

func (stringWrapper *StringWrapper) ValueWithoutPtr() string {
	return *stringWrapper.content
}

// There is a difference between length in bytes (doesn't represent proper unicode chars) and
//
// length in runes (represents actual char length in unicode format).
//
// If care about unicode chars count then use Length version.
// It returns the len(str) cached version.
// Note :
//  - This version returns cached version of len(str) which is saved during the instantiation of the object creation.
//  - Less expensive (returns StringWrapper.length). Effective for bytes knowledge only.
//  - Doesn't yield accurate characters length of unicode characters but only ascii.
//  - There is a difference between len(str) and utf8.RuneCountInString(str) or len([]rune(str)).
//  - Example : https://play.golang.org/p/78uFF8s-Dw1
func (stringWrapper *StringWrapper) BytesLength() int {
	return stringWrapper.lengthInBytes
}

// Returns len(ToRunesPtr()) cached version. If once runeLength generated then it will not generate again.
//
// If don't care about unicode then use len(str) which is BytesLength
// Note :
//  - A bit expensive to generate. However, this one is cached version.
//  - Yields accurate characters length regardless of unicode.
//  - As there is a difference between len(str) and utf8.RuneCountInString(str) or len([]rune(str))
//  - Example : https://play.golang.org/p/78uFF8s-Dw1
func (stringWrapper *StringWrapper) Length() int {
	if stringWrapper.runesLength == nil {
		runesLength := len(stringWrapper.ToRunes())
		stringWrapper.runesLength = &runesLength
	}

	return *stringWrapper.runesLength
}

func (stringWrapper *StringWrapper) IsEquals(s string, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return s == *stringWrapper.content
	}

	// insensitive
	lower := stringWrapper.ToLowerPtr()

	return *lower == stringWrapper.ToLower()
}

func (stringWrapper *StringWrapper) IsAnyEquals(
	isCaseSensitive bool,
	contents ...*string,
) bool {
	for _, content := range contents {
		if stringWrapper.IsEquals(*content, isCaseSensitive) {
			return true
		}
	}

	return false
}

func (stringWrapper *StringWrapper) IsAllEquals(
	isCaseSensitive bool,
	contents ...*string,
) bool {
	for _, content := range contents {
		if !stringWrapper.IsEquals(*content, isCaseSensitive) {
			return false
		}
	}

	return true
}

// Returns all indexes where findingString is found.
//
// @limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
//
// Results:
//  - Invalid result can be nil if any (content == nil || findingString == nil) results nil.
//  - If no indexes found returns nil
func (stringWrapper *StringWrapper) IndexesOfAll(
	findingString *string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) *[]int {
	return index.OfAllPtr(
		stringWrapper.content,
		findingString,
		startsAtIndex,
		limits,
		isCaseSensitive)
}

// Find all the indexes for all the finding strings given.
//
// limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
func (stringWrapper *StringWrapper) ManyIndexesOfAll(
	findingStrings *[]string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) *strhelpercore.IndexesResultSet {
	return index.OfAllMany(
		stringWrapper.content,
		findingStrings,
		startsAtIndex,
		limits,
		isCaseSensitive)
}

// Returns true based on text compare case sensitive.
func (stringWrapper *StringWrapper) IsSensitiveEquals(s *string) bool {
	if s == nil {
		return false
	}

	return *s == stringWrapper.ValueWithoutPtr()
}

// returns s == nil || len(s) == 0 || s == ""
//
// Thread safety is NOT guaranteed, for parallel programming use swasync.StringWrapper pointer for async mode.
func (stringWrapper *StringWrapper) IsNullOrEmpty() bool {
	if stringWrapper.isEmptyOrWhitespace == nil {
		value := stringWrapper.content
		isEmptyOrNull := value == nil || *value == constants.EmptyString || stringWrapper.lengthInBytes == 0
		stringWrapper.isNullOrEmpty = &isEmptyOrNull
		isEmptyOrWhitespace := isEmptyOrNull

		if isEmptyOrNull == false {
			allRunes := stringWrapper.ToRunesPtr()
			isEmptyOrWhitespace = whitespace.IsRunesWhitespaces(allRunes)
		}

		stringWrapper.isEmptyOrWhitespace = &isEmptyOrWhitespace
	}

	return *stringWrapper.isNullOrEmpty
}

// IsNull(s) || IsNullOrEmpty(s)
func (stringWrapper *StringWrapper) IsNull() bool {
	return stringWrapper.content == nil
}

// IsNullOrEmpty(s) || isstr.BlankPtr(stringWrapper.content)
//
// Thread safety is NOT guaranteed, for parallel programming use swasync.StringWrapper pointer for async mode.
func (stringWrapper *StringWrapper) IsNullOrWhitespace() bool {
	return stringWrapper.IsNullOrEmpty() || *stringWrapper.isEmptyOrWhitespace
}

// Trimmed space string (use caching, doesn't run multiple times)
//
// Thread safety is NOT guaranteed, for parallel programming use swasync.StringWrapper pointer for async mode.
func (stringWrapper *StringWrapper) TrimSpace() *string {
	if stringWrapper.trimmedSpaceContent == nil && !stringWrapper.IsNull() {
		content := stringWrapper.ValueWithoutPtr()
		trimmed := strings.TrimSpace(content)
		stringWrapper.trimmedSpaceContent = &trimmed
	}

	return stringWrapper.trimmedSpaceContent
}

func (stringWrapper *StringWrapper) Trim(cutSet string) *string {
	trimmed := strings.Trim((*stringWrapper).ValueWithoutPtr(), cutSet)

	return &trimmed
}

func (stringWrapper *StringWrapper) TrimLeft(cutSet string) *string {
	trimmed := strings.TrimLeft((*stringWrapper).ValueWithoutPtr(), cutSet)

	return &trimmed
}

func (stringWrapper *StringWrapper) TrimRight(cutSet string) *string {
	trimmed := strings.TrimRight((*stringWrapper).ValueWithoutPtr(), cutSet)

	return &trimmed
}

// GetLines splits by newline of os
//
// Windows (`\r\n`), unix (`\n`) - darwin/macos/linux
//
// Thread safety is NOT guaranteed, for parallel programming use swasync.StringWrapper pointer for async mode.
func (stringWrapper *StringWrapper) GetLines() *[]string {
	if stringWrapper.lines == nil && !stringWrapper.IsNull() {
		stringWrapper.lines = lines.GetPtr(stringWrapper.content)
	}

	return stringWrapper.lines
}

// GetLines splits by newline using unix Split `\n`
//
// Thread safety is NOT guaranteed, for parallel programming use swasync.StringWrapper pointer for async mode.
func (stringWrapper *StringWrapper) GetUnixLines() *[]string {
	isRequiresSetting := stringWrapper.linesUnix == nil &&
		!stringWrapper.IsNull()
	isNewLineSameAsUnix := isinternal.IsCurrentOsUnix()

	if isRequiresSetting && isNewLineSameAsUnix {
		// same no need to process
		stringWrapper.linesUnix = stringWrapper.GetLines()
	}

	if isRequiresSetting && !isNewLineSameAsUnix {
		// requires processing
		linesUnix := lines.UnixGet(stringWrapper.content)
		stringWrapper.linesUnix = &linesUnix
	}

	return stringWrapper.linesUnix
}

// GetLinesAsWrappers splits by newline of os
//
// Windows (`\r\n`), unix (`\n`) - darwin/macos/linux
func (stringWrapper *StringWrapper) GetLinesAsWrappers() *[]*StringWrapper {
	allLines := stringWrapper.GetLines()
	length := len(*allLines)
	wrappers := make([]*StringWrapper, length)

	for i := 0; i < length; i++ {
		wrappers[i] = New(&(*allLines)[i])
	}

	return &wrappers
}

// GetUnixLinesAsWrappers splits by newline using unix Split `\n`
func (stringWrapper *StringWrapper) GetUnixLinesAsWrappers() *[]*StringWrapper {
	allLines := stringWrapper.GetUnixLines()
	length := len(*allLines)
	wrappers := make([]*StringWrapper, length)

	for i := 0; i < length; i++ {
		wrappers[i] = New(&(*allLines)[i])
	}

	return &wrappers
}

// get uint8 array
func (stringWrapper *StringWrapper) ToUInt8s() []uint8 {
	return *stringWrapper.ToUInt8sPtr()
}

// get uint8 array ptr
//
// Thread safety is NOT guaranteed, for parallel programming use swasync.StringWrapper pointer for async mode.
func (stringWrapper *StringWrapper) ToUInt8sPtr() *[]uint8 {
	if (stringWrapper.uint8s == nil || *stringWrapper.uint8s == nil) && !stringWrapper.IsNull() {
		*stringWrapper.uint8s = []uint8(*stringWrapper.content)
	}

	return stringWrapper.uint8s
}

// get bytes array
func (stringWrapper *StringWrapper) ToBytes() []byte {
	return *stringWrapper.ToBytesPtr()
}

// get bytes array ptr
//
// Thread safety is NOT guaranteed, for parallel programming use swasync.StringWrapper pointer for async mode.
func (stringWrapper *StringWrapper) ToBytesPtr() *[]byte {
	if stringWrapper.bytes == nil || *stringWrapper.bytes == nil {
		allBytes := []byte(*stringWrapper.content)

		stringWrapper.bytes = &allBytes
	}

	return stringWrapper.bytes
}

// get rune array
func (stringWrapper *StringWrapper) ToRunes() []rune {
	return *stringWrapper.ToRunesPtr()
}

// get rune array ptr
//
// Thread safety is NOT guaranteed, for parallel programming use swasync.StringWrapper pointer for async mode.
func (stringWrapper *StringWrapper) ToRunesPtr() *[]rune {
	if (stringWrapper.runes == nil || *stringWrapper.runes == nil) && !stringWrapper.IsNull() {
		runes := []rune(*stringWrapper.content)

		stringWrapper.runes = &runes
	}

	return stringWrapper.runes
}

// get lowercase rune array ptr
//
// Thread safety is NOT guaranteed, for parallel programming use swasync.StringWrapper pointer for async mode.
func (stringWrapper *StringWrapper) ToLowerRunesPtr() *[]rune {
	if (stringWrapper.lowerRunes == nil || *stringWrapper.lowerRunes == nil) && !stringWrapper.IsNull() {
		lowerRunes := chars.ToLowerRunes(stringWrapper.ToRunesPtr())
		stringWrapper.lowerRunes = lowerRunes
	}

	return stringWrapper.lowerRunes
}

// get uppercase rune array ptr
//
// Thread safety is NOT guaranteed, for parallel programming use swasync.StringWrapper pointer for async mode.
func (stringWrapper *StringWrapper) ToUpperRunesPtr() *[]rune {
	if stringWrapper.upperRunes == nil || *stringWrapper.upperRunes == nil {
		upperRunes := chars.ToUpperRunes(stringWrapper.ToRunesPtr())
		stringWrapper.upperRunes = upperRunes
	}

	return stringWrapper.upperRunes
}

// returns true if IsNullOrWhitespace(s)
func (stringWrapper *StringWrapper) IsBlank() bool {
	return stringWrapper.IsNullOrWhitespace()
}

// Has at least one character other than space or whitespace
func (stringWrapper *StringWrapper) HasCharacter() bool {
	return !stringWrapper.IsNullOrWhitespace()
}

// Has at least one character other than space or whitespace
func (stringWrapper *StringWrapper) IsDefined() bool {
	return !stringWrapper.IsNullOrWhitespace()
}

func (stringWrapper *StringWrapper) ToLower() string {
	return *stringWrapper.ToLowerPtr()
}

// Return lower string (used caching)
//
// Thread safety is NOT guaranteed, for parallel programming use swasync.StringWrapper pointer for async mode.
func (stringWrapper *StringWrapper) ToLowerPtr() *string {
	if stringWrapper.lowerContent == nil && !stringWrapper.IsNull() {
		lowerCase := string(*stringWrapper.ToLowerRunesPtr())
		stringWrapper.lowerContent = &lowerCase
	}

	return stringWrapper.lowerContent
}

func (stringWrapper *StringWrapper) ToUpper() string {
	return *stringWrapper.ToUpperPtr()
}

// Return uppercase string (used caching)
//
// Thread safety is NOT guaranteed, for parallel programming use swasync.StringWrapper pointer for async mode.
func (stringWrapper *StringWrapper) ToUpperPtr() *string {
	if stringWrapper.upperContent == nil && !stringWrapper.IsNullOrEmpty() {
		upperContent := string(*stringWrapper.ToUpperRunesPtr())
		stringWrapper.upperContent = &upperContent
	}

	return stringWrapper.upperContent
}

func (stringWrapper *StringWrapper) ToLowerWrapperPtr() *StringWrapper {
	return New(stringWrapper.ToLowerPtr())
}

func (stringWrapper *StringWrapper) ToLowerWrapper() StringWrapper {
	return *stringWrapper.ToLowerWrapperPtr()
}

// Returns strings to upper case as StringWrapper
func (stringWrapper *StringWrapper) ToUpperWrapper() StringWrapper {
	return *stringWrapper.ToUpperWrapperPtr()
}

// Returns strings to upper case as StringWrapper
func (stringWrapper *StringWrapper) ToUpperWrapperPtr() *StringWrapper {
	return New(stringWrapper.ToUpperPtr())
}

// Returns character at the given index, if not exist then panic.
//
// Slower than direct access
//
// Use loop version if want to loop through
func (stringWrapper *StringWrapper) At(index int) uint8 {
	return stringWrapper.ToUInt8s()[index]
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
	allLines := stringWrapper.GetUnixLines()

	return lines.Process(
		stringWrapper.content,
		allLines,
		&lineProcessor)
}

// Loops through all the lines.
func (stringWrapper *StringWrapper) LoopParallelUnixLinesToStringArray(
	lineProcessor strhelpercore.LineProcessor,
) *[]*string {
	allLines := stringWrapper.GetUnixLines()

	return lines.ProcessAsync(
		stringWrapper.content,
		allLines,
		&lineProcessor)
}

// Create regular expression from current string.
// Recommendation, do not create regular expressions inside a function call, keep it on top of the file as variables
func (stringWrapper *StringWrapper) CreateRegularExpression() *regexp.Regexp {
	return regexp.MustCompile(*stringWrapper.content)
}

// Same as Value()
func (stringWrapper *StringWrapper) String() string {
	return stringWrapper.ValueWithoutPtr()
}

// returns -1 if the index is not present in strings lengthInBytes.
// or else returns the character value from that index
// performance should be very slow, use direct access of str.
func (stringWrapper *StringWrapper) GetSafeIndexAt(index int) int16 {
	if !stringWrapper.HasIndex(index) {
		return constants.InvalidNotFoundCase
	}

	return int16(stringWrapper.ValueWithoutPtr()[index])
}

// returns -1 if the index is not present in strings lengthInBytes.
// or else returns the character value from that index
// performance should be very slow, use direct access of stringWrapper.ToRunesPtr().
func (stringWrapper *StringWrapper) GetSafeRuneIndexAt(index int) rune {
	if !stringWrapper.HasIndex(index) {
		return constants.InvalidNotFoundCase
	}

	return stringWrapper.ToRunes()[index]
}

func (stringWrapper *StringWrapper) IsEqualAtIndex(
	index int,
	char uint8,
	isCaseSensitive bool,
) bool {
	valueAt := stringWrapper.ValueWithoutPtr()[index]

	if isCaseSensitive {
		return valueAt == char
	}

	return chars.IsMatchCaseInsensitive(valueAt, char)
}

// stringWrapper.BytesLength()-1 >= index
func (stringWrapper *StringWrapper) HasIndex(index int) bool {
	return stringWrapper.BytesLength()-1 >= index
}

// stringWrapper.Length()-1 >= index
func (stringWrapper *StringWrapper) HasRuneIndex(index int) bool {
	return stringWrapper.Length()-1 >= index
}

// Returns a new string builder contains text of stringWrapper and has a
// growth = stringWrapper.lengthInBytes + additionalGrowLength
func (stringWrapper *StringWrapper) Builder(additionalGrowLength int) strings.Builder {
	builder := strings.Builder{}
	currentString := stringWrapper.content
	length := stringWrapper.BytesLength() + additionalGrowLength
	builder.Grow(length)
	builder.WriteString(*currentString)
	return builder
}

// Returns a new string builder contains text of stringWrapper + str and has a
// growth = stringWrapper.lengthInBytes + additionalGrowLength + len(str)
func (stringWrapper *StringWrapper) BuilderWithStr(str *string, additionalGrowLength int) strings.Builder {
	builder := strings.Builder{}
	currentString := stringWrapper.content
	length := stringWrapper.BytesLength() + len(*str) + additionalGrowLength
	builder.Grow(length)
	builder.WriteString(*currentString)
	builder.WriteString(*str)

	return builder
}

// Add the contents before the content of StringWrapper.Value()
func (stringWrapper *StringWrapper) Prepend(contents ...string) StringWrapper {
	return *stringWrapper.Prepends(
		constants.EmptyString,
		false, // must add everything
		&contents)
}

// Line is the separator for add the content before the content of StringWrapper.Value()
func (stringWrapper *StringWrapper) PrependLines(contents ...string) StringWrapper {
	return *stringWrapper.Prepends(
		constants.NewLine,
		false, // must add everything
		&contents)
}

// Prepends array @contents before @StringWrapper.Value()
// (@combinedContents + *separator + @StringWrapper.Value())
// and compiles to a single string using @separator.
//
// @Expression:
//  - @combinedContents = contents joined to single one using separator (skip empty if flag is enabled)
//  - returns @combinedContents + *separator + @StringWrapper.Value()
//
// @isSkipEmptyOrNil:
//  - Skip nil or empty string in elements. (not the whitespace)
//  - If final string compiled string from contents is a whitespace then ignored.
//
// @separator:
//  - used to concat each strings / elements.
//
// @Returns:
//  - @isSkipEmptyOrNil false , @combinedContents + separator + @StringWrapper.Value()
//  - @isSkipEmptyOrNil true ,
//    - if not empty or whitespace (@StringWrapper.Value()) then @combinedContents
//    - if not empty or whitespace (@combinedContents) then returns @StringWrapper.Value()
//    - if both are not empty and combined contents is not whitespace then
//          returns (@combinedContents) + separator + @StringWrapper.Value()
func (stringWrapper *StringWrapper) Prepends(
	separator string,
	isSkipOnEmpty bool,
	contents *[]string,
) *StringWrapper {
	combinedResult := concat.PrependArrayWithCurrentStringUsingSeparator(
		stringWrapper.content,
		&separator,
		isSkipOnEmpty,
		contents)

	return New(&combinedResult)
}

// Prepends array @contents before
// @StringWrapper.Value() (@combinedContents + *separator + @StringWrapper.Value())
// and compiles to a single string using @separator.
//
// @Expression:
//  - @combinedContents = contents joined to single one using separator (skip empty if flag is enabled)
//  - returns @combinedContents + *separator + @StringWrapper.Value()
//
// @isSkipEmptyOrNil:
//  - Skip nil or empty string in elements. (not the whitespace)
//  - If final string compiled string from contents is a whitespace then ignored.
//
// @separator:
//  - used to concat each strings / elements.
//
// @Returns:
//  - @isSkipEmptyOrNil false , @combinedContents + separator + @StringWrapper.Value()
//  - @isSkipEmptyOrNil true ,
//    - if not empty or whitespace (@StringWrapper.Value()) then returns @combinedContents
//    - if not empty or whitespace (@combinedContents) then returns @StringWrapper.Value()
//    - if both are not empty and combined contents is not whitespace then
//          returns (@combinedContents) + separator + @StringWrapper.Value()
func (stringWrapper *StringWrapper) PrependAsString(
	separator string,
	isSkipOnEmpty bool,
	contents *[]string,
) *string {
	combinedResult := concat.PrependArrayWithCurrentStringUsingSeparator(
		stringWrapper.content,
		&separator,
		isSkipOnEmpty,
		contents)

	return &combinedResult
}

// Better to use slice or builder for appending lines in a loop.
//
// stringWrapper.content + separator + JoinAll(separator, stringWrappers)
func (stringWrapper *StringWrapper) AppendLines(isSkipOnEmpty bool, contents ...string) *StringWrapper {
	return stringWrapper.Concatenates(constants.NewLine, isSkipOnEmpty, &contents)
}

// Better to use slice or builder for appending or concatenating lines in a loop.
//
// stringWrapper.content + separator + JoinAll(separator, stringWrappers)
func (stringWrapper *StringWrapper) Concat(contents ...string) *StringWrapper {
	return stringWrapper.Concatenates(
		constants.EmptyString,
		false, // must add everything
		&contents)
}

// Better to use slice or builder for appending lines.
//
// stringWrapper.content + separator + JoinAll(separator, stringWrappers)
func (stringWrapper *StringWrapper) ConcatWrappers(
	separator string,
	isSkipOnEmpty bool,
	stringWrappers ...StringWrapper,
) *StringWrapper {
	strArray := make([]*string, len(stringWrappers))

	for i, wrapper := range stringWrappers {
		strArray[i] = wrapper.content
	}

	return stringWrapper.ConcatPtrStr(
		separator,
		isSkipOnEmpty,
		&strArray)
}

// Better to use slice or builder for appending lines.
//
// stringWrapper.content + separator + JoinAll(separator, stringWrappers)
func (stringWrapper *StringWrapper) ConcatWrappersPointers(
	separator string,
	isSkipOnEmpty bool,
	stringWrappers ...*StringWrapper,
) *StringWrapper {
	strArray := make([]*string, len(stringWrappers))

	for i, wrapper := range stringWrappers {
		strArray[i] = wrapper.content
	}

	return stringWrapper.ConcatPtrStr(
		separator,
		isSkipOnEmpty,
		&strArray)
}

// Better to use slice or builder for appending lines.
//
// stringWrapper.content + separator + JoinAll(separator, stringWrappers)
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

// Better to use slice or builder for appending lines.
//
// stringWrapper.content + separator + JoinAll(separator, stringWrappers)
func (stringWrapper *StringWrapper) ConcatStrPointersWithSeparator(
	separator string,
	isSkipOnEmpty bool,
	contents ...*string,
) *StringWrapper {
	return stringWrapper.ConcatPtrStr(
		separator,
		isSkipOnEmpty,
		&contents)
}

// combine current wrapper strings + all given ones with given separator and returns as wrapper
//
// stringWrapper.content + separator + JoinAll(separator, stringWrappers)
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

	return New(&combinedResult)
}

// combine current wrapper strings + all given ones with given separator and returns as wrapper
func (stringWrapper *StringWrapper) ConcatPtrStr(
	separator string,
	isSkipOnEmpty bool,
	contents *[]*string,
) *StringWrapper {
	combinedResult := concat.PtrStringsArrayWithSeparator(
		stringWrapper.content,
		&separator,
		isSkipOnEmpty,
		contents)

	return New(&combinedResult)
}

func (stringWrapper *StringWrapper) ReplaceWrapper(
	searchingWrapper,
	replacingWrapper *StringWrapper,
	isCaseSensitive bool,
	startsAt int,
	replaceCount int,
) *StringWrapper {
	replacedText := replace.GetPtr(
		stringWrapper.content,
		searchingWrapper.content,
		replacingWrapper.content,
		startsAt,
		replaceCount,
		isCaseSensitive,
	)

	return New(&replacedText)
}

// For better performance use Ptr version.
func (stringWrapper *StringWrapper) Replace(
	search,
	replaceText string,
	isCaseSensitive bool,
	startsAt int,
	replaceCount int,
) string {
	return replace.GetPtr(
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
) *string {
	replacedText := replace.GetPtr(
		stringWrapper.content,
		search,
		replaceText,
		startsAt,
		replaceCount,
		isCaseSensitive,
	)

	return &replacedText
}

func (stringWrapper *StringWrapper) ReplaceAll(
	search,
	replaceText string,
	isCaseSensitive bool,
	startsAt int,
) string {
	return replace.GetPtr(
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

func (stringWrapper *StringWrapper) ReplaceMultiple(
	searchReplaceMap *map[string]string,
	startsAt int,
) string {
	return replace.ManyPtr(
		stringWrapper.content,
		searchReplaceMap,
		startsAt,
		-1,
		true,
	)
}

func (stringWrapper *StringWrapper) ReplaceMultipleCase(
	searchReplaceMap *map[string]string,
	startsAt int,
	limits int,
	isCaseSensitive bool,
) string {
	return replace.ManyPtr(
		stringWrapper.content,
		searchReplaceMap,
		startsAt,
		limits,
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

func (stringWrapper *StringWrapper) ReversePtr() string {
	return reverse.Ptr(stringWrapper.content)
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

// For better performance use isstr.StartsWithPtr
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

func (stringWrapper *StringWrapper) PadLeftWithSpace(width int) string {
	return padding2.SpaceLeft(stringWrapper.content, width)
}

func (stringWrapper *StringWrapper) PadRightWithSpace(width int) string {
	return padding2.SpaceRight(stringWrapper.content, width)
}

func (stringWrapper *StringWrapper) PadLeft(width int, padding string) string {
	return padding2.Left(stringWrapper.content, &padding, width)
}

func (stringWrapper *StringWrapper) PadRight(width int, padding string) string {
	return padding2.Right(stringWrapper.content, &padding, width)
}

func (stringWrapper *StringWrapper) Pad(width int, padding string, isLeft, isRight bool) string {
	return padding2.Pad(stringWrapper.content, &padding, width, isLeft, isRight)
}

// Multiple split occur from the given array of splits.
//
// Basics of split("Hello World", " ") -> ["Hello", "World"] splitter will not be available in the result.
//
// limit :
//  - number of times split will performed for all
//  - if -1 then all split will occur
//
// splitStartsAt:
//  - where split searching will start from.
func (stringWrapper *StringWrapper) MultiSplit(
	startsAt, limits int,
	splitters ...string,
) *strhelpercore.SplitResultOverview {
	return splits.ManyPtr(
		stringWrapper.content,
		&splitters,
		startsAt,
		limits,
		true)
}

// Multiple split occur from the given array of splits.
//
// Basics of split("Hello World", " ") -> ["Hello", "World"] splitter will not be available in the result.
//
// limit :
//  - number of times split will performed for all
//  - if -1 then all split will occur
//
// splitStartsAt:
//  - where split searching will start from.
func (stringWrapper *StringWrapper) MultiSplitCase(
	startsAt, limits int,
	isCaseSensitive bool,
	splitters ...string,
) *strhelpercore.SplitResultOverview {
	return splits.ManyPtr(
		stringWrapper.content,
		&splitters,
		startsAt,
		limits,
		isCaseSensitive)
}

func (stringWrapper *StringWrapper) Split(splitter string) []string {
	return strings.Split(
		*stringWrapper.content,
		splitter)
}

func (stringWrapper *StringWrapper) SplitN(splitter string, count int) []string {
	return strings.SplitN(
		*stringWrapper.content,
		splitter,
		count)
}
