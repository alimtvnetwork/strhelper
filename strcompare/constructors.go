package strcompare

import "gitlab.com/evatix-go/strhelper/strswasync"

func New(left, right string) *Result {
	result := Result{
		leftText:   &left,
		rightText:  &right,
		createType: PlainText,
	}

	return &result
}

func NewPtr(left, right *string) *Result {
	result := Result{
		leftText:   left,
		rightText:  right,
		createType: PlainText,
	}

	return &result
}

func NewUsingLines(left, right *[]string) *Result {
	result := Result{
		leftLines:  strswasync.NewPtr(left),
		rightLines: strswasync.NewPtr(right),
		createType: Lines,
	}

	return &result
}
