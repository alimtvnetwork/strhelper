package strhelpercore

import "gitlab.com/evatix-go/strhelper/internal/misc"

type SplitResultOverview struct {
	Results                 *[]*string
	NonEmptyResults         *[]*string
	SplitResults            *[]*SplitResult
	NonEmptySplitResults    *[]*SplitResult
	resultsLength           *int
	nonEmptyResultsLength   *int
	regExWrappersCollection *RegExWrappersCollection
	IsEmptyResult           bool
}

// ResultsToRegexMap returns Results to regular expressions map
func (splitResultOverview *SplitResultOverview) ResultsToRegexMap() *map[string]*RegExWrapper {
	splitResultOverview.initializeRegExWrappersCollection()

	return splitResultOverview.regExWrappersCollection.RegexesMap()
}

func (splitResultOverview *SplitResultOverview) initializeRegExWrappersCollection() {
	if splitResultOverview.regExWrappersCollection == nil {
		splitResultOverview.regExWrappersCollection =
			NewRegExWrappersCollectionUsingStringPointer(splitResultOverview.Results)
	}
}

// ResultsToRegexArray returns Results to regular expressions array
func (splitResultOverview *SplitResultOverview) ResultsToRegexArray() *[]*RegExWrapper {
	splitResultOverview.initializeRegExWrappersCollection()

	return splitResultOverview.regExWrappersCollection.Value()
}

// Returns the cached length of Results
func (splitResultOverview *SplitResultOverview) ResultsLength() int {
	if splitResultOverview.resultsLength == nil {
		length := len(*splitResultOverview.Results)
		splitResultOverview.resultsLength = &length
	}

	return *splitResultOverview.resultsLength
}

// Returns the cached length of NonEmptyResults
func (splitResultOverview *SplitResultOverview) NonEmptyResultsLength() int {
	if splitResultOverview.nonEmptyResultsLength == nil {
		length := len(*splitResultOverview.NonEmptyResults)
		splitResultOverview.nonEmptyResultsLength = &length
	}

	return *splitResultOverview.nonEmptyResultsLength
}

// Returns Results from *[]*string to *[]string
func (splitResultOverview *SplitResultOverview) ToSimpleArray() *[]string {
	return misc.ConvertPointerStringsToStrings(splitResultOverview.Results)
}

// Returns NonEmptyResults *[]*string to *[]string
func (splitResultOverview *SplitResultOverview) NonEmptyToSimpleArray() *[]string {
	return misc.ConvertPointerStringsToStrings(splitResultOverview.NonEmptyResults)
}

func NewEmptySplitResultOverview(str *string) *SplitResultOverview {
	strArray := []*string{str}

	return &SplitResultOverview{
		Results:       &strArray,
		SplitResults:  nil,
		IsEmptyResult: true,
	}
}
