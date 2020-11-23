package strhelpercore

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/strconst"
)

type ErrorWrapper struct {
	currentError *error
	fullMsg      *string
	msg          *string
	typeName     *string
}

func NewErrorWrapperPtr(error *error) *ErrorWrapper {
	return &ErrorWrapper{currentError: error}
}

func NewErrorWrapper(error error) *ErrorWrapper {
	if error == nil {
		return &ErrorWrapper{currentError: nil}
	}

	return &ErrorWrapper{currentError: &error}
}

func (errorWrapper *ErrorWrapper) HasError() bool {
	return errorWrapper.currentError != nil && *errorWrapper.currentError != nil
}

func (errorWrapper *ErrorWrapper) Error() *error {
	return errorWrapper.currentError
}

func (errorWrapper *ErrorWrapper) IsErrorEmpty() bool {
	return errorWrapper.currentError == nil || *errorWrapper.currentError == nil
}

// Only call panic if has currentError
func (errorWrapper *ErrorWrapper) HandleError() {
	if !(*errorWrapper).HasError() {
		return
	}

	panic(errorWrapper.FullString())
}

// Only call panic if has currentError
func (errorWrapper *ErrorWrapper) HandleErrorWithMsg(newMessage string) {
	if !(*errorWrapper).HasError() {
		return
	}

	message := fmt.Sprintf("%s %s", newMessage, *errorWrapper.FullString())

	panic(message)
}

func (errorWrapper *ErrorWrapper) Value() *error {
	return errorWrapper.currentError
}

func (errorWrapper *ErrorWrapper) FullString() *string {
	if errorWrapper.fullMsg == nil {
		msg := fmt.Sprintf(strconst.SprintFullPropertyNameValueFormat, *errorWrapper.currentError)
		errorWrapper.fullMsg = &msg
	}

	return errorWrapper.fullMsg
}

func (errorWrapper *ErrorWrapper) TypeString() *string {
	if errorWrapper.typeName == nil {
		msg := fmt.Sprintf(strconst.SprintTypeFormat, *errorWrapper.currentError)
		errorWrapper.typeName = &msg
	}

	return errorWrapper.typeName
}

func (errorWrapper *ErrorWrapper) StringPtr() *string {
	if errorWrapper.msg == nil {
		msg := fmt.Sprintf(strconst.SprintTypeFormat, *errorWrapper.currentError)
		errorWrapper.msg = &msg
	}

	return errorWrapper.msg
}

func (errorWrapper *ErrorWrapper) String() string {
	return *errorWrapper.StringPtr()
}
