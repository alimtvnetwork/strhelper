package splitstestwrapper

import "gitlab.com/evatix-go/core/coretests"

type GenericSplit struct {
	Content          string
	SearchingContent string
	IsCaseSensitive  bool
	Limits           int
	HasPanic         bool
	funcName         coretests.TestFuncName
	expected         interface{}
	actual           interface{}
}

func (genericSplit *GenericSplit) Actual() interface{} {
	return genericSplit.actual
}

func (genericSplit *GenericSplit) SetActual(actual interface{}) {
	genericSplit.actual = actual
}

func (genericSplit *GenericSplit) FuncName() string {
	return genericSplit.funcName.Value()
}

func (genericSplit *GenericSplit) Value() interface{} {
	return genericSplit
}

func (genericSplit *GenericSplit) Expected() interface{} {
	return genericSplit.expected
}

func (genericSplit *GenericSplit) ExpectedAsStringsArray() *[]string {
	intArray, isOkay := genericSplit.expected.(*[]string)

	if isOkay {
		return intArray
	}

	return nil
}

func (genericSplit *GenericSplit) AsTestCaseMessenger() coretests.TestCaseMessenger {
	var testCaseMessenger coretests.TestCaseMessenger = genericSplit

	return testCaseMessenger
}
