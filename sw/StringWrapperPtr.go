package sw

import (
	"regexp"
	"strings"
	"sync"

	"gitlab.com/evatix-go/strhelper"
	"gitlab.com/evatix-go/strhelper/charhelper"
	"gitlab.com/evatix-go/strhelper/concat"
	"gitlab.com/evatix-go/strhelper/constants"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

type StringWrapperPtr struct {
	content             *string
	lowerContent        *string
	upperContent        *string
	trimmedSpaceContent *string
	// len(string) not the actual character size
	lengthInBytes       int
	isNullOrEmpty       *bool
	isEmptyOrWhitespace *bool
	uint8s              *[]uint8
	bytes               *[]byte
	runes               *[]rune
	lines               *[]string
	linesUnix           *[]string
	lowerRunes          *[]rune
	upperRunes          *[]rune
	sync.Mutex

	// This represents actual characters length
	runesLength *int
}

func NewSwPtr(stringInput *string) *StringWrapperPtr {
	return &StringWrapperPtr{
		content:             stringInput,
		trimmedSpaceContent: nil,
		lengthInBytes:       len(*stringInput),
		isNullOrEmpty:       nil,
		isEmptyOrWhitespace: nil,
		uint8s:              nil,
		bytes:               nil,
		runes:               nil,
		runesLength:         nil,
		lowerRunes:          nil,
		upperRunes:          nil,
		Mutex:               sync.Mutex{},
	}
}

func (stringWrapperPtr *StringWrapperPtr) Lock() {
	stringWrapperPtr.Mutex.Lock()
}

func (stringWrapperPtr *StringWrapperPtr) Unlock() {
	stringWrapperPtr.Mutex.Unlock()
}

func (stringWrapperPtr *StringWrapperPtr) Value() *string {
	return stringWrapperPtr.content
}

func (stringWrapperPtr *StringWrapperPtr) ValueWithoutPtr() string {
	return *stringWrapperPtr.content
}

// As there is a difference between len(str) and utf8.RuneCountInString(str)
//
// It returns the len(str) cached version.
func (stringWrapperPtr *StringWrapperPtr) LengthInBytes() int {
	return stringWrapperPtr.lengthInBytes
}

// As there is a difference between len(str) and utf8.RuneCountInString(str) or len([]rune(str))
//
// It returns the len([]rune(str)) or len([]runeCached) cached version.
func (stringWrapperPtr *StringWrapperPtr) Length() int {
	stringWrapperPtr.Lock()
	defer stringWrapperPtr.Unlock()

	if stringWrapperPtr.runesLength == nil {
		runesLength := len(stringWrapperPtr.ToRunes())
		(*stringWrapperPtr).runesLength = &runesLength
	}

	return *stringWrapperPtr.runesLength
}

func (stringWrapperPtr *StringWrapperPtr) IsEquals(s string, isCaseSensitive bool) bool {
	if isCaseSensitive {
		return s == *stringWrapperPtr.content
	}

	// insensitive
	lower := strings.ToLower(s)

	return lower == stringWrapperPtr.ToLower()
}

func (stringWrapperPtr *StringWrapperPtr) IsAnyEquals(
	isCaseSensitive bool,
	contents ...*string,
) bool {
	for _, content := range contents {
		if (*stringWrapperPtr).IsEquals(*content, isCaseSensitive) {
			return true
		}
	}

	return false
}

func (stringWrapperPtr *StringWrapperPtr) IsAllEquals(
	isCaseSensitive bool,
	contents ...*string,
) bool {
	for _, content := range contents {
		if !(*stringWrapperPtr).IsEquals(*content, isCaseSensitive) {
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
func (stringWrapperPtr *StringWrapperPtr) IndexesOfAll(
	findingString *string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) []int {
	return strhelper.IndexesOfAllPtr(
		stringWrapperPtr.content,
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
func (stringWrapperPtr *StringWrapperPtr) MultipleStringIndexesOfAll(
	findingStrings *[]string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) *strhelpercore.IndexesResultSet {
	return strhelper.MultiStrIndexesOfAllUsingSimpleArrayPtr(
		stringWrapperPtr.content,
		findingStrings,
		startsAtIndex,
		limits,
		isCaseSensitive)
}

// Returns true based on text compare case sensitive.
func (stringWrapperPtr *StringWrapperPtr) IsSensitiveEquals(s *string) bool {
	if s == nil {
		return false
	}

	return *s == stringWrapperPtr.ValueWithoutPtr()
}

// returns s == nil || len(s) == 0 || s == ""
func (stringWrapperPtr *StringWrapperPtr) IsNullOrEmpty() bool {
	stringWrapperPtr.Lock()
	defer stringWrapperPtr.Unlock()

	if stringWrapperPtr.isEmptyOrWhitespace == nil {
		value := stringWrapperPtr.content
		isEmptyOrNull := value == nil || *value == constants.EmptyString || (*stringWrapperPtr).lengthInBytes == 0
		stringWrapperPtr.isNullOrEmpty = &isEmptyOrNull
		isEmptyOrWhitespace := isEmptyOrNull || whitespace.IsWhitespaceOnly(value)
		(*stringWrapperPtr).isEmptyOrWhitespace = &isEmptyOrWhitespace
	}

	return *(*stringWrapperPtr).isNullOrEmpty
}

// IsNull(s) || IsNullOrEmpty(s)
func (stringWrapperPtr *StringWrapperPtr) IsNull() bool {
	return (*stringWrapperPtr).content == nil
}

// IsNullOrEmpty(s) || strhelper.IsBlankPtr(stringWrapperPtr.content)
func (stringWrapperPtr *StringWrapperPtr) IsNullOrWhitespace() bool {
	return (*stringWrapperPtr).IsNullOrEmpty() || *(*stringWrapperPtr).isEmptyOrWhitespace
}

func (stringWrapperPtr *StringWrapperPtr) TrimSpace() *string {
	stringWrapperPtr.Lock()
	defer stringWrapperPtr.Unlock()

	if stringWrapperPtr.trimmedSpaceContent == nil && !(*stringWrapperPtr).IsNull() {
		content := stringWrapperPtr.ValueWithoutPtr()
		trimmed := strings.TrimSpace(content)
		stringWrapperPtr.trimmedSpaceContent = &trimmed
	}

	return stringWrapperPtr.trimmedSpaceContent
}

func (stringWrapperPtr *StringWrapperPtr) Trim(cutSet string) *string {
	trimmed := strings.Trim((*stringWrapperPtr).ValueWithoutPtr(), cutSet)

	return &trimmed
}

func (stringWrapperPtr *StringWrapperPtr) TrimLeft(cutSet string) *string {
	trimmed := strings.TrimLeft((*stringWrapperPtr).ValueWithoutPtr(), cutSet)

	return &trimmed
}

func (stringWrapperPtr *StringWrapperPtr) TrimRight(cutSet string) *string {
	trimmed := strings.TrimRight((*stringWrapperPtr).ValueWithoutPtr(), cutSet)

	return &trimmed
}

// GetLines splitted by newline of os
//
// Windows (`\r\n`), unix (`\n`) - darwin/macos/linux
func (stringWrapperPtr *StringWrapperPtr) GetLines() *[]string {
	stringWrapperPtr.Lock()
	defer stringWrapperPtr.Unlock()

	if stringWrapperPtr.lines == nil && !(*stringWrapperPtr).IsNull() {
		lines := strhelper.GetLines(stringWrapperPtr.content)
		(*stringWrapperPtr).lines = &lines
	}

	return stringWrapperPtr.lines
}

// GetLines splitted by newline using unix Split `\n`
func (stringWrapperPtr *StringWrapperPtr) GetLinesUnix() *[]string {
	stringWrapperPtr.Lock()
	defer stringWrapperPtr.Unlock()

	isRequiresSetting := stringWrapperPtr.linesUnix == nil &&
		!(*stringWrapperPtr).IsNull()
	isNewLineSameAsUnix := constants.NewLine == constants.NewLineUnix

	if isRequiresSetting && isNewLineSameAsUnix {
		// same no need to process
		stringWrapperPtr.linesUnix = stringWrapperPtr.lines
	}

	if isRequiresSetting && !isNewLineSameAsUnix {
		// requires processing
		linesUnix := strhelper.GetLinesUnix(stringWrapperPtr.content)
		(*stringWrapperPtr).linesUnix = &linesUnix
	}

	return stringWrapperPtr.linesUnix
}

// get uint8 array
func (stringWrapperPtr *StringWrapperPtr) ToUInt8s() []uint8 {
	return *stringWrapperPtr.ToUInt8sPtr()
}

// get uint8 array ptr
func (stringWrapperPtr *StringWrapperPtr) ToUInt8sPtr() *[]uint8 {
	stringWrapperPtr.Lock()
	defer stringWrapperPtr.Unlock()

	if (stringWrapperPtr.uint8s == nil || *stringWrapperPtr.uint8s == nil) && !(*stringWrapperPtr).IsNull() {
		*stringWrapperPtr.uint8s = []uint8(*stringWrapperPtr.content)
	}

	return stringWrapperPtr.uint8s
}

// get bytes array
func (stringWrapperPtr *StringWrapperPtr) ToBytes() []byte {
	return *stringWrapperPtr.ToBytesPtr()
}

// get bytes array ptr
func (stringWrapperPtr *StringWrapperPtr) ToBytesPtr() *[]byte {
	stringWrapperPtr.Lock()
	defer stringWrapperPtr.Unlock()

	if stringWrapperPtr.bytes == nil || *stringWrapperPtr.bytes == nil {
		*stringWrapperPtr.bytes = []byte(*stringWrapperPtr.content)
	}

	return stringWrapperPtr.bytes
}

// get rune array
func (stringWrapperPtr *StringWrapperPtr) ToRunes() []rune {
	return *stringWrapperPtr.ToRunesPtr()
}

// get rune array ptr
func (stringWrapperPtr *StringWrapperPtr) ToRunesPtr() *[]rune {
	stringWrapperPtr.Lock()
	defer stringWrapperPtr.Unlock()

	if (stringWrapperPtr.runes == nil || *stringWrapperPtr.runes == nil) && !(*stringWrapperPtr).IsNull() {
		*stringWrapperPtr.runes = []rune(*stringWrapperPtr.content)
	}

	return stringWrapperPtr.runes
}

// get lowercase rune array ptr
func (stringWrapperPtr *StringWrapperPtr) ToLowerRunesPtr() *[]rune {
	stringWrapperPtr.Lock()
	defer stringWrapperPtr.Unlock()

	if (stringWrapperPtr.lowerRunes == nil || *stringWrapperPtr.lowerRunes == nil) && !(*stringWrapperPtr).IsNull() {
		lowerRunes := charhelper.ToLowerRunes(stringWrapperPtr.ToRunesPtr())
		stringWrapperPtr.lowerRunes = lowerRunes
	}

	return stringWrapperPtr.lowerRunes
}

// get uppercase rune array ptr
func (stringWrapperPtr *StringWrapperPtr) ToUpperRunesPtr() *[]rune {
	stringWrapperPtr.Lock()
	defer stringWrapperPtr.Unlock()

	if stringWrapperPtr.upperRunes == nil || *stringWrapperPtr.upperRunes == nil {
		upperRunes := charhelper.ToUpperRunes(stringWrapperPtr.ToRunesPtr())
		stringWrapperPtr.upperRunes = upperRunes
	}

	return stringWrapperPtr.upperRunes
}

// returns true if IsNullOrWhitespace(s)
func (stringWrapperPtr *StringWrapperPtr) IsBlank() bool {
	return (*stringWrapperPtr).IsNullOrWhitespace()
}

// Has at least one character other than space or whitespace
func (stringWrapperPtr *StringWrapperPtr) HasCharacter() bool {
	return !(*stringWrapperPtr).IsNullOrWhitespace()
}

// Has at least one character other than space or whitespace
func (stringWrapperPtr *StringWrapperPtr) IsDefined() bool {
	return !(*stringWrapperPtr).IsNullOrWhitespace()
}

func (stringWrapperPtr *StringWrapperPtr) ToLower() string {
	return *stringWrapperPtr.ToLowerPtr()
}

func (stringWrapperPtr *StringWrapperPtr) ToLowerPtr() *string {
	stringWrapperPtr.Lock()
	defer stringWrapperPtr.Unlock()

	if stringWrapperPtr.lowerContent == nil && !(*stringWrapperPtr).IsNull() {
		lowerCase := string(*stringWrapperPtr.ToLowerRunesPtr())
		stringWrapperPtr.lowerContent = &lowerCase
	}

	return stringWrapperPtr.lowerContent
}

func (stringWrapperPtr *StringWrapperPtr) ToUpper() string {
	return *stringWrapperPtr.ToUpperPtr()
}

func (stringWrapperPtr *StringWrapperPtr) ToUpperPtr() *string {
	stringWrapperPtr.Lock()
	defer stringWrapperPtr.Unlock()

	if stringWrapperPtr.upperContent == nil && !(*stringWrapperPtr).IsNullOrEmpty() {
		upperContent := string(*stringWrapperPtr.ToUpperRunesPtr())
		stringWrapperPtr.upperContent = &upperContent
	}

	return stringWrapperPtr.upperContent
}

func (stringWrapperPtr *StringWrapperPtr) ToLowerWrapperPtr() *StringWrapperPtr {
	return NewSwPtr(stringWrapperPtr.ToLowerPtr())
}

func (stringWrapperPtr *StringWrapperPtr) ToLowerWrapper() StringWrapperPtr {
	return *stringWrapperPtr.ToLowerWrapperPtr()
}

// Returns strings to upper case as StringWrapperPtr
func (stringWrapperPtr *StringWrapperPtr) ToUpperWrapper() StringWrapperPtr {
	return *stringWrapperPtr.ToUpperWrapperPtr()
}

// Returns strings to upper case as StringWrapperPtr
func (stringWrapperPtr *StringWrapperPtr) ToUpperWrapperPtr() *StringWrapperPtr {
	return NewSwPtr(stringWrapperPtr.ToUpperPtr())
}

// Returns character at the given index, if not exist then panic.
//
// Slower than direct access
//
// Use loop version if want to loop through
func (stringWrapperPtr *StringWrapperPtr) At(index int) uint8 {
	return stringWrapperPtr.ToUInt8s()[index]
}

// Loops through all the rune characters
//
// Slower than direct access
func (stringWrapperPtr *StringWrapperPtr) LoopRunes(
	simpleLoopProcessor func(args *strhelpercore.StringWrapperRuneLoopArgs) *string,
) *[]*string {
	runes := *stringWrapperPtr.ToRunesPtr()
	newStrings := make([]*string, len(runes))

	args := strhelpercore.StringWrapperRuneLoopArgs{
		Content: stringWrapperPtr.content,
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
func (stringWrapperPtr *StringWrapperPtr) LoopRunesToGetAnys(
	loopProcessorInterface func(args *strhelpercore.StringWrapperRuneLoopArgs) *interface{},
) *[]*interface{} {
	runes := *stringWrapperPtr.ToRunesPtr()
	results := make([]*interface{}, len(runes))

	args := strhelpercore.StringWrapperRuneLoopArgs{
		Content: stringWrapperPtr.content,
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
func (stringWrapperPtr *StringWrapperPtr) LoopRunesToGetRunes(
	loopProcessorRune func(args *strhelpercore.StringWrapperRuneLoopArgs) rune,
) *[]rune {
	runes := *stringWrapperPtr.ToRunesPtr()
	newRunes := make([]rune, len(runes))

	args := strhelpercore.StringWrapperRuneLoopArgs{
		Content: stringWrapperPtr.content,
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
func (stringWrapperPtr *StringWrapperPtr) LoopLinesToStringArray(
	loopProcessorRune func(args *strhelpercore.StringWrapperLineLoopArgs) *string,
) *[]*string {
	lines := stringWrapperPtr.GetLines()
	newStrings := make([]*string, len(*lines))

	args := strhelpercore.StringWrapperLineLoopArgs{
		Content: stringWrapperPtr.content,
		Lines:   lines,
		Index:   -1,                    // it will change per line
		Line:    constants.EmptyString, // it will change per line
	}

	for args.Index, args.Line = range *lines {
		newStrings[args.Index] = loopProcessorRune(&args)
	}

	return &newStrings
}

// Loops through all the lines.
func (stringWrapperPtr *StringWrapperPtr) LoopUnixLinesToStringArray(
	loopProcessorRune func(args *strhelpercore.StringWrapperLineLoopArgs) *string,
) *[]*string {
	lines := stringWrapperPtr.GetLinesUnix()
	newStrings := make([]*string, len(*lines))

	args := strhelpercore.StringWrapperLineLoopArgs{
		Content: stringWrapperPtr.content,
		Lines:   lines,
		Index:   -1,                    // it will change per line
		Line:    constants.EmptyString, // it will change per line
	}

	for args.Index, args.Line = range *lines {
		newStrings[args.Index] = loopProcessorRune(&args)
	}

	return &newStrings
}

// Create regular expression from current string.
// Recommendation, do not create regular expressions inside a function call, keep it on top of the file as variables
func (stringWrapperPtr *StringWrapperPtr) CreateRegularExpression() *regexp.Regexp {
	return regexp.MustCompile(*stringWrapperPtr.content)
}

// Same as Value()
func (stringWrapperPtr *StringWrapperPtr) String() string {
	return stringWrapperPtr.ValueWithoutPtr()
}

// returns -1 if the index is not present in strings lengthInBytes.
// or else returns the character value from that index
// performance should be very slow, use direct access of str.
func (stringWrapperPtr *StringWrapperPtr) GetSafeIndexAt(index int) int16 {
	if !stringWrapperPtr.HasIndex(index) {
		return constants.InvalidNotFoundCase
	}

	return int16(stringWrapperPtr.ValueWithoutPtr()[index])
}

// returns -1 if the index is not present in strings lengthInBytes.
// or else returns the character value from that index
// performance should be very slow, use direct access of stringWrapperPtr.ToRunesPtr().
func (stringWrapperPtr *StringWrapperPtr) GetSafeRuneIndexAt(index int) rune {
	if !stringWrapperPtr.HasIndex(index) {
		return constants.InvalidNotFoundCase
	}

	return stringWrapperPtr.ToRunes()[index]
}

func (stringWrapperPtr *StringWrapperPtr) IsEqualAtIndex(
	index int,
	char uint8,
	isCaseSensitive bool,
) bool {
	valueAt := stringWrapperPtr.ValueWithoutPtr()[index]

	if isCaseSensitive {
		return valueAt == char
	}

	return charhelper.IsMatchCaseInsensitive(valueAt, char)
}

// (*stringWrapperPtr).LengthInBytes()-1 >= index
func (stringWrapperPtr *StringWrapperPtr) HasIndex(index int) bool {
	return (*stringWrapperPtr).LengthInBytes()-1 >= index
}

// (*stringWrapperPtr).Length()-1 >= index
func (stringWrapperPtr *StringWrapperPtr) HasRuneIndex(index int) bool {
	return (*stringWrapperPtr).Length()-1 >= index
}

// Returns a new string builder contains text of stringWrapperPtr and has a
// growth = stringWrapperPtr.lengthInBytes + additionalGrowLength
func (stringWrapperPtr *StringWrapperPtr) Builder(additionalGrowLength int) strings.Builder {
	builder := strings.Builder{}
	currentString := stringWrapperPtr.content
	length := stringWrapperPtr.LengthInBytes() + additionalGrowLength
	builder.Grow(length)
	builder.WriteString(*currentString)
	return builder
}

// Returns a new string builder contains text of stringWrapperPtr + str and has a
// growth = stringWrapperPtr.lengthInBytes + additionalGrowLength + len(str)
func (stringWrapperPtr *StringWrapperPtr) BuilderWithStr(str *string, additionalGrowLength int) strings.Builder {
	builder := strings.Builder{}
	currentString := stringWrapperPtr.content
	length := stringWrapperPtr.LengthInBytes() + len(*str) + additionalGrowLength
	builder.Grow(length)
	builder.WriteString(*currentString)
	builder.WriteString(*str)

	return builder
}

// Better to use slice or builder for appending lines in a loop.
//
// stringWrapperPtr.content + separator + JoinAll(separator, stringWrapperPtrs)
func (stringWrapperPtr *StringWrapperPtr) AppendLines(isSkipOnEmpty bool, contents ...string) *StringWrapperPtr {
	return stringWrapperPtr.concat(constants.NewLine, isSkipOnEmpty, &contents)
}

// Better to use slice or builder for appending or concatenating lines in a loop.
//
// stringWrapperPtr.content + separator + JoinAll(separator, stringWrapperPtrs)
func (stringWrapperPtr *StringWrapperPtr) Concat(contents ...string) *StringWrapperPtr {
	return stringWrapperPtr.concat(
		constants.EmptyString,
		false, // must add everything
		&contents)
}

// Better to use slice or builder for appending lines.
//
// stringWrapperPtr.content + separator + JoinAll(separator, stringWrapperPtrs)
func (stringWrapperPtr *StringWrapperPtr) ConcatWrappers(
	separator string,
	isSkipOnEmpty bool,
	stringWrapperPtrs ...StringWrapperPtr,
) *StringWrapperPtr {
	strArray := make([]*string, len(stringWrapperPtrs))

	for i, wrapper := range stringWrapperPtrs {
		strArray[i] = wrapper.content
	}

	return stringWrapperPtr.ConcatStrPtr(
		separator,
		isSkipOnEmpty,
		&strArray)
}

// Better to use slice or builder for appending lines.
//
// stringWrapperPtr.content + separator + JoinAll(separator, stringWrapperPtrs)
func (stringWrapperPtr *StringWrapperPtr) ConcatWrappersPtrs(
	separator string,
	isSkipOnEmpty bool,
	stringWrapperPtrs ...*StringWrapperPtr,
) *StringWrapperPtr {
	strArray := make([]*string, len(stringWrapperPtrs))

	for i, wrapper := range stringWrapperPtrs {
		strArray[i] = wrapper.content
	}

	return stringWrapperPtr.ConcatStrPtr(
		separator,
		isSkipOnEmpty,
		&strArray)
}

// Better to use slice or builder for appending lines.
//
// stringWrapperPtr.content + separator + JoinAll(separator, stringWrapperPtrs)
func (stringWrapperPtr *StringWrapperPtr) ConcatWithSeparator(
	separator string,
	isSkipOnEmpty bool,
	contents ...string,
) *StringWrapperPtr {
	return stringWrapperPtr.concat(
		separator,
		isSkipOnEmpty,
		&contents)
}

// Better to use slice or builder for appending lines.
//
// stringWrapperPtr.content + separator + JoinAll(separator, stringWrapperPtrs)
func (stringWrapperPtr *StringWrapperPtr) ConcatPtrsWithSeparator(
	separator string,
	isSkipOnEmpty bool,
	contents ...*string,
) *StringWrapperPtr {
	return stringWrapperPtr.ConcatStrPtr(
		separator,
		isSkipOnEmpty,
		&contents)
}

// combine current wrapper strings + all given ones with given separator and returns as wrapper
//
// stringWrapperPtr.content + separator + JoinAll(separator, stringWrapperPtrs)
func (stringWrapperPtr *StringWrapperPtr) concat(
	separator string,
	isSkipOnEmpty bool,
	contents *[]string,
) *StringWrapperPtr {
	combinedResult := concat.StringsArrayWithSeparator(
		stringWrapperPtr.content,
		&separator,
		isSkipOnEmpty,
		contents)

	return NewSwPtr(&combinedResult)
}

// combine current wrapper strings + all given ones with given separator and returns as wrapper
func (stringWrapperPtr *StringWrapperPtr) ConcatStrPtr(
	separator string,
	isSkipOnEmpty bool,
	contents *[]*string,
) *StringWrapperPtr {
	combinedResult := concat.PtrStringsArrayWithSeparator(
		stringWrapperPtr.content,
		&separator,
		isSkipOnEmpty,
		contents)

	return NewSwPtr(&combinedResult)
}

func (stringWrapperPtr *StringWrapperPtr) ReplaceWrapper(
	searchingWrapper,
	replacingWrapper *StringWrapperPtr,
	isCaseSensitive bool,
	startsAt int,
	replaceCount int,
) *StringWrapperPtr {
	replacedText := strhelper.ReplacePtr(
		stringWrapperPtr.content,
		searchingWrapper.content,
		replacingWrapper.content,
		startsAt,
		replaceCount,
		isCaseSensitive,
	)

	return NewSwPtr(&replacedText)
}

// For better performance use Ptr version.
func (stringWrapperPtr *StringWrapperPtr) Replace(
	search,
	replaceText string,
	isCaseSensitive bool,
	startsAt int,
	replaceCount int,
) string {
	return strhelper.ReplacePtr(
		stringWrapperPtr.content,
		&search,
		&replaceText,
		startsAt,
		replaceCount,
		isCaseSensitive,
	)
}

func (stringWrapperPtr *StringWrapperPtr) ReplacePtr(
	search,
	replaceText *string,
	startsAt int,
	replaceCount int,
	isCaseSensitive bool,
) *string {
	replacedText := strhelper.ReplacePtr(
		stringWrapperPtr.content,
		search,
		replaceText,
		startsAt,
		replaceCount,
		isCaseSensitive,
	)

	return &replacedText
}

func (stringWrapperPtr *StringWrapperPtr) ReplaceAll(
	search,
	replaceText string,
	isCaseSensitive bool,
	startsAt int,
) string {
	return strhelper.ReplacePtr(
		stringWrapperPtr.content,
		&search,
		&replaceText,
		startsAt,
		-1,
		isCaseSensitive,
	)
}

func (stringWrapperPtr *StringWrapperPtr) ReplaceMultiple(
	searchReplaceMap *map[string]string,
	startsAt int,
) string {
	return strhelper.ReplaceMultiplePtr(
		stringWrapperPtr.content,
		searchReplaceMap,
		startsAt,
		-1,
		true,
	)
}

func (stringWrapperPtr *StringWrapperPtr) ReplaceMultipleCase(
	searchReplaceMap *map[string]string,
	startsAt int,
	limits int,
	isCaseSensitive bool,
) string {
	return strhelper.ReplaceMultiplePtr(
		stringWrapperPtr.content,
		searchReplaceMap,
		startsAt,
		limits,
		isCaseSensitive,
	)
}

func (stringWrapperPtr *StringWrapperPtr) LastIndexOf(
	search string,
	lastStartIndexReducedBy int,
	isCaseSensitive bool,
) int {
	return strhelper.LastIndexOfPtr(
		stringWrapperPtr.content,
		&search,
		lastStartIndexReducedBy,
		isCaseSensitive)
}

func (stringWrapperPtr *StringWrapperPtr) LastIndexOfPtr(
	search *string,
	lastStartIndexReducedBy int,
	isCaseSensitive bool,
) int {
	return strhelper.LastIndexOfPtr(
		stringWrapperPtr.content,
		search,
		lastStartIndexReducedBy,
		isCaseSensitive)
}

// For better performance use strhelper.IsStartsWithPtr
func (stringWrapperPtr *StringWrapperPtr) IsStartsWith(
	search string,
	isCaseSensitive bool,
	startsAt int,
) bool {
	return strhelper.IsStartsWithPtr(
		stringWrapperPtr.content,
		&search,
		startsAt,
		isCaseSensitive)
}

// Use direct strhelper.IsEndsWithPtr will be faster
func (stringWrapperPtr *StringWrapperPtr) IsEndsWith(
	endsWith string,
	isCaseSensitive bool,
	startsAt int,
) bool {
	return strhelper.IsEndsWithPtr(
		stringWrapperPtr.content,
		&endsWith,
		startsAt,
		isCaseSensitive)
}

// Use direct strhelper.IsEndsWithPtr will be faster
func (stringWrapperPtr *StringWrapperPtr) IsEndsWithPtr(
	endsWith *string,
	isCaseSensitive bool,
	startsAt int,
) bool {
	return strhelper.IsEndsWithPtr(
		stringWrapperPtr.content,
		endsWith,
		startsAt,
		isCaseSensitive)
}

func (stringWrapperPtr *StringWrapperPtr) PadLeftWithSpace(width int) string {
	return strhelper.PadSpaceLeft(stringWrapperPtr.content, width)
}

func (stringWrapperPtr *StringWrapperPtr) PadRightWithSpace(width int) string {
	return strhelper.PadSpaceRight(stringWrapperPtr.content, width)
}

func (stringWrapperPtr *StringWrapperPtr) PadLeft(width int, padding string) string {
	return strhelper.PadLeft(stringWrapperPtr.content, &padding, width)
}

func (stringWrapperPtr *StringWrapperPtr) PadRight(width int, padding string) string {
	return strhelper.PadRight(stringWrapperPtr.content, &padding, width)
}

func (stringWrapperPtr *StringWrapperPtr) Pad(width int, padding string, isLeft, isRight bool) string {
	return strhelper.Pad(stringWrapperPtr.content, &padding, width, isLeft, isRight)
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
func (stringWrapperPtr *StringWrapperPtr) MultiSplit(
	startsAt, limits int,
	splitters ...string,
) *strhelpercore.SplitResultOverview {
	return strhelper.MultipleSplitsPtr(
		stringWrapperPtr.content,
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
func (stringWrapperPtr *StringWrapperPtr) MultiSplitCase(
	startsAt, limits int,
	isCaseSensitive bool,
	splitters ...string,
) *strhelpercore.SplitResultOverview {
	return strhelper.MultipleSplitsPtr(
		stringWrapperPtr.content,
		&splitters,
		startsAt,
		limits,
		isCaseSensitive)
}

func (stringWrapperPtr *StringWrapperPtr) Split(splitter string) []string {
	return strings.Split(
		*stringWrapperPtr.content,
		splitter)
}

func (stringWrapperPtr *StringWrapperPtr) SplitN(splitter string, count int) []string {
	return strings.SplitN(
		*stringWrapperPtr.content,
		splitter,
		count)
}
