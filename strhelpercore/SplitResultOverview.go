package strhelpercore

type SplitResultOverview struct {
	Results                 []string
	NonEmptyResults         []string
	SplitResults            []*SplitResult
	NonEmptySplitResults    []*SplitResult
	resultsLength           *int // todo optimize to non ptr
	nonEmptyResultsLength   *int
	regExWrappersCollection *RegExWrappersCollection
	IsEmptyResult           bool
}

// ResultsToRegexMap returns Results to regular expressions map
func (it *SplitResultOverview) ResultsToRegexMap() map[string]*RegExWrapper {
	it.initializeRegExWrappersCollection()

	return it.regExWrappersCollection.RegexesMap()
}

func (it *SplitResultOverview) initializeRegExWrappersCollection() {
	if it.regExWrappersCollection == nil {
		it.regExWrappersCollection =
			NewRegExWrappersCollection(it.Results...)
	}
}

// ResultsToRegexArray returns Results to regular expressions array
func (it *SplitResultOverview) ResultsToRegexArray() []*RegExWrapper {
	it.initializeRegExWrappersCollection()

	return it.regExWrappersCollection.Value()
}

// ResultsLength Returns the cached length of Results
func (it *SplitResultOverview) ResultsLength() int {
	if it.resultsLength != nil {
		return *it.resultsLength
	}

	if it.resultsLength == nil {
		length := len(it.Results)
		it.resultsLength = &length
	}

	return *it.resultsLength
}

// NonEmptyResultsLength  Returns the cached length of NonEmptyResults
func (it *SplitResultOverview) NonEmptyResultsLength() int {
	if it.nonEmptyResultsLength == nil {
		length := len(it.NonEmptyResults)
		it.nonEmptyResultsLength = &length
	}

	return *it.nonEmptyResultsLength
}

// ToSimpleArray Returns Results from *[]*string to *[]string
func (it *SplitResultOverview) ToSimpleArray() []string {
	return it.Results
}

// NonEmptyToSimpleArray Returns NonEmptyResults *[]*string to *[]string
func (it *SplitResultOverview) NonEmptyToSimpleArray() []string {
	return it.NonEmptyResults
}

func NewEmptySplitResultOverview(str string) *SplitResultOverview {
	strArray := []string{str}

	return &SplitResultOverview{
		Results:       strArray,
		SplitResults:  nil,
		IsEmptyResult: true,
	}
}
