package splitstestwrapper

import "gitlab.com/evatix-go/core/coretests"

type LastByRunes struct {
	Content           string
	SearchingContents []string
	IsCaseSensitive   bool
	Limits            int
	HasPanic          bool
	funcName          coretests.TestFuncName
	expected          interface{}
	actual            interface{}
}

func (lastByRunes *LastByRunes) Actual() interface{} {
	return lastByRunes.actual
}

func (lastByRunes *LastByRunes) SetActual(actual interface{}) {
	lastByRunes.actual = actual
}

func (lastByRunes *LastByRunes) FuncName() string {
	return lastByRunes.funcName.Value()
}

func (lastByRunes *LastByRunes) Value() interface{} {
	return lastByRunes
}

func (lastByRunes *LastByRunes) Expected() interface{} {
	return lastByRunes.expected
}

func (lastByRunes *LastByRunes) ExpectedAsStringsArray() *[]string {
	intArray, isOkay := lastByRunes.expected.(*[]string)

	if isOkay {
		return intArray
	}

	return nil
}

func (lastByRunes *LastByRunes) AsTestCaseMessenger() coretests.TestCaseMessenger {
	var testCaseMessenger coretests.TestCaseMessenger = lastByRunes

	return testCaseMessenger
}
