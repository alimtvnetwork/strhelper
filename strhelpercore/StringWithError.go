package strhelpercore

import (
	"encoding/json"
	"errors"
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/issetter"

	"gitlab.com/auk-go/strhelper/internal/isstrinternal"
	"gitlab.com/auk-go/strhelper/internal/whitespacesinternal"
)

type StringWithError struct {
	content      string
	bytes        []byte
	runes        []rune
	error        error
	bytesLength  int
	runeLength   *int
	isWhitespace issetter.Value
}

func NewStringWithErrorOnlyError(err error) *StringWithError {
	return &StringWithError{
		error: err,
	}
}

func NewStringWithError(str string, err error) *StringWithError {
	length := len(str)

	return &StringWithError{
		content:     str,
		error:       err,
		bytesLength: length,
	}
}

func NewStringWithNoError(str string) *StringWithError {
	length := len(str)

	return &StringWithError{
		content:     str,
		bytesLength: length,
	}
}

// ToBytes represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (it *StringWithError) ToBytes() []byte {
	if it.bytes != nil {
		return it.bytes
	}

	it.bytes = []byte(it.content)

	return it.bytes
}

// Returns len(ToRunes()) cached version. If once runeLength generated then it will not generate again.
//
// If don't care about unicode then use len(str) which is BytesLength
// Note :
//  - A bit expensive to generate. However, this one is cached version.
//  - Yields accurate characters length regardless of unicode.
//  - As there is a difference between len(str) and utf8.RuneCountInString(str) or len([]rune(str))
//  - Example : https://play.golang.org/p/78uFF8s-Dw1
func (it *StringWithError) Length() int {
	if it.runeLength != nil {
		return *it.runeLength
	}

	allRunes := it.ToRunes()
	runesLength := len(allRunes)
	it.runeLength = &runesLength

	return *it.runeLength
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
func (it *StringWithError) BytesLength() int {
	return it.bytesLength
}

// ToRunes represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (it *StringWithError) ToRunes() []rune {
	if it.runes == nil {
		it.runes = []rune(it.content)
	}

	return it.runes
}

func (it *StringWithError) IsNull() bool {
	return it == nil
}

func (it *StringWithError) IsNullOrEmpty() bool {
	return it == nil ||
		it.content == constants.EmptyString
}

// IsNullOrEmptyOrWhitespaces
//
// Returns true if nil or "" or all whitespaces (including unicode whitespaces)
func (it *StringWithError) IsNullOrEmptyOrWhitespaces() bool {
	if it.isWhitespace.IsInitBoolean() {
		return it.isWhitespace.IsTrue()
	}

	isWhitespace := it.IsNullOrEmpty()

	if !isWhitespace {
		allRunes := it.ToRunes()
		isWhitespace = whitespacesinternal.IsRunesWhitespaces(allRunes)
	}

	it.isWhitespace = issetter.GetBool(isWhitespace)

	return it.isWhitespace.IsTrue()
}

// IsDefined
//
// Returns true if no currentError and has at least one characters other than whitespace
func (it *StringWithError) IsDefined() bool {
	return it.IsErrorEmpty() && !it.IsNullOrEmptyOrWhitespaces()
}

// HasValidCharacters
//
// Returns true meaning has at least one characters other than whitespace
func (it *StringWithError) HasValidCharacters() bool {
	return !it.IsNullOrEmptyOrWhitespaces()
}

func (it *StringWithError) Error() error {
	return it.error
}

func (it *StringWithError) IsErrorEmpty() bool {
	return it.error == nil
}

func (it *StringWithError) HasError() bool {
	return it.error != nil
}

// HandleError
//
// Only call panic if has currentError
func (it *StringWithError) HandleError() {
	if !it.HasError() {
		return
	}

	message := fmt.Sprintf("%#v", it.error)

	panic(message)
}

// HandleErrorWithMsg
//
// Only call panic if has currentError
func (it *StringWithError) HandleErrorWithMsg(newMessage string) {
	if !it.HasError() {
		return
	}

	message := fmt.Sprintf("%s %#v", newMessage, it.error)

	panic(message)
}

func (it *StringWithError) Value() string {
	return it.content
}

// UnmarshalStringAt usages json unmarshal to convert to object
func (it *StringWithError) UnmarshalStringAt(result interface{}) error {
	return json.Unmarshal(it.ToBytes(), result)
}

// String
//
// note: that it makes a copy of the content so use it wisely
func (it *StringWithError) String() string {
	return it.content
}

func (it *StringWithError) ToByesWithError() *BytesWithError {
	if it.IsNull() {
		err := errors.New("content has nil string pointer and nothing to add")

		return NewBytesWithErrorOnlyError(err)
	}

	return NewBytesWithNoError(it.ToBytes())
}

func (it *StringWithError) IsEquals(another *StringWithError) bool {
	return it.IsEqualsCase(true, another)
}

func (it *StringWithError) IsEqualsCase(isCaseSensitive bool, another *StringWithError) bool {
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

	return isstrinternal.EqualsCase(
		isCaseSensitive,
		it.String(),
		another.String(),
	)
}

func (it *StringWithError) IsStringEquals(another string) bool {
	return it.IsStringCaseEquals(true, another)
}

func (it *StringWithError) IsStringCaseEquals(isCaseSensitive bool, another string) bool {
	return isstrinternal.EqualsCase(
		isCaseSensitive,
		it.String(),
		another,
	)
}
