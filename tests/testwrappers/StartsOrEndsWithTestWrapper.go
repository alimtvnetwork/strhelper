package testwrappers

import (
	"gitlab.com/evatix-go/core/coretests"
)

type StartsOrEndsWithTestWrapper struct {
	WholeText       string
	Search          string
	StartsAt        int
	IsCaseSensitive bool
	funcName        coretests.TestFuncName
	expected        bool
	actual          bool
}

func (startsOrEndsWithTestWrapper StartsOrEndsWithTestWrapper) Actual() interface{} {
	return startsOrEndsWithTestWrapper.actual
}

func (startsOrEndsWithTestWrapper *StartsOrEndsWithTestWrapper) SetActual(actual bool) {
	startsOrEndsWithTestWrapper.actual = actual
}

func (startsOrEndsWithTestWrapper StartsOrEndsWithTestWrapper) FuncName() string {
	return startsOrEndsWithTestWrapper.funcName.Value()
}

func (startsOrEndsWithTestWrapper StartsOrEndsWithTestWrapper) Value() interface{} {
	return startsOrEndsWithTestWrapper
}

func (startsOrEndsWithTestWrapper StartsOrEndsWithTestWrapper) Expected() interface{} {
	return startsOrEndsWithTestWrapper.expected
}
