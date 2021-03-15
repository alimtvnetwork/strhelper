package indextestwrappers

import (
	"gitlab.com/evatix-go/core/coretests"
)

type GenericIndexOfTestWrapper struct {
	Content             string
	SearchingContent    string
	InitializedPosition int
	IsCaseSensitive     bool
	Limits              int
	funcName            coretests.TestFuncName
	expected            interface{}
	actual              interface{}
}

func (compareTestWrapper *GenericIndexOfTestWrapper) Actual() interface{} {
	return compareTestWrapper.actual
}

func (compareTestWrapper *GenericIndexOfTestWrapper) SetActual(actual interface{}) {
	compareTestWrapper.actual = actual
}

func (compareTestWrapper *GenericIndexOfTestWrapper) FuncName() string {
	return compareTestWrapper.funcName.Value()
}

func (compareTestWrapper *GenericIndexOfTestWrapper) Value() interface{} {
	return compareTestWrapper
}

func (compareTestWrapper *GenericIndexOfTestWrapper) Expected() interface{} {
	return compareTestWrapper.expected
}

func (compareTestWrapper *GenericIndexOfTestWrapper) ExpectedAsIntArray() *[]int {
	intArray, isOkay := compareTestWrapper.expected.(*[]int)

	if isOkay {
		return intArray
	}

	return nil
}

func (compareTestWrapper *GenericIndexOfTestWrapper) AsTestCaseMessenger() coretests.TestCaseMessenger {
	var testCaseMessenger coretests.TestCaseMessenger = compareTestWrapper

	return testCaseMessenger
}
