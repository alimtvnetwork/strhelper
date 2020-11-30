package strhelpercore

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/anyto"
	"gitlab.com/evatix-go/strhelper/internal/pkg/misc"
	"gitlab.com/evatix-go/strhelper/internal/pkg/whitespacesinternal"
	"gitlab.com/evatix-go/strhelper/strerror"
)

type BytesWithError struct {
	bytes        *[]byte
	lines        *[]string
	content      *string
	errorWrapper strerror.ErrorWrapper
	bytesLength  int
	stringLength *int
	isWhitespace *bool
}

func NewBytesWithErrorOnlyError(err error) *BytesWithError {
	return &BytesWithError{
		errorWrapper: *strerror.NewErrorWrapperPtr(&err),
	}
}

func NewBytesWithErrorOnlyErrorPtr(err *error) *BytesWithError {
	return &BytesWithError{
		errorWrapper: *strerror.NewErrorWrapperPtr(err),
	}
}

func NewBytesWithError(bytes *[]byte, err *error) *BytesWithError {
	length := 0

	if bytes != nil {
		length = len(*bytes)
	}

	return &BytesWithError{
		bytes:        bytes,
		errorWrapper: *strerror.NewErrorWrapperPtr(err),
		bytesLength:  length,
	}
}

// NewBytesWithErrorUsingAny converts interface object using encoder
// If any is still, still it creates BytesWithError with errorWrapper details and nil contents.
func NewBytesWithErrorUsingAny(any interface{}) *BytesWithError {
	length := 0

	if any == nil {
		return NewBytesWithErrorOnlyErrorPtr(nil)
	}

	bytes, err := anyto.Bytes(any)

	if bytes != nil && *bytes != nil {
		length = len(*bytes)
	}

	return &BytesWithError{
		bytes:        bytes,
		errorWrapper: strerror.NewErrorWrapper(err),
		bytesLength:  length,
	}
}

// NewBytesWithNoError Creates new BytesWithError
func NewBytesWithNoError(bytes *[]byte) *BytesWithError {
	return NewBytesWithError(bytes, nil)
}

// ContentAsString represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (bytesWithError *BytesWithError) ContentAsString() *string {
	return bytesWithError.StringPtr()
}

func (bytesWithError *BytesWithError) BytesLength() int {
	return bytesWithError.bytesLength
}

func (bytesWithError *BytesWithError) StringLength() int {
	if bytesWithError.stringLength == nil {
		length := len(*bytesWithError.StringPtr())
		bytesWithError.stringLength = &length
	}

	return *bytesWithError.stringLength
}

// Error must be initialize have it or not. Then check Error().IsEmpty()
func (bytesWithError *BytesWithError) Error() strerror.ErrorWrapper {
	return bytesWithError.errorWrapper
}

func (bytesWithError *BytesWithError) IsNull() bool {
	return bytesWithError.bytes == nil
}

func (bytesWithError *BytesWithError) IsNullOrEmpty() bool {
	// checking bytesLength == 0 is enough to prove empty string ""
	// reference : https://play.golang.org/p/6vU5y92LKYg
	return bytesWithError.bytes == nil ||
		bytesWithError.bytesLength == 0
}

// IsNullOrEmptyOrWhitespaces returns true if nil or "" or all whitespaces
// (excluding unicode whitespaces, only limited to ASCII spaces)
//
// To check unicode whitespace, Get the String() then use whitespace.IsWhitespaces(...)
func (bytesWithError *BytesWithError) IsNullOrEmptyOrWhitespaces() bool {
	if bytesWithError.isWhitespace == nil {
		// checking bytesLength == 0 is enough to prove empty string ""
		// reference : https://play.golang.org/p/6vU5y92LKYg
		*bytesWithError.isWhitespace = bytesWithError.bytes == nil ||
			bytesWithError.bytesLength == 0 ||
			whitespacesinternal.IsAsciiWhitespacesBytes(bytesWithError.bytes)
	}

	return *bytesWithError.isWhitespace
}

// IsDefined returns true if no currentError and has at least one characters other than whitespace (Ascii only)
func (bytesWithError *BytesWithError) IsDefined() bool {
	return bytesWithError.errorWrapper.IsEmpty() && !bytesWithError.IsNullOrEmptyOrWhitespaces()
}

// HasValidCharacters returns true meaning has at least one characters other than whitespace (Ascii only)
func (bytesWithError *BytesWithError) HasValidCharacters() bool {
	return !bytesWithError.IsNullOrEmptyOrWhitespaces()
}

// IsEqualAny returns true if bytes contents are same as the any converted bytes contents
//
// Steps :
//  - converts any using the same method as NewBytesWithErrorUsingAny / encoder to bytes (strs.ToBytesOfAny(any))
//  - then compare with bytes
func (bytesWithError *BytesWithError) IsEqualsAny(any interface{}, isPanicOnErrorParse bool) bool {
	if bytesWithError.IsNull() && any == nil {
		return true
	}

	if any == nil {
		return false
	}

	bytes, err := anyto.Bytes(any)

	if err != nil && isPanicOnErrorParse {
		panic(fmt.Sprintf("%s %s", "any parse failed:", err))
	}

	return misc.IsBytesEquals(
		bytesWithError.bytes,
		bytes,
		0)
}

func (bytesWithError *BytesWithError) IsEqualBytes(bytes *[]byte) bool {
	if bytesWithError.IsNull() && bytes == nil {
		return true
	}

	// both are not nil confirmed, so if any nil returns false.
	if bytes == nil || bytesWithError.IsNull() {
		return false
	}

	return misc.IsBytesEquals(
		bytesWithError.bytes,
		bytes,
		0)
}

func (bytesWithError *BytesWithError) IsEquals(another *BytesWithError) bool {
	if another == nil {
		return false
	}

	// same pointer
	if bytesWithError == another {
		return true
	}

	if bytesWithError.IsNullOrEmpty() == another.IsNullOrEmpty() {
		return true
	}

	if bytesWithError.BytesLength() != another.BytesLength() {
		return false
	}

	return misc.IsBytesEquals(
		bytesWithError.bytes,
		another.bytes,
		0)
}

// Bytes represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (bytesWithError *BytesWithError) Bytes() *[]byte {
	return bytesWithError.bytes
}

// Value represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (bytesWithError *BytesWithError) Value() *[]byte {
	return bytesWithError.bytes
}

// note: that it makes a copy of the content so use it wisely
func (bytesWithError *BytesWithError) ValueWithoutPtr() []byte {
	return *bytesWithError.bytes
}

// note: that it makes a copy of the content so use it wisely
func (bytesWithError *BytesWithError) String() string {
	return *bytesWithError.StringPtr()
}

// StringPtr is an expensive operation, it creates new memory using string(*bytesWithError.bytes)
// However, it does it at once, so calling it 3 times will only create once and cached result will be returned.
//
// StringPtr represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (bytesWithError *BytesWithError) StringPtr() *string {
	if bytesWithError.content == nil && bytesWithError.bytes != nil {
		newString := string(*bytesWithError.bytes)
		bytesWithError.content = &newString
	}

	return bytesWithError.content
}
