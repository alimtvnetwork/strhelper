package byteserror

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/issetter"
	"gitlab.com/evatix-go/errorwrapper"

	"gitlab.com/evatix-go/strhelper/encodingbytetype"
	"gitlab.com/evatix-go/strhelper/internal/misc"
	"gitlab.com/evatix-go/strhelper/internal/whitespacesinternal"
)

type Wrapper struct {
	bytes        *[]byte
	content      *string
	errorWrapper *errorwrapper.Wrapper
	byteType     encodingbytetype.Variant
	bytesLength  int
	stringLength *int
	isWhitespace issetter.Value
}

func (wrapper *Wrapper) ByteType() encodingbytetype.Variant {
	return wrapper.byteType
}

// ContentAsString represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (wrapper *Wrapper) ContentAsString() *string {
	return wrapper.StringPtr()
}

func (wrapper *Wrapper) BytesLength() int {
	return wrapper.bytesLength
}

func (wrapper *Wrapper) StringLength() int {
	if wrapper.stringLength == nil {
		length := len(*wrapper.StringPtr())
		wrapper.stringLength = &length
	}

	return *wrapper.stringLength
}

func (wrapper *Wrapper) ErrorWrapper() *errorwrapper.Wrapper {
	return wrapper.errorWrapper
}

func (wrapper *Wrapper) IsNull() bool {
	return wrapper.bytes == nil
}

func (wrapper *Wrapper) IsNullOrEmpty() bool {
	// checking bytesLength == 0 is enough to prove empty string ""
	// reference : https://play.golang.org/p/6vU5y92LKYg
	return wrapper.bytes == nil ||
		wrapper.bytesLength == 0
}

// IsNullOrEmptyOrWhitespaces returns true if nil or "" or all whitespaces
// (excluding unicode whitespaces, only limited to ASCII spaces)
//
// To check unicode whitespace, Get the String() then use whitespace.IsWhitespaces(...)
func (wrapper *Wrapper) IsNullOrEmptyOrWhitespaces() bool {
	if wrapper.isWhitespace.IsUninitialized() {
		isWhitespace := wrapper.bytes == nil ||
			wrapper.bytesLength == 0 ||
			whitespacesinternal.IsAsciiWhitespacesBytes(wrapper.bytes)

		// checking bytesLength == 0 is enough to prove empty string ""
		// reference : https://play.golang.org/p/6vU5y92LKYg
		wrapper.isWhitespace = issetter.GetBool(isWhitespace)
	}

	return wrapper.isWhitespace.IsTrue()
}

// IsDefined returns true if no currentError and has at least one characters other than whitespace (Ascii only)
func (wrapper *Wrapper) IsDefined() bool {
	return wrapper.errorWrapper.IsEmpty() && !wrapper.IsNullOrEmptyOrWhitespaces()
}

// HasValidCharacters returns true meaning has at least one characters other than whitespace (Ascii only)
func (wrapper *Wrapper) HasValidCharacters() bool {
	return !wrapper.IsNullOrEmptyOrWhitespaces()
}

func (wrapper *Wrapper) IsEqualBytes(bytes *[]byte) bool {
	if wrapper.IsNull() && bytes == nil {
		return true
	}

	// both are not nil confirmed, so if any nil returns false.
	if bytes == nil || wrapper.IsNull() {
		return false
	}

	return misc.IsBytesEquals(
		wrapper.bytes,
		bytes,
		0)
}

func (wrapper *Wrapper) IsEquals(another *Wrapper) bool {
	if another == nil {
		return false
	}

	// same pointer
	if wrapper == another {
		return true
	}

	if wrapper.IsNullOrEmpty() == another.IsNullOrEmpty() {
		return true
	}

	if wrapper.BytesLength() != another.BytesLength() {
		return false
	}

	return misc.IsBytesEquals(
		wrapper.bytes,
		another.bytes,
		0)
}

// Bytes represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (wrapper *Wrapper) Bytes() *[]byte {
	return wrapper.bytes
}

// Value represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (wrapper *Wrapper) Value() *[]byte {
	return wrapper.bytes
}

// note: that it makes a copy of the content so use it wisely
func (wrapper *Wrapper) ValueWithoutPtr() []byte {
	return *wrapper.bytes
}

// note: that it makes a copy of the content so use it wisely
func (wrapper *Wrapper) String() string {
	return *wrapper.StringPtr()
}

// StringPtr is an expensive operation, it creates new memory using string(*wrapper.bytes)
// However, it does it at once, so calling it 3 times will only create once and cached result will be returned.
//
// StringPtr represents the pointer to optimize memory copying.
//
// Warning:
//  - Returns cached value from a field. Expects no modification in data.
//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
//  - Reviewer should check the mutation of the pointers.
func (wrapper *Wrapper) StringPtr() *string {
	if wrapper.content == nil && wrapper.bytes != nil {
		newString := string(*wrapper.bytes)
		wrapper.content = &newString
	}

	return wrapper.content
}

func (wrapper *Wrapper) JsonModel() *WrapperDataModel {
	return &WrapperDataModel{
		Bytes:        wrapper.bytes,
		ErrorWrapper: wrapper.errorWrapper,
		ByteType:     wrapper.byteType,
		BytesLength:  wrapper.bytesLength,
		IsWhitespace: wrapper.isWhitespace,
	}
}

func (wrapper *Wrapper) JsonModelAny() interface{} {
	return wrapper.JsonModel()
}

func (wrapper *Wrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(*wrapper.JsonModel())
}

func (wrapper *Wrapper) UnmarshalJSON(data []byte) error {
	var dataModel WrapperDataModel
	err := json.Unmarshal(data, &dataModel)

	if err == nil {
		wrapper.bytes = dataModel.Bytes
		wrapper.errorWrapper = dataModel.ErrorWrapper
		wrapper.byteType = dataModel.ByteType
		wrapper.bytesLength = dataModel.BytesLength
		wrapper.isWhitespace = dataModel.IsWhitespace
	}

	return err
}

//goland:noinspection GoLinterLocal
func (wrapper *Wrapper) Json() *corejson.Result {
	if wrapper.IsNullOrEmpty() {
		return corejson.EmptyWithoutErrorPtr()
	}

	return corejson.NewFromAny(wrapper)
}

//goland:noinspection GoLinterLocal
func (wrapper *Wrapper) ParseInjectUsingJson(
	jsonResult *corejson.Result,
) (*Wrapper, error) {
	err := jsonResult.Unmarshal(&wrapper)

	if err != nil {
		return nil, err
	}

	return wrapper, nil
}

// Panic if error
//goland:noinspection GoLinterLocal
func (wrapper *Wrapper) ParseInjectUsingJsonMust(
	jsonResult *corejson.Result,
) *Wrapper {
	newUsingJson, err :=
		wrapper.ParseInjectUsingJson(jsonResult)

	if err != nil {
		panic(err)
	}

	return newUsingJson
}

func (wrapper *Wrapper) JsonParseSelfInject(
	jsonResult *corejson.Result,
) error {
	_, err := wrapper.ParseInjectUsingJson(
		jsonResult,
	)

	return err
}

func (wrapper *Wrapper) AsJsoner() corejson.Jsoner {
	return wrapper
}

func (wrapper *Wrapper) AsJsonParseSelfInjector() corejson.JsonParseSelfInjector {
	return wrapper
}

func (wrapper *Wrapper) AsJsonMarshaller() corejson.JsonMarshaller {
	return wrapper
}
