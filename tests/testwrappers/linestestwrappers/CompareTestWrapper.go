package linestestwrappers

import (
	"gitlab.com/evatix-go/core/coretests"
)

type CompareTestWrapper struct {
	LeftLines                *[]string
	RightLines               *[]string
	StartsAt                 int
	IsCaseSensitive          bool
	IsPanicOnLengthDifferent bool
	funcName                 coretests.TestFuncName
	expected                 int
	actual                   int
}

func (compareTestWrapper CompareTestWrapper) Actual() interface{} {
	return compareTestWrapper.actual
}

func (compareTestWrapper *CompareTestWrapper) SetActual(actual int) {
	compareTestWrapper.actual = actual
}

func (compareTestWrapper CompareTestWrapper) FuncName() string {
	return compareTestWrapper.funcName.Value()
}

func (compareTestWrapper CompareTestWrapper) Value() interface{} {
	return compareTestWrapper
}

func (compareTestWrapper CompareTestWrapper) Expected() interface{} {
	return compareTestWrapper.expected
}
