package indextestwrappers

import (
	"gitlab.com/evatix-go/core/coretests"
)

type GenericIndexOf struct {
	Content             string
	SearchingContent    string
	InitializedPosition int
	IsCaseSensitive     bool
	Limits              int
	HasPanic            bool
	funcName            coretests.TestFuncName
	expected            interface{}
	actual              interface{}
}

func (compareTestWrapper *GenericIndexOf) Actual() interface{} {
	return compareTestWrapper.actual
}

func (compareTestWrapper *GenericIndexOf) SetActual(actual interface{}) {
	compareTestWrapper.actual = actual
}

func (compareTestWrapper *GenericIndexOf) FuncName() string {
	return compareTestWrapper.funcName.Value()
}

func (compareTestWrapper *GenericIndexOf) Value() interface{} {
	return compareTestWrapper
}

func (compareTestWrapper *GenericIndexOf) Expected() interface{} {
	return compareTestWrapper.expected
}

func (compareTestWrapper *GenericIndexOf) ExpectedAsIntArray() *[]int {
	intArray, isOkay := compareTestWrapper.expected.(*[]int)

	if isOkay {
		return intArray
	}

	return nil
}

func (compareTestWrapper *GenericIndexOf) AsTestCaseMessenger() coretests.TestCaseMessenger {
	var testCaseMessenger coretests.TestCaseMessenger = compareTestWrapper

	return testCaseMessenger
}
