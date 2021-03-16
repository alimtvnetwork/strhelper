package linestestwrappers

import (
	"gitlab.com/evatix-go/core/coretests"
)

type Compare struct {
	LeftLines                *[]string
	RightLines               *[]string
	StartsAt                 int
	IsCaseSensitive          bool
	IsPanicOnLengthDifferent bool
	funcName                 coretests.TestFuncName
	expected                 int
	actual                   int
}

func (compare Compare) Actual() interface{} {
	return compare.actual
}

func (compare *Compare) SetActual(actual int) {
	compare.actual = actual
}

func (compare Compare) FuncName() string {
	return compare.funcName.Value()
}

func (compare Compare) Value() interface{} {
	return compare
}

func (compare Compare) Expected() interface{} {
	return compare.expected
}
