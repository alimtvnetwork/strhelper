package splitstestwrapper

import "gitlab.com/auk-go/core/coretests"

type GenericSplitWrapper struct {
	Content          string
	SearchingContent string
	IsCaseSensitive  bool
	Limits           int
	HasPanic         bool
	funcName         coretests.TestFuncName
	expected         interface{}
	actual           interface{}
}

func (genericSplitWrapper *GenericSplitWrapper) Actual() interface{} {
	return genericSplitWrapper.actual
}

func (genericSplitWrapper *GenericSplitWrapper) SetActual(actual interface{}) {
	genericSplitWrapper.actual = actual
}

func (genericSplitWrapper *GenericSplitWrapper) FuncName() string {
	return genericSplitWrapper.funcName.Value()
}

func (genericSplitWrapper *GenericSplitWrapper) Value() interface{} {
	return genericSplitWrapper
}

func (genericSplitWrapper *GenericSplitWrapper) Expected() interface{} {
	return genericSplitWrapper.expected
}

func (genericSplitWrapper *GenericSplitWrapper) ExpectedAsStringsArray() []string {
	intArray, isOkay := genericSplitWrapper.expected.([]string)

	if isOkay {
		return intArray
	}

	return nil
}

func (genericSplitWrapper *GenericSplitWrapper) AsTestCaseMessenger() coretests.TestCaseMessenger {
	var testCaseMessenger coretests.TestCaseMessenger = genericSplitWrapper

	return testCaseMessenger
}
