package indextestwrappers

import (
	"gitlab.com/evatix-go/core/coretests"
)

type GenericIndexOfWrapper struct {
	Content             string
	SearchingContent    string
	InitializedPosition int
	Limits              int
	IsCaseSensitive     bool
	HasPanic            bool
	funcName            coretests.TestFuncName
	expected            interface{}
	actual              interface{}
}

func (compareTestWrapper *GenericIndexOfWrapper) Actual() interface{} {
	return compareTestWrapper.actual
}

func (compareTestWrapper *GenericIndexOfWrapper) SetActual(actual interface{}) {
	compareTestWrapper.actual = actual
}

func (compareTestWrapper *GenericIndexOfWrapper) FuncName() string {
	return compareTestWrapper.funcName.Value()
}

func (compareTestWrapper *GenericIndexOfWrapper) Value() interface{} {
	return compareTestWrapper
}

func (compareTestWrapper *GenericIndexOfWrapper) Expected() interface{} {
	return compareTestWrapper.expected
}

func (compareTestWrapper *GenericIndexOfWrapper) ExpectedAsIntArray() []int {
	intArray, isOkay := compareTestWrapper.expected.([]int)

	if isOkay {
		return intArray
	}

	return nil
}

func (compareTestWrapper *GenericIndexOfWrapper) AsTestCaseMessenger() coretests.TestCaseMessenger {
	var testCaseMessenger coretests.TestCaseMessenger = compareTestWrapper

	return testCaseMessenger
}
