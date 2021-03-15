package indextestwrappers

import (
	"gitlab.com/evatix-go/core/coretests"
)

type GenericIndexOfTestWrapper struct {
	Content                  string
	SearchingContent         string
	InitializedPosition      int
	IsCaseSensitive          bool
	IsPanicOnLengthDifferent bool
	Limits                   int
	funcName                 coretests.TestFuncName
	expected                 int
	actual                   interface{}
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
