package strhelpercore

type SplitResultOverview struct {
	Results              *[]*string
	NonEmptyResults      *[]*string
	SplitResults         *[]*SplitResult
	NonEmptySplitResults *[]*SplitResult
	IsEmptyResult        bool
}

func (splitResultOverview *SplitResultOverview) ToSimpleArray(ptrStrArray *[]*string) *[]string {
	return convertPtrStringArrayToStringArray(ptrStrArray)
}

func NewEmptySplitResultOverview(str *string) *SplitResultOverview {
	strArray := []*string{str}

	return &SplitResultOverview{
		Results:       &strArray,
		SplitResults:  nil,
		IsEmptyResult: true,
	}
}
