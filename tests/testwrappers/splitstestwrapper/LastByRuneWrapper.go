package splitstestwrapper

import "gitlab.com/auk-go/core/coretests"

type LastByRuneWrapper struct {
	Content          string
	SearchingContent rune
	IsCaseSensitive  bool
	Limits           int
	HasPanic         bool
	funcName         coretests.TestFuncName
	expected         interface{}
	actual           interface{}
}

func (lastByRuneWrapper *LastByRuneWrapper) Actual() interface{} {
	return lastByRuneWrapper.actual
}

func (lastByRuneWrapper *LastByRuneWrapper) SetActual(actual interface{}) {
	lastByRuneWrapper.actual = actual
}

func (lastByRuneWrapper *LastByRuneWrapper) FuncName() string {
	return lastByRuneWrapper.funcName.Value()
}

func (lastByRuneWrapper *LastByRuneWrapper) Value() interface{} {
	return lastByRuneWrapper
}

func (lastByRuneWrapper *LastByRuneWrapper) Expected() interface{} {
	return lastByRuneWrapper.expected
}

func (lastByRuneWrapper *LastByRuneWrapper) ExpectedAsStringsArray() []string {
	intArray, isOkay := lastByRuneWrapper.expected.([]string)

	if isOkay {
		return intArray
	}

	return nil
}

func (lastByRuneWrapper *LastByRuneWrapper) AsTestCaseMessenger() coretests.TestCaseMessenger {
	var testCaseMessenger coretests.TestCaseMessenger = lastByRuneWrapper

	return testCaseMessenger
}
