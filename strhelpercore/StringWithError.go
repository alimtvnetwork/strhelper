package strhelpercore

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/constants"
)

type StringWithError struct {
	content      *string
	error        *error
	isWhitespace *bool
}

func NewStringWithErrorOnlyError(err *error) *StringWithError {
	return &StringWithError{
		content: nil,
		error:   err,
	}
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

func (stringWithError *StringWithError) IsNull() bool {
	return stringWithError.content == nil || (*stringWithError).content == nil
}

func (stringWithError *StringWithError) IsNullOrEmpty() bool {
	return stringWithError.content == nil ||
		(*stringWithError).content == nil ||
		*stringWithError.content == constants.EmptyString
}

// Returns true if nil or "" or all whitespaces (including unicode whitespaces)
func (stringWithError *StringWithError) IsNullOrEmptyOrWhitespaces() bool {
	if stringWithError.isWhitespace == nil {
		*stringWithError.isWhitespace = stringWithError.content == nil ||
			(*stringWithError).content == nil ||
			*stringWithError.content == constants.EmptyString ||
			isWhitespaces(stringWithError.content)
	}

	return *(*stringWithError).isWhitespace
}

// Returns true if no error and has at least one characters other than whitespace
func (stringWithError *StringWithError) IsDefined() bool {
	return (*stringWithError).IsErrorEmpty() && !(*stringWithError).IsNullOrEmptyOrWhitespaces()
}

// Returns true meaning has at least one characters other than whitespace
func (stringWithError *StringWithError) HasValidCharacters() bool {
	return !(*stringWithError).IsNullOrEmptyOrWhitespaces()
}

func (stringWithError *StringWithError) Error() *error {
	return stringWithError.error
}

func (stringWithError *StringWithError) IsErrorEmpty() bool {
	return stringWithError.error == nil || (*stringWithError).error == nil
}

func (stringWithError *StringWithError) HasError() bool {
	return stringWithError.error != nil && (*stringWithError).error != nil
}

// Only call panic if has error
func (stringWithError *StringWithError) HandleError() {
	if !(*stringWithError).HasError() {
		return
	}

	message := fmt.Sprintf("%#v", *stringWithError.error)

	panic(message)
}

// Only call panic if has error
func (stringWithError *StringWithError) HandleErrorMsg(newMessage string) {
	if !(*stringWithError).HasError() {
		return
	}

	message := fmt.Sprintf("%s %#v", newMessage, *stringWithError.error)

	panic(message)
}

func (stringWithError *StringWithError) Value() *string {
	return stringWithError.content
}

// note: that it makes a copy of the content so use it wisely
func (stringWithError *StringWithError) ValueWithoutPtr() string {
	return *stringWithError.content
}

// note: that it makes a copy of the content so use it wisely
func (stringWithError *StringWithError) String() string {
	return *stringWithError.content
}

func (stringWithError *StringWithError) StringPtr() *string {
	return stringWithError.content
}
