package strhelpercore

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
)

type StringResult struct {
	FoundIndex int
	Line       string
	IsFound    bool
}

func InvalidStringResult() *StringResult {
	return &StringResult{
		FoundIndex: constants.InvalidNotFoundCase,
		Line:       constants.EmptyString,
		IsFound:    false,
	}
}

func (it *StringResult) String() string {
	return converters.AnyToValueString(*it)
}
