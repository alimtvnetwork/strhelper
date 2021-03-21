package paddingwrappers

import "gitlab.com/evatix-go/core/coretests"

type RepeatWrapper struct {
	Content     string
	RepeatWidth int
	HasPanic    bool
	funcName    coretests.TestFuncName
	expected    interface{}
	actual      interface{}
}

func (wrapper *RepeatWrapper) Actual() interface{} {
	return wrapper.actual
}

func (wrapper *RepeatWrapper) SetActual(actual interface{}) {
	wrapper.actual = actual
}

func (wrapper *RepeatWrapper) FuncName() string {
	return wrapper.funcName.Value()
}

func (wrapper *RepeatWrapper) Value() interface{} {
	return wrapper
}

func (wrapper *RepeatWrapper) Expected() interface{} {
	return wrapper.expected
}

func (wrapper *RepeatWrapper) AsTestCaseMessenger() coretests.TestCaseMessenger {
	var testCaseMessenger coretests.TestCaseMessenger = wrapper

	return testCaseMessenger
}
