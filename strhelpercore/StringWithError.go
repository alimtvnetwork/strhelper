package strhelpercore

import (
	"encoding/json"
	"errors"
	"fmt"

	"gitlab.com/evatix-go/strhelper/internal/pkg/isinternal"
	"gitlab.com/evatix-go/strhelper/internal/pkg/whitespacesinternal"
	"gitlab.com/evatix-go/strhelper/strconst"
)

type StringWithError struct {
	content      *string
	bytes        *[]byte
	runes        *[]rune
	error        *error
	bytesLength  int
	runeLength   *int
	isWhitespace *bool
}

func NewStringWithErrorOnlyError(err *error) *StringWithError {
	return &StringWithError{
		content: nil,
		error:   err,
	}
}

func NewStringWithError(str *string, err *error) *StringWithError {
	length := 0

	if str != nil {
		length = len(*str)
	}

	return &StringWithError{
		content:     str,
		error:       err,
		bytesLength: length,
	}
}

func NewStringWithNoError(str *string) *StringWithError {
	length := 0

	if str != nil {
		length = len(*str)
	}

	return &StringWithError{
		content:     str,
		error:       nil,
		bytesLength: length,
	}
}

// ToBytesPtr represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (stringWithError *StringWithError) ToBytesPtr() *[]byte {
	if stringWithError.bytes != nil {
		return stringWithError.bytes
	}

	bytes := []byte(*stringWithError.content)
	stringWithError.bytes = &bytes

	return stringWithError.bytes
}

// Returns len(ToRunesPtr()) cached version. If once runeLength generated then it will not generate again.
//
// If don't care about unicode then use len(str) which is BytesLength
// Note :
//  - A bit expensive to generate. However, this one is cached version.
//  - Yields accurate characters length regardless of unicode.
//  - As there is a difference between len(str) and utf8.RuneCountInString(str) or len([]rune(str))
//  - Example : https://play.golang.org/p/78uFF8s-Dw1
func (stringWithError *StringWithError) Length() int {
	if stringWithError.runeLength == nil {
		allRunes := stringWithError.ToRunesPtr()
		runesLength := len(*allRunes)
		stringWithError.runeLength = &runesLength
	}

	return *stringWithError.runeLength
}

// There is a difference between length in bytes (doesn't represent proper unicode chars) and
//
// length in runes (represents actual char length in unicode format).
//
// If care about unicode chars count then use Length version.
// It returns the len(str) cached version.
// Note :
//  - This version returns cached version of len(str) which is saved during the instantiation of the object creation.
//  - Less expensive (returns StringWithError.bytesLength). Effective for bytes knowledge only.
//  - Doesn't yield accurate characters length of unicode characters but only ascii.
//  - There is a difference between len(str) and utf8.RuneCountInString(str) or len([]rune(str)).
//  - Example : https://play.golang.org/p/78uFF8s-Dw1
func (stringWithError *StringWithError) BytesLength() int {
	return stringWithError.bytesLength
}

// ToRunesPtr represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (stringWithError *StringWithError) ToRunesPtr() *[]rune {
	if stringWithError.runes == nil {
		allRunes := []rune(*stringWithError.content)
		stringWithError.runes = &allRunes
	}

	return stringWithError.runes
}

func (stringWithError *StringWithError) IsNull() bool {
	return stringWithError.content == nil
}

func (stringWithError *StringWithError) IsNullOrEmpty() bool {
	return stringWithError.content == nil ||
		*stringWithError.content == strconst.EmptyString
}

// Returns true if nil or "" or all whitespaces (including unicode whitespaces)
func (stringWithError *StringWithError) IsNullOrEmptyOrWhitespaces() bool {
	if stringWithError.isWhitespace == nil {
		isWhitespace := stringWithError.content == nil ||
			*stringWithError.content == strconst.EmptyString

		if !isWhitespace {
			allRunes := stringWithError.ToRunesPtr()
			isWhitespace = whitespacesinternal.IsRunesWhitespaces(allRunes)
		}

		stringWithError.isWhitespace = &isWhitespace
	}

	return *stringWithError.isWhitespace
}

// Returns true if no currentError and has at least one characters other than whitespace
func (stringWithError *StringWithError) IsDefined() bool {
	return stringWithError.IsErrorEmpty() && !stringWithError.IsNullOrEmptyOrWhitespaces()
}

// Returns true meaning has at least one characters other than whitespace
func (stringWithError *StringWithError) HasValidCharacters() bool {
	return !stringWithError.IsNullOrEmptyOrWhitespaces()
}

func (stringWithError *StringWithError) Error() *error {
	return stringWithError.error
}

func (stringWithError *StringWithError) IsErrorEmpty() bool {
	return stringWithError.error == nil
}

func (stringWithError *StringWithError) HasError() bool {
	return stringWithError.error != nil
}

// Only call panic if has currentError
func (stringWithError *StringWithError) HandleError() {
	if !stringWithError.HasError() {
		return
	}

	message := fmt.Sprintf("%#v", *stringWithError.error)

	panic(message)
}

// Only call panic if has currentError
func (stringWithError *StringWithError) HandleErrorWithMsg(newMessage string) {
	if !stringWithError.HasError() {
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

// ToType usages json unmarshal to convert to object
func (stringWithError *StringWithError) ToType(result *interface{}) error {
	return json.Unmarshal(*stringWithError.ToBytesPtr(), result)
}

// note: that it makes a copy of the content so use it wisely
func (stringWithError *StringWithError) String() string {
	return *stringWithError.content
}

func (stringWithError *StringWithError) ToByesWithError() *BytesWithError {
	if stringWithError.IsNull() {
		err := errors.New("content has nil string pointer and nothing to add")

		return NewBytesWithErrorOnlyError(err)
	}

	return NewBytesWithNoError(stringWithError.ToBytesPtr())
}

func (stringWithError *StringWithError) StringPtr() *string {
	return stringWithError.content
}

func (stringWithError *StringWithError) IsEquals(another *StringWithError) bool {
	return stringWithError.IsEqualsCase(another, true)
}

func (stringWithError *StringWithError) IsEqualsCase(another *StringWithError, isCaseSensitive bool) bool {
	if another == nil {
		return false
	}

	// same pointer
	if stringWithError == another {
		return true
	}

	if stringWithError.IsNullOrEmpty() == another.IsNullOrEmpty() {
		return true
	}

	if stringWithError.BytesLength() != another.BytesLength() {
		return false
	}

	return isinternal.EqualsCasePtr(
		stringWithError.StringPtr(),
		another.StringPtr(),
		isCaseSensitive)
}

func (stringWithError *StringWithError) IsStringEquals(another *string) bool {
	return stringWithError.IsStringCaseEquals(another, true)
}

func (stringWithError *StringWithError) IsStringCaseEquals(another *string, isCaseSensitive bool) bool {
	if stringWithError.IsNull() && another == nil {
		return true
	}

	if stringWithError.IsNull() || another == nil {
		return false
	}

	// same pointer
	if stringWithError.content == another {
		return true
	}

	if stringWithError.BytesLength() != len(*another) {
		return false
	}

	return isinternal.EqualsCasePtr(
		stringWithError.StringPtr(),
		another,
		isCaseSensitive)
}
