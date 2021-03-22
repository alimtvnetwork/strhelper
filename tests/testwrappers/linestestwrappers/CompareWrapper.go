package linestestwrappers

import (
	"gitlab.com/evatix-go/core/coretests"
)

type CompareWrapper struct {
	LeftLines                *[]string
	RightLines               *[]string
	StartsAt                 int
	IsCaseSensitive          bool
	IsPanicOnLengthDifferent bool
	funcName                 coretests.TestFuncName
	expected                 int
	actual                   int
}

func (compareWrapper CompareWrapper) Actual() interface{} {
	return compareWrapper.actual
}

func (compareWrapper *CompareWrapper) SetActual(actual int) {
	compareWrapper.actual = actual
}

func (compareWrapper CompareWrapper) FuncName() string {
	return compareWrapper.funcName.Value()
}

func (compareWrapper CompareWrapper) Value() interface{} {
	return compareWrapper
}

func (compareWrapper CompareWrapper) Expected() interface{} {
	return compareWrapper.expected
}
