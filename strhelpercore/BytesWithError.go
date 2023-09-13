package strhelpercore

import (
	"bytes"
	"fmt"

	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/issetter"
	"gitlab.com/auk-go/errorwrapper"
	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"

	"gitlab.com/auk-go/strhelper/anyto"
	"gitlab.com/auk-go/strhelper/internal/whitespacesinternal"
)

type BytesWithError struct {
	bytes        []byte
	lines        []string
	content      corestr.SimpleStringOnce
	errorWrapper *errorwrapper.Wrapper
	bytesLength  int
	stringLength *int // todo fix not ptr
	isWhitespace issetter.Value
}

func NewBytesWithErrorOnlyError(err error) *BytesWithError {
	return &BytesWithError{
		errorWrapper: errorwrapper.NewError(codestack.Skip1, err),
	}
}

func NewBytesWithError(errType errtype.Variation, err error) *BytesWithError {
	return &BytesWithError{
		errorWrapper: errorwrapper.NewUsingError(
			codestack.Skip1,
			errType,
			err),
	}
}

func NewBytesWithErrorWrapper(errWrapper *errorwrapper.Wrapper) *BytesWithError {
	return &BytesWithError{
		errorWrapper: errWrapper,
	}
}

// NewBytesWithErrorUsingAny converts interface object using encoder
// If any is still, still it creates BytesWithError with errorWrapper details and nil contents.
func NewBytesWithErrorUsingAny(any interface{}) *BytesWithError {
	if any == nil {
		return NewBytesWithErrorWrapper(
			errnew.Null.UsingStackSkip(
				codestack.Skip1,
				"has issues with BytesWithError",
				any))
	}

	rawBytes, err := anyto.Bytes(any)
	errWp := errnew.Error.TypeMsgUsingStackSkip(
		codestack.Skip1,
		errtype.ConversionFailed,
		err,
		coredynamic.TypeName(any))

	return &BytesWithError{
		bytes:        rawBytes,
		errorWrapper: errWp,
		bytesLength:  len(rawBytes),
	}
}

// NewBytesWithNoError Creates new BytesWithError
func NewBytesWithNoError(bytes []byte) *BytesWithError {
	return &BytesWithError{
		bytes:        bytes,
		errorWrapper: nil,
		bytesLength:  len(bytes),
	}
}

// Content represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (it *BytesWithError) Content() string {
	if it.content.IsInitialized() {
		return it.content.Value()
	}

	return it.
		content.
		GetPlusSetOnUninitialized(
			it.String())
}

func (it *BytesWithError) BytesLength() int {
	return it.bytesLength
}

func (it *BytesWithError) StringLength() int {
	if it.stringLength == nil {
		length := len([]rune(it.Content()))
		it.stringLength = &length
	}

	return *it.stringLength
}

func (it *BytesWithError) Error() *errorwrapper.Wrapper {
	return it.errorWrapper
}

func (it *BytesWithError) IsNull() bool {
	return it.bytes == nil
}

func (it *BytesWithError) IsNullOrEmpty() bool {
	// checking bytesLength == 0 is enough to prove empty string ""
	// reference : https://play.golang.org/p/6vU5y92LKYg
	return it.bytes == nil ||
		it.bytesLength == 0
}

// IsNullOrEmptyOrWhitespaces returns true if nil or "" or all whitespaces
// (excluding unicode whitespaces, only limited to ASCII spaces)
//
// To check unicode whitespace, Get the String() then use whitespace.IsWhitespaces(...)
func (it *BytesWithError) IsNullOrEmptyOrWhitespaces() bool {
	if it.isWhitespace.IsInitBoolean() {
		return it.isWhitespace.IsTrue()
	}

	isWhitespace := it.bytes == nil ||
		it.bytesLength == 0 ||
		whitespacesinternal.IsAsciiWhitespacesBytes(it.bytes)

	// checking bytesLength == 0 is enough to prove empty string ""
	// reference : https://play.golang.org/p/6vU5y92LKYg
	it.isWhitespace = issetter.GetBool(isWhitespace)

	return it.isWhitespace.IsTrue()
}

// IsDefined
//
// returns true if no currentError and has at least one character other than whitespace (Ascii only)
func (it *BytesWithError) IsDefined() bool {
	return it.errorWrapper.IsEmpty() && !it.IsNullOrEmptyOrWhitespaces()
}

// HasValidCharacters
//
// returns true meaning has at least one character other than whitespace (Ascii only)
func (it *BytesWithError) HasValidCharacters() bool {
	return !it.IsNullOrEmptyOrWhitespaces()
}

// IsEqualAny returns true if bytes contents are same as the any converted bytes contents
//
// Steps :
//  - converts any using the same method as NewBytesWithErrorUsingAny / encoder to bytes (strs.ToBytesOfAny(any))
//  - then compare with bytes
func (it *BytesWithError) IsEqualsAny(
	any interface{}, isPanicOnErrorParse bool,
) bool {
	if it.IsNull() && any == nil {
		return true
	}

	if any == nil {
		return false
	}

	allBytes, err := anyto.Bytes(any)

	if err != nil && isPanicOnErrorParse {
		panic(fmt.Sprintf("%s %s", "any parse failed:", err))
	}

	return bytes.Equal(allBytes, it.bytes)
}

func (it *BytesWithError) IsEqualBytes(anotherRawBytes []byte) bool {
	if it.IsNull() && anotherRawBytes == nil {
		return true
	}

	// both are not nil confirmed, so if any nil returns false.
	if anotherRawBytes == nil || it.IsNull() {
		return false
	}

	return bytes.Equal(it.bytes, anotherRawBytes)
}

func (it *BytesWithError) IsEquals(another *BytesWithError) bool {
	if another == nil {
		return false
	}

	// same pointer
	if it == another {
		return true
	}

	if it.IsNullOrEmpty() == another.IsNullOrEmpty() {
		return true
	}

	if it.BytesLength() != another.BytesLength() {
		return false
	}

	return bytes.Equal(it.bytes, another.bytes)
}

// Bytes represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (it *BytesWithError) Bytes() []byte {
	return it.bytes
}

// Value represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (it *BytesWithError) Value() []byte {
	return it.bytes
}

func (it *BytesWithError) ValueWithoutPtr() []byte {
	return it.bytes
}

// note: that it makes a copy of the content so use it wisely
func (it *BytesWithError) String() string {
	if it.bytesLength == 0 {
		return constants.EmptyString
	}

	return string(it.bytes)
}
