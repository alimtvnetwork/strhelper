// Refers to StringsWrapperAsync alias as `strswasync`
//
// Thread safety ensured.
//
// This for async transaction purpose.
package strswasync

import (
	"strings"
	"sync"
	"unsafe"

	"gitlab.com/evatix-go/strhelper/concat"
	"gitlab.com/evatix-go/strhelper/ds/strhashset"
	"gitlab.com/evatix-go/strhelper/isstr"
	"gitlab.com/evatix-go/strhelper/lines"
	"gitlab.com/evatix-go/strhelper/remove"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
	"gitlab.com/evatix-go/strhelper/strs"
	"gitlab.com/evatix-go/strhelper/strs/isstrs"
	"gitlab.com/evatix-go/strhelper/strs/strsindex"
	"gitlab.com/evatix-go/strhelper/swasync"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

// Refers to StringWrapperPointer Async alias as `swasync`
//
// Thread safety ensured.
//
// This for async transaction purpose.
type Wrapper struct {
	lines                        *[]string
	linesWrappers                *[]*swasync.StringWrapper
	trimmedLines                 *[]string
	trimmedLinesWrappers         *[]*swasync.StringWrapper
	lowerLines                   *[]string
	upperLines                   *[]string
	lowerWrappers                *[]*swasync.StringWrapper
	upperWrappers                *[]*swasync.StringWrapper
	regExWrappersCollection      *strhelpercore.RegExWrappersCollection
	bytes                        *[]byte
	jsonBytes                    *[]byte
	isNullOrEmptyFirstItem       *bool
	isEmptyOrWhitespaceFirstItem *bool
	bytesLength                  *int
	sync.Mutex
	length int
}

func NewPtr(lines *[]string) *Wrapper {
	return &Wrapper{
		lines:         lines,
		linesWrappers: nil,
		lowerLines:    nil,
		upperLines:    nil,
		Mutex:         sync.Mutex{},
		length:        len(*lines),
	}
}

func New(lines []string) *Wrapper {
	if lines == nil {
		return &Wrapper{
			lines:         nil,
			linesWrappers: nil,
			lowerLines:    nil,
			upperLines:    nil,
			Mutex:         sync.Mutex{},
			length:        0,
		}
	}

	return &Wrapper{
		lines:         &lines,
		linesWrappers: nil,
		lowerLines:    nil,
		upperLines:    nil,
		Mutex:         sync.Mutex{},
		length:        len(lines),
	}
}

func (wrapper *Wrapper) Lock() {
	wrapper.Mutex.Lock()
}

func (wrapper *Wrapper) Unlock() {
	wrapper.Mutex.Unlock()
}

func (wrapper *Wrapper) Value() *[]string {
	return wrapper.lines
}

func (wrapper *Wrapper) ValueWithoutPtr() []string {
	return *wrapper.lines
}

func (wrapper *Wrapper) Length() int {
	return wrapper.length
}

func (wrapper *Wrapper) BytesLength() int {
	if wrapper.bytesLength == nil {
		bytesLength := 0
		for _, line := range *wrapper.lines {
			bytesLength += len(line)
		}

		wrapper.bytesLength = &bytesLength
	}

	return *wrapper.bytesLength
}

func (wrapper *Wrapper) IsEquals(items *[]string, isCaseSensitive bool) bool {
	if len(*items) != wrapper.Length() {
		return false
	}

	if isCaseSensitive {
		for i := 0; i < wrapper.Length(); i++ {
			if (*wrapper.lines)[i] != (*items)[i] {
				return false
			}
		}

		return true
	}

	// insensitive
	lowerLines := *wrapper.ToLowersPtr()
	searchLowers := strs.ToLowerStrings(items)
	for i := 0; i < wrapper.Length(); i++ {
		if (lowerLines)[i] != (*searchLowers)[i] {
			return false
		}
	}

	return true
}

func (wrapper *Wrapper) IsAnyEquals(
	isCaseSensitive bool,
	searchItems ...*[]string,
) bool {
	for _, searchItem := range searchItems {
		if wrapper.IsEquals(searchItem, isCaseSensitive) {
			return true
		}
	}

	return false
}

func (wrapper *Wrapper) IsAllEquals(
	isCaseSensitive bool,
	searchItems ...*[]string,
) bool {
	for _, searchItem := range searchItems {
		if wrapper.IsEquals(searchItem, isCaseSensitive) == false {
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
func (wrapper *Wrapper) IndexesOfAll(
	findingString *string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) *[]int {
	return strsindex.OfAll(
		wrapper.lines,
		findingString,
		startsAtIndex,
		limits,
		isCaseSensitive)
}

// Find all the indexes for all the searchTerms given.
//
// limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
func (wrapper *Wrapper) ManyIndexesOfAll(
	searchTerms *[]string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) *strhelpercore.IndexesResultSet {
	return strsindex.OfAllMany(
		wrapper.lines,
		searchTerms,
		startsAtIndex,
		limits,
		isCaseSensitive)
}

func (wrapper *Wrapper) IsNullOrEmptyFirstItem() bool {
	if wrapper.isNullOrEmptyFirstItem == nil {
		isNullOrEmptyFirstItem := wrapper.IsNull() || (wrapper.length > 0 && (*wrapper.lines)[0] == "")
		wrapper.isNullOrEmptyFirstItem = &isNullOrEmptyFirstItem
	}

	return *wrapper.isNullOrEmptyFirstItem
}

func (wrapper *Wrapper) IsNullOrEmptyFirstItemWhitespace() bool {
	if wrapper.isEmptyOrWhitespaceFirstItem == nil {
		isEmptyOrWhitespaceFirstItem := wrapper.IsNullOrEmptyFirstItem() ||
			(wrapper.length > 0 && whitespace.IsWhitespaces(&(*wrapper.lines)[0]))
		wrapper.isEmptyOrWhitespaceFirstItem = &isEmptyOrWhitespaceFirstItem
	}

	return *wrapper.isEmptyOrWhitespaceFirstItem
}

// IsNull(s) || IsNullOrEmpty(s)
func (wrapper *Wrapper) IsNull() bool {
	return wrapper.lines == nil || *wrapper.lines == nil
}

func (wrapper *Wrapper) Lines() *[]string {
	return wrapper.lines
}

func (wrapper *Wrapper) GetAsWrappers() *[]*swasync.StringWrapper {
	wrapper.Lock()
	defer wrapper.Unlock()

	if wrapper.linesWrappers == nil {
		wrappers := make([]*swasync.StringWrapper, wrapper.length)

		for i := 0; i < wrapper.length; i++ {
			wrappers[i] = swasync.New(&(*wrapper.lines)[i])
		}

		*wrapper.linesWrappers = wrappers
	}

	return wrapper.linesWrappers
}

// get bytes array ptr
func (wrapper *Wrapper) ToBytesPtr() *[]byte {
	wrapper.Lock()
	defer wrapper.Unlock()

	if wrapper.bytes == nil || *wrapper.bytes == nil {
		// Reference : https://bit.ly/2ITGTaU
		*wrapper.bytes = *(*[]byte)(unsafe.Pointer(wrapper.lines))
	}

	return wrapper.bytes
}

// Returns:
//  - true : if @lines are nil.
//  - true : if all lines are blank (whitespace or empty or nil)
func (wrapper *Wrapper) IsAllBlank() bool {
	return isstrs.AllBlank(wrapper.lines)
}

// Returns:
//  - false : if @lines are nil.
//  - true : if all defined (NOT whitespace or empty or nil)
func (wrapper *Wrapper) IsAllDefined() bool {
	return isstrs.AllDefined(wrapper.lines)
}

// Returns:
//  - true : if @lines are nil.
//  - true : if any line in lines is blank (whitespace or empty or nil)
func (wrapper *Wrapper) IsAnyBlank() bool {
	return isstrs.AnyBlank(wrapper.lines)
}

// Returns:
//  - false : if @lines are nil.
//  - true : if all @lines are defined (not whitespace or empty or nil)
func (wrapper *Wrapper) IsAnyDefined() bool {
	return isstrs.AnyDefined(wrapper.lines)
}

func (wrapper *Wrapper) ToLowersPtr() *[]string {
	wrapper.Lock()
	defer wrapper.Unlock()

	if wrapper.lowerLines == nil && !wrapper.IsNull() {
		lowerLines := make([]string, wrapper.length)
		allLines := wrapper.lines
		for i := 0; i < wrapper.length; i++ {
			lowerLines[i] = strings.ToLower((*allLines)[i])
		}

		wrapper.lowerLines = &lowerLines
	}

	return wrapper.lowerLines
}

func (wrapper *Wrapper) ToUppersPtr() *[]string {
	wrapper.Lock()
	defer wrapper.Unlock()

	if wrapper.lowerLines == nil && !wrapper.IsNull() {
		upperLines := make([]string, wrapper.length)
		allLines := wrapper.lines
		for i := 0; i < wrapper.length; i++ {
			upperLines[i] = strings.ToUpper((*allLines)[i])
		}

		wrapper.upperLines = &upperLines
	}

	return wrapper.upperLines
}

func (wrapper *Wrapper) ToLowersWrappersPtr() *[]*swasync.StringWrapper {
	wrapper.Lock()
	defer wrapper.Unlock()

	if wrapper.lowerWrappers == nil {
		allLowers := wrapper.ToLowersPtr()
		wrappers := make([]*swasync.StringWrapper, wrapper.length)

		for i, lowerLine := range *allLowers {
			wrappers[i] = swasync.New(&lowerLine)
		}

		wrapper.lowerWrappers = &wrappers
	}

	return wrapper.lowerWrappers
}

func (wrapper *Wrapper) ToUpperWrappersPtr() *[]*swasync.StringWrapper {
	wrapper.Lock()
	defer wrapper.Unlock()

	if wrapper.upperWrappers == nil {
		allUppers := wrapper.ToUppersPtr()
		wrappers := make([]*swasync.StringWrapper, wrapper.length)

		for i, upperLine := range *allUppers {
			wrappers[i] = swasync.New(&upperLine)
		}

		wrapper.upperWrappers = &wrappers
	}

	return wrapper.upperWrappers
}

func (wrapper *Wrapper) At(index int) *string {
	return &(*wrapper.lines)[index]
}

func (wrapper *Wrapper) WrapperAt(index int) *swasync.StringWrapper {
	return (*wrapper.GetAsWrappers())[index]
}

// Loops through all the lines.
func (wrapper *Wrapper) LoopLines(
	lineProcessor strhelpercore.LineProcessor,
) *[]*string {
	allLines := wrapper.lines

	return lines.Process(
		nil,
		allLines,
		&lineProcessor)
}

// Loops through all the lines.
func (wrapper *Wrapper) LoopParallel(
	lineProcessor strhelpercore.LineProcessor,
) *[]*string {
	allLines := wrapper.lines

	return lines.ProcessAsync(
		nil,
		allLines,
		&lineProcessor)
}

// Returns
//  - defaultValue : if the index is not valid.
//  - *string : if index is valid and within the range
func (wrapper *Wrapper) GetSafeIndexAt(index int, defaultValue *string) *string {
	if !wrapper.HasIndex(index) || index < 0 {
		return defaultValue
	}

	return &(*wrapper.lines)[index]
}

func (wrapper *Wrapper) IsEqualAtIndex(
	index int,
	compareStr *string,
	isCaseSensitive bool,
) bool {
	valueAt := (*wrapper.lines)[index]

	return isstr.EqualsPtr(&valueAt, compareStr, isCaseSensitive)
}

// wrapper.BytesLength()-1 >= index
func (wrapper *Wrapper) HasIndex(index int) bool {
	return wrapper.length-1 >= index
}

func (wrapper *Wrapper) HasLengthOf(length int) bool {
	return wrapper.Length() >= length
}

// Returns a new string builder contains text of wrapper.lines with separator and has a
// growth = wrapper.bytesLength + additionalGrowLength
//
// Warning :
//  - It doesn't care about empty line or anything.
//  - panic if str is nil.
func (wrapper *Wrapper) Builder(separator *string, additionalGrowLength int) strings.Builder {
	return concat.GetBuilder(wrapper.lines, separator, additionalGrowLength)
}

// Returns a new string builder contains text of wrapper.lines with separator and has a
// growth = wrapper.bytesLength + additionalGrowLength + len(str)
//
// Warning :
//  - It doesn't care about empty line or anything.
//  - panic if str is nil.
func (wrapper *Wrapper) BuilderWithStr(separator *string, str *string, additionalGrowLength int) strings.Builder {
	length := len(*str)
	builder := concat.GetBuilder(wrapper.lines, separator, additionalGrowLength+length)
	builder.WriteString(*separator)
	builder.WriteString(*str)

	return builder
}

// Add the @contents before the content of Wrapper.Lines()
func (wrapper *Wrapper) Prepend(contents ...string) *[]string {
	return wrapper.Prepends(
		true,
		nil,
		&contents)
}

// Prepends array @contents before @StringWrapper.Lines()
//
// Look for more details in concat.ArraysOfArraysToArray
func (wrapper *Wrapper) Prepends(
	isSkipEmptyOrNil bool,
	skipFilter *strhashset.Hashset,
	contents *[]string,
) *[]string {
	combinedResult := concat.ArraysOfArraysToArray(
		isSkipEmptyOrNil,
		skipFilter,
		contents, // pre
		wrapper.lines) // post

	return combinedResult
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
func (wrapper *Wrapper) PrependAsString(
	separator string,
	isSkipOnEmpty bool,
	contents *[]string,
) *string {
	combinedResult := concat.CombineArrayWithAnotherUsingSeparator(
		&separator,
		isSkipOnEmpty,
		contents,
		wrapper.lines)

	return &combinedResult
}

// Wrapper.Lines() + @contents all joined to single string using separator.
func (wrapper *Wrapper) ConcatAsString(
	separator string,
	isSkipOnEmpty bool,
	contents *[]string,
) *string {
	combinedResult := concat.CombineArrayWithAnotherUsingSeparator(
		&separator,
		isSkipOnEmpty,
		wrapper.lines, // pre
		contents) // post

	return &combinedResult
}

// Wrapper.Lines() + @contents
func (wrapper *Wrapper) Concat(contents ...string) *[]string {
	return wrapper.Concatenates(
		true,
		nil,
		&contents)
}

// Wrapper.Lines() + @contents
//
// @skipFilter:
//  - Items will be skipped from the compiled one.
func (wrapper *Wrapper) Concatenates(
	isSkipEmptyOrNil bool,
	skipFilter *strhashset.Hashset,
	contents *[]string,
) *[]string {
	combinedResult := concat.ArraysOfArraysToArray(
		isSkipEmptyOrNil,
		skipFilter,
		wrapper.lines, // pre
		contents,      // post
	)

	return combinedResult
}

// Creates new lines where removeStr will not appear.
//
// @count:
// - `-1` meaning remove all.
// - or else remove up to the given number only.
func (wrapper *Wrapper) Remove(
	removeString *string,
	isCaseSensitive bool,
	startsAt int,
	count int,
) *[]string {
	return remove.Lines(
		wrapper.lines,
		removeString,
		startsAt,
		count,
		isCaseSensitive,
	)
}

// Creates new lines where removeStr will not appear.
//
// @count (default given):
// - `-1` meaning remove all.
func (wrapper *Wrapper) RemoveAll(
	removeString *string,
	isCaseSensitive bool,
	startsAt int,
) *[]string {
	return remove.Lines(
		wrapper.lines,
		removeString,
		startsAt,
		-1,
		isCaseSensitive,
	)
}

func (wrapper *Wrapper) LastIndexOf(
	search string,
	lastStartIndexReducedBy int,
	isCaseSensitive bool,
) int {
	return strsindex.OfLastPtr(
		wrapper.lines,
		&search,
		lastStartIndexReducedBy,
		isCaseSensitive)
}

// ReversePtr doesn't reverse in place but create a new one.
func (wrapper *Wrapper) ReversePtr() *[]string {
	return strs.ReverseStringsPtr(wrapper.lines)
}

func (wrapper *Wrapper) LastIndexOfPtr(
	search *string,
	lastStartIndexReducedBy int,
	isCaseSensitive bool,
) int {
	return strsindex.OfLastPtr(
		wrapper.lines,
		search,
		lastStartIndexReducedBy,
		isCaseSensitive)
}

// IsContains alias of isstrs.Contains and
// returns true if the search text contains any where in the text after the start index.
//
// Use direct isstr.ContainsPtr will be faster
func (wrapper *Wrapper) IsContains(
	search *string,
	isCaseSensitive bool,
	startsAt int,
) bool {
	return isstrs.Contains(
		wrapper.lines,
		search,
		startsAt,
		isCaseSensitive)
}

// Has is alias of IsContains and returns true if the search text contains any where in the text after the start index.
//
// Use direct isstr.ContainsPtr will be faster
func (wrapper *Wrapper) Has(search *string) bool {
	return isstrs.Contains(
		wrapper.lines,
		search,
		0,
		true)
}

// HasAll usages isstrs.IsContains(...) with loop and returns true if the all search text exists.
func (wrapper *Wrapper) HasAll(searchTerms ...string) bool {
	for _, searchTerm := range searchTerms {
		if isstrs.Contains(
			wrapper.lines,
			&searchTerm,
			0,
			true) == false {
			// not found
			return false
		}
	}

	// all found
	return true
}

// HasAny usages isstrs.IsContains(...) with loop and returns true if the any search text exists.
func (wrapper *Wrapper) HasAny(searchTerms ...string) bool {
	for _, searchTerm := range searchTerms {
		if isstrs.Contains(
			wrapper.lines,
			&searchTerm,
			0,
			true) {
			// any found
			return true
		}
	}

	// not found any
	return false
}

func (wrapper *Wrapper) ToHashset() *strhashset.Hashset {
	return strhashset.NewUsingArray(wrapper.lines)
}

func (wrapper *Wrapper) ToLowerHashset() *strhashset.Hashset {
	return strhashset.NewUsingArray(wrapper.ToLowersPtr())
}

func (wrapper *Wrapper) initializeRegExWrappersCollection() {
	if wrapper.regExWrappersCollection == nil {
		wrapper.regExWrappersCollection = strhelpercore.NewRegExWrappersCollection(wrapper.lines)
	}
}

func (wrapper *Wrapper) RegExWrappersCollection() *strhelpercore.RegExWrappersCollection {
	wrapper.initializeRegExWrappersCollection()

	return wrapper.regExWrappersCollection
}

func (wrapper *Wrapper) RegexWrappers() *[]*strhelpercore.RegExWrapper {
	return wrapper.RegExWrappersCollection().Regexes()
}

func (wrapper *Wrapper) RegexWrappersMap() *map[string]*strhelpercore.RegExWrapper {
	return wrapper.RegExWrappersCollection().RegexesMap()
}
