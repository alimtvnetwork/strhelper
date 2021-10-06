package splitstestwrapper

import "gitlab.com/evatix-go/core/coretests"

type LastByRunesWrapper struct {
	Content           string
	SearchingContents []rune
	IsCaseSensitive   bool
	Limits            int
	HasPanic          bool
	funcName          coretests.TestFuncName
	expected          interface{}
	actual            interface{}
}

func (lastByRunesWrapper *LastByRunesWrapper) Actual() interface{} {
	return lastByRunesWrapper.actual
}

func (lastByRunesWrapper *LastByRunesWrapper) SetActual(actual interface{}) {
	lastByRunesWrapper.actual = actual
}

func (lastByRunesWrapper *LastByRunesWrapper) FuncName() string {
	return lastByRunesWrapper.funcName.Value()
}

func (lastByRunesWrapper *LastByRunesWrapper) Value() interface{} {
	return lastByRunesWrapper
}

func (lastByRunesWrapper *LastByRunesWrapper) Expected() interface{} {
	return lastByRunesWrapper.expected
}

func (lastByRunesWrapper *LastByRunesWrapper) ExpectedAsStringsArray() []string {
	intArray, isOkay := lastByRunesWrapper.expected.([]string)

	if isOkay {
		return intArray
	}

	return nil
}

func (lastByRunesWrapper *LastByRunesWrapper) AsTestCaseMessenger() coretests.TestCaseMessenger {
	var testCaseMessenger coretests.TestCaseMessenger = lastByRunesWrapper

	return testCaseMessenger
}
