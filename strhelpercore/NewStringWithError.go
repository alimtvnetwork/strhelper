package strhelpercore

import "fmt"

type StringWithError struct {
	content *string
	error   *error
}

func NewStringWithError(str *string, err *error) *StringWithError {
	return &StringWithError{
		content: str,
		error:   err,
	}
}

func NewStringWithNoError(str *string) *StringWithError {
	return &StringWithError{
		content: str,
		error:   nil,
	}
}

func (stringWithError *StringWithError) Error() *error {
	return stringWithError.error
}

func (stringWithError *StringWithError) HasError() bool {
	return stringWithError.error != nil && (*stringWithError).error != nil
}

// Only call panic if has error
func (stringWithError *StringWithError) HandleError(newMessage string) {
	if !(*stringWithError).HasError() {
		return
	}

	message := fmt.Sprintf("%s %#v", newMessage, *stringWithError.error)

	panic(message)
}

func (stringWithError *StringWithError) Value() *string {
	return stringWithError.content
}

func (stringWithError *StringWithError) String() string {
	return *stringWithError.content
}

func (stringWithError *StringWithError) StringPtr() *string {
	return stringWithError.content
}
