package strhelpercore

import (
	"gitlab.com/evatix-go/strhelper/internal/pkg/misc"
	"gitlab.com/evatix-go/strhelper/internal/pkg/whitespacesinternal"
	"gitlab.com/evatix-go/strhelper/strconst"
)

type BytesWithError struct {
	bytes        *[]byte
	lines        *[]string
	content      *string
	error        *ErrorWrapper
	bytesLength  int
	stringLength *int
	isWhitespace *bool
}

func NewBytesWithErrorOnlyError(err *error) *BytesWithError {
	return &BytesWithError{
		error: NewErrorWrapperPtr(err),
	}
}

func NewBytesWithError(bytes *[]byte, err *error) *BytesWithError {
	length := 0

	if bytes != nil {
		length = len(*bytes)
	}

	return &BytesWithError{
		bytes:       bytes,
		error:       NewErrorWrapperPtr(err),
		bytesLength: length,
	}
}

func NewBytesWithErrorUsingAny(any interface{}) *BytesWithError {
	length := 0

	bytes := misc.ToBytesOfAny(any)

	if bytes != nil {
		length = len(*bytes)
	}

	return &BytesWithError{
		bytes:       bytes,
		error:       NewErrorWrapperPtr(nil),
		bytesLength: length,
	}
}

// NewBytesWithNoError Creates new BytesWithError
func NewBytesWithNoError(bytes *[]byte) *BytesWithError {
	return NewBytesWithError(bytes, nil)
}

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

func (bytesWithError *BytesWithError) Error() *ErrorWrapper {
	return bytesWithError.error
}

func (bytesWithError *BytesWithError) IsNull() bool {
	return bytesWithError.bytes == nil
}

func (bytesWithError *BytesWithError) IsNullOrEmpty() bool {
	return bytesWithError.bytes == nil ||
		*bytesWithError.StringPtr() == strconst.EmptyString
}

// IsNullOrEmptyOrWhitespaces returns true if nil or "" or all whitespaces (including unicode whitespaces)
func (bytesWithError *BytesWithError) IsNullOrEmptyOrWhitespaces() bool {
	if bytesWithError.isWhitespace == nil {
		str := bytesWithError.StringPtr()
		*bytesWithError.isWhitespace = str == nil ||
			*str == strconst.EmptyString ||
			whitespacesinternal.IsWhitespaces(str)
	}

	return *bytesWithError.isWhitespace
}

// IsDefined returns true if no currentError and has at least one characters other than whitespace
func (bytesWithError *BytesWithError) IsDefined() bool {
	return bytesWithError.error.IsErrorEmpty() && !bytesWithError.IsNullOrEmptyOrWhitespaces()
}

// HasValidCharacters returns true meaning has at least one characters other than whitespace
func (bytesWithError *BytesWithError) HasValidCharacters() bool {
	return !bytesWithError.IsNullOrEmptyOrWhitespaces()
}

func (bytesWithError *BytesWithError) Bytes() *[]byte {
	return bytesWithError.bytes
}

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

func (bytesWithError *BytesWithError) StringPtr() *string {
	if bytesWithError.content == nil && bytesWithError.bytes != nil {
		newString := string(*bytesWithError.bytes)
		bytesWithError.content = &newString
	}

	return bytesWithError.content
}
