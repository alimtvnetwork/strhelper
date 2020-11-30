package strerror

import (
	"fmt"
	"strings"

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

func NewErrorWrapper(error error) ErrorWrapper {
	if error == nil {
		return ErrorWrapper{currentError: nil}
	}

	return ErrorWrapper{currentError: &error}
}

func (errorWrapper *ErrorWrapper) HasError() bool {
	return errorWrapper.currentError != nil && *errorWrapper.currentError != nil
}

func (errorWrapper *ErrorWrapper) ErrorPtr() *error {
	return errorWrapper.currentError
}

func (errorWrapper *ErrorWrapper) Error() error {
	if errorWrapper.IsEmpty() {
		return nil
	}

	return *errorWrapper.currentError
}

// ErrorString if empty error then returns ""
func (errorWrapper *ErrorWrapper) ErrorString() string {
	if errorWrapper.IsEmpty() {
		return ""
	}

	return (*errorWrapper.currentError).Error()
}

func (errorWrapper *ErrorWrapper) IsEmpty() bool {
	return errorWrapper.currentError == nil || *errorWrapper.currentError == nil
}

func (errorWrapper *ErrorWrapper) IsEquals(another *ErrorWrapper) bool {
	if another == nil {
		return false
	}

	if errorWrapper == another {
		return true
	}

	if errorWrapper.IsEmpty() == another.IsEmpty() {
		return true
	}

	// both are not nil confirmed, so if any nil returns false.
	if another.IsEmpty() || errorWrapper.IsEmpty() {
		return false
	}

	// both are defined, now if both pointers are same then it is same object.
	if errorWrapper.Error() == another.Error() {
		return true
	}

	if errorWrapper.ErrorString() == another.ErrorString() {
		return true
	}

	return false
}

func (errorWrapper *ErrorWrapper) IsErrorEquals(err error) bool {
	if err == nil && errorWrapper.IsEmpty() {
		return true
	}

	if err == nil || errorWrapper.IsEmpty() {
		return false
	}

	if errorWrapper.Error() == err {
		return true
	}

	if errorWrapper.ErrorString() == err.Error() {
		return true
	}

	return false
}

// If error IsEmpty then returns false regardless
func (errorWrapper *ErrorWrapper) IsErrorMessage(msg string, isCaseSensitive bool) bool {
	if errorWrapper.IsEmpty() {
		return false
	}

	errMsg := errorWrapper.ErrorString()

	if errMsg == msg {
		// same string or empty or also for case sensitivity
		return true
	}

	if !isCaseSensitive {
		lowerErrorMsg := strings.ToLower(errMsg)
		lowerMsg := strings.ToLower(msg)

		return lowerErrorMsg == lowerMsg
	}

	return false
}

// If error IsEmpty then returns false regardless
func (errorWrapper *ErrorWrapper) IsErrorMessageContains(msg string, isCaseSensitive bool) bool {
	if errorWrapper.IsEmpty() {
		return false
	}

	errMsg := errorWrapper.ErrorString()

	if errMsg == msg && msg == "" {
		// same string or empty or also for case sensitivity
		return true
	}

	if !isCaseSensitive {
		lowerErrorMsg := strings.ToLower(errMsg)
		lowerMsg := strings.ToLower(msg)

		return strings.Index(lowerErrorMsg, lowerMsg) > -1
	}

	return strings.Index(errMsg, msg) > -1
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
