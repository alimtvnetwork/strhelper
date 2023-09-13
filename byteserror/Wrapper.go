package byteserror

import (
	"bytes"
	"encoding/json"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/issetter"
	"gitlab.com/auk-go/errorwrapper"

	"gitlab.com/auk-go/strhelper/encodingbytetype"
	"gitlab.com/auk-go/strhelper/internal/whitespacesinternal"
)

type Wrapper struct {
	bytes        []byte
	content      corestr.SimpleStringOnce
	errorWrapper *errorwrapper.Wrapper
	byteType     encodingbytetype.Variant
	bytesLength  int
	stringLength *int // todo fix , ptr
	isWhitespace issetter.Value
}

func (it *Wrapper) Content() string {
	if it.content.IsInitialized() {
		return it.content.Value()
	}

	return it.content.GetPlusSetOnUninitializedFunc(it.String)
}

func (it *Wrapper) ByteType() encodingbytetype.Variant {
	return it.byteType
}

func (it *Wrapper) BytesLength() int {
	if it == nil {
		return 0
	}

	return it.bytesLength
}

func (it *Wrapper) StringLength() int {
	if it == nil {
		return 0
	}

	if it.stringLength == nil {
		length := len(it.Content())
		it.stringLength = &length
	}

	return *it.stringLength
}

func (it *Wrapper) ErrorWrapper() *errorwrapper.Wrapper {
	return it.errorWrapper
}

func (it *Wrapper) IsNull() bool {
	return it.bytes == nil
}

func (it *Wrapper) IsNullOrEmpty() bool {
	if it == nil {
		return true
	}

	// checking bytesLength == 0 is enough to prove empty string ""
	// reference : https://play.golang.org/p/6vU5y92LKYg
	return it.bytes == nil ||
		it.bytesLength == 0
}

// IsNullOrEmptyOrWhitespaces returns true if nil or "" or all whitespaces
// (excluding unicode whitespaces, only limited to ASCII spaces)
//
// To check unicode whitespace, Get the String() then use whitespace.IsWhitespaces(...)
func (it *Wrapper) IsNullOrEmptyOrWhitespaces() bool {
	if it.isWhitespace.IsUninitialized() {
		isWhitespace := it.bytes == nil ||
			it.bytesLength == 0 ||
			whitespacesinternal.IsAsciiWhitespacesBytes(it.bytes)

		// checking bytesLength == 0 is enough to prove empty string ""
		// reference : https://play.golang.org/p/6vU5y92LKYg
		it.isWhitespace = issetter.GetBool(isWhitespace)
	}

	return it.isWhitespace.IsTrue()
}

// IsDefined returns true if no currentError and has at least one character other than whitespace (Ascii only)
func (it *Wrapper) IsDefined() bool {
	return it.BytesLength() > 0 && it.errorWrapper.IsEmpty() && !it.IsNullOrEmptyOrWhitespaces()
}

// HasValidCharacters returns true meaning has at least one character other than whitespace (Ascii only)
func (it *Wrapper) HasValidCharacters() bool {
	return !it.IsNullOrEmptyOrWhitespaces()
}

func (it *Wrapper) IsEqualBytes(rawBytes []byte) bool {
	if it == nil && rawBytes == nil {
		return true
	}

	if it.IsNull() && rawBytes == nil {
		return true
	}

	return bytes.Equal(
		it.bytes,
		rawBytes)
}

func (it *Wrapper) IsEquals(another *Wrapper) bool {
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

	return bytes.Equal(
		it.bytes,
		another.bytes)
}

// Bytes represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (it *Wrapper) Bytes() []byte {
	return it.bytes
}

// Value represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (it *Wrapper) Value() []byte {
	return it.bytes
}

// note: that it makes a copy of the content so use it wisely
func (it *Wrapper) String() string {
	if it == nil {
		return constants.EmptyString
	}

	return string(it.bytes)
}

func (it *Wrapper) JsonModel() *WrapperDataModel {
	return &WrapperDataModel{
		Bytes:        it.bytes,
		ErrorWrapper: it.errorWrapper,
		ByteType:     it.byteType,
		BytesLength:  it.bytesLength,
		IsWhitespace: it.isWhitespace,
	}
}

func (it *Wrapper) JsonModelAny() interface{} {
	return it.JsonModel()
}

func (it *Wrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(*it.JsonModel())
}

func (it *Wrapper) UnmarshalJSON(data []byte) error {
	var dataModel WrapperDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		it.bytes = dataModel.Bytes
		it.errorWrapper = dataModel.ErrorWrapper
		it.byteType = dataModel.ByteType
		it.bytesLength = dataModel.BytesLength
		it.isWhitespace = dataModel.IsWhitespace
	}

	return err
}

func (it Wrapper) Json() corejson.Result {
	return corejson.New(it)
}

func (it Wrapper) JsonPtr() *corejson.Result {
	return corejson.NewPtr(it)
}

//goland:noinspection GoLinterLocal
func (it *Wrapper) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Wrapper, error) {
	err := jsonResult.Unmarshal(&it)

	if err != nil {
		return nil, err
	}

	return it, nil
}

// Panic if error
//goland:noinspection GoLinterLocal
func (it *Wrapper) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Wrapper {
	newUsingJson, err :=
		it.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (it *Wrapper) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := it.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (it *Wrapper) AsJsonContractsBinder() corejson.JsonContractsBinder {
	return it
}

func (it *Wrapper) AsJsoner() corejson.Jsoner {
	return it
}

func (it *Wrapper) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return it
}

func (it *Wrapper) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}
