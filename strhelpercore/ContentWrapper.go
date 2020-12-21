package strhelpercore

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/anyto"
	"gitlab.com/evatix-go/strhelper/content"
	"gitlab.com/evatix-go/strhelper/internal/isinternal"
	"gitlab.com/evatix-go/strhelper/internal/isstrsinternal"
	"gitlab.com/evatix-go/strhelper/internal/misc"
)

// Changing data outside makes it non stable.
//
// Kindly use pointers respectfully.
//
// Don't modify pointers values. Make copy when need to modify.
type ContentWrapper struct {
	str                   *string
	strs                  *[]string
	bytes                 *[]byte
	bytesWithError        *BytesWithError
	anyItems              *AnyItems
	valueToBytes          *[]byte
	valueToBytesWithError *BytesWithError
	valueToString         *string
	valueToStrings        *[]string
	valueToAnyItems       *AnyItems
	isEmpty               *bool
	contentType           content.Type
}

func NewContentWrapperString(str *string) *ContentWrapper {
	return &ContentWrapper{
		str:         str,
		contentType: content.String,
	}
}

func NewContentWrapperStrings(strs *[]string) *ContentWrapper {
	return &ContentWrapper{
		strs:        strs,
		contentType: content.Strings,
	}
}

func NewContentWrapperBytes(bytes *[]byte) *ContentWrapper {
	return &ContentWrapper{
		bytes:       bytes,
		contentType: content.Bytes,
	}
}

func NewContentWrapperBytesWithError(bytesWithError *BytesWithError) *ContentWrapper {
	return &ContentWrapper{
		bytesWithError: bytesWithError,
		contentType:    content.BytesWithError,
	}
}

func NewContentWrapperAnyItems(anyItems *AnyItems) *ContentWrapper {
	return &ContentWrapper{
		anyItems:    anyItems,
		contentType: content.AnyItems,
	}
}

var (
	emptyAnyItems   = AnyItems{}
	emptyBytesError = BytesWithError{}
)

func (contentWrapper *ContentWrapper) IsStringContent() bool {
	return contentWrapper.contentType == content.String
}

func (contentWrapper *ContentWrapper) IsStringsContent() bool {
	return contentWrapper.contentType == content.Strings
}

func (contentWrapper *ContentWrapper) IsBytesContent() bool {
	return contentWrapper.contentType == content.Bytes
}

func (contentWrapper *ContentWrapper) IsAnyItemsContentEmpty() bool {
	return contentWrapper.contentType == content.AnyItems && contentWrapper.IsEmpty()
}

func (contentWrapper *ContentWrapper) IsStringContentEmpty() bool {
	return contentWrapper.contentType == content.String && contentWrapper.IsEmpty()
}

func (contentWrapper *ContentWrapper) IsStringsContentEmpty() bool {
	return contentWrapper.contentType == content.Strings && contentWrapper.IsEmpty()
}

func (contentWrapper *ContentWrapper) IsBytesContentEmpty() bool {
	return contentWrapper.contentType == content.Bytes && contentWrapper.IsEmpty()
}

func (contentWrapper *ContentWrapper) IsBytesWithErrorContentEmpty() bool {
	return contentWrapper.contentType == content.Bytes && contentWrapper.IsEmpty()
}

func (contentWrapper *ContentWrapper) IsAnyItemsContent() bool {
	return contentWrapper.contentType == content.AnyItems
}

func (contentWrapper *ContentWrapper) IsBytesWithErrorContent() bool {
	return contentWrapper.contentType == content.BytesWithError
}

func (contentWrapper *ContentWrapper) IsEmpty() bool {
	if contentWrapper.isEmpty == nil {
		isEmpty := true
		switch contentWrapper.contentType {
		case content.String:
			isEmpty = contentWrapper.str == nil || *contentWrapper.str == ""
		case content.Strings:
			isEmpty = contentWrapper.strs == nil || len(*contentWrapper.strs) == 0
		case content.Bytes:
			isEmpty = contentWrapper.bytes == nil || len(*contentWrapper.bytes) == 0
		case content.AnyItems:
			isEmpty = contentWrapper.anyItems == nil || contentWrapper.anyItems.IsNullOrEmpty()
		case content.BytesWithError:
			isEmpty = contentWrapper.bytesWithError == nil || contentWrapper.bytesWithError.IsNullOrEmpty()
		default:
			panic(contentWrapper.notSupportedMessage())
		}

		contentWrapper.isEmpty = &isEmpty
	}

	return *contentWrapper.isEmpty
}

// HasAnyContent returns true if the given data is not null or nil.
func (contentWrapper *ContentWrapper) HasAnyContent() bool {
	switch contentWrapper.contentType {
	case content.String:
		return contentWrapper.str != nil
	case content.Strings:
		return contentWrapper.strs != nil
	case content.Bytes:
		return contentWrapper.bytes != nil
	case content.AnyItems:
		return contentWrapper.anyItems != nil
	case content.BytesWithError:
		return contentWrapper.bytesWithError != nil
	default:
		panic(contentWrapper.notSupportedMessage())
	}
}

func (contentWrapper *ContentWrapper) StringPtr() *string {
	return contentWrapper.str
}

func (contentWrapper *ContentWrapper) String() string {
	if contentWrapper.str == nil {
		return ""
	}

	return *contentWrapper.str
}

func (contentWrapper *ContentWrapper) Strings() []string {
	if contentWrapper.strs == nil {
		return nil
	}

	return *contentWrapper.strs
}

func (contentWrapper *ContentWrapper) StringsPtr() *[]string {
	return contentWrapper.strs
}

func (contentWrapper *ContentWrapper) Bytes() []byte {
	if contentWrapper.bytes == nil {
		return nil
	}

	return *contentWrapper.bytes
}

func (contentWrapper *ContentWrapper) BytesPtr() *[]byte {
	return contentWrapper.bytes
}

func (contentWrapper *ContentWrapper) BytesWithErrorPtr() *BytesWithError {
	return contentWrapper.bytesWithError
}

func (contentWrapper *ContentWrapper) BytesWithError() BytesWithError {
	if contentWrapper.bytesWithError == nil {
		return emptyBytesError
	}

	return *contentWrapper.bytesWithError
}

//goland:noinspection ALL
func (contentWrapper *ContentWrapper) AnyItems() AnyItems {
	if contentWrapper.anyItems == nil {
		return emptyAnyItems
	}

	return *contentWrapper.anyItems
}

func (contentWrapper *ContentWrapper) AnyItemsPtr() *AnyItems {
	if contentWrapper.anyItems == nil {
		return nil
	}

	return contentWrapper.anyItems
}

func (contentWrapper *ContentWrapper) ContentType() content.Type {
	return contentWrapper.contentType
}

func (contentWrapper *ContentWrapper) GetContent() interface{} {
	switch contentWrapper.contentType {
	case content.String:
		return contentWrapper.String()
	case content.Strings:
		return contentWrapper.Strings()
	case content.Bytes:
		return contentWrapper.Bytes()
	case content.AnyItems:
		return contentWrapper.AnyItems()
	case content.BytesWithError:
		return contentWrapper.BytesWithError()
	default:
		panic(contentWrapper.notSupportedMessage())
	}
}

func (contentWrapper *ContentWrapper) GetContentPtr() interface{} {
	switch contentWrapper.contentType {
	case content.String:
		return contentWrapper.StringPtr()
	case content.Strings:
		return contentWrapper.StringsPtr()
	case content.Bytes:
		return contentWrapper.BytesPtr()
	case content.AnyItems:
		return contentWrapper.AnyItemsPtr()
	case content.BytesWithError:
		return contentWrapper.BytesWithErrorPtr()
	default:
		panic(contentWrapper.notSupportedMessage())
	}
}

func (contentWrapper *ContentWrapper) notSupportedMessage() string {
	return fmt.Sprintf("%s %+v", "Content type not supported. CurrentType : ", contentWrapper.contentType)
}

func (contentWrapper *ContentWrapper) IsEqual(another *ContentWrapper) bool {
	if another == nil {
		return false
	}

	t1 := contentWrapper.contentType
	t2 := another.contentType

	if t1 != t2 {
		return false
	}

	if contentWrapper == another {
		return true
	}

	if contentWrapper.IsEmpty() && another.IsEmpty() {
		return true
	}

	switch contentWrapper.contentType {
	case content.String:
		return isinternal.EqualsPtr(contentWrapper.StringPtr(), another.StringPtr())
	case content.Strings:
		return isstrsinternal.Equals(contentWrapper.StringsPtr(), another.StringsPtr(), 0, true)
	case content.Bytes:
		return misc.IsBytesEquals(contentWrapper.BytesPtr(), another.BytesPtr(), 0)
	case content.AnyItems:
		return contentWrapper.AnyItemsPtr().IsEquals(contentWrapper.AnyItemsPtr())
	case content.BytesWithError:
		return contentWrapper.BytesWithErrorPtr().IsEquals(contentWrapper.BytesWithErrorPtr())
	default:
		panic(contentWrapper.notSupportedMessage())
	}
}

func (contentWrapper *ContentWrapper) extractToBytes() []byte {
	switch contentWrapper.contentType {
	case content.String:
		return []byte(contentWrapper.String())
	case content.Strings:
		return contentWrapper.stringsToBytes()
	case content.Bytes:
		return contentWrapper.Bytes()
	case content.AnyItems:
		return contentWrapper.anyItems.ToBytes()
	default:
		panic(contentWrapper.notSupportedMessage())
	}
}

func (contentWrapper *ContentWrapper) stringsToBytes() []byte {
	strs := contentWrapper.Strings()
	strsToBytes, err := anyto.Bytes(strs)

	if err != nil {
		panic(err)
	}

	if strsToBytes != nil {
		return *strsToBytes
	}

	return nil
}

func (contentWrapper *ContentWrapper) ValueToBytesPtr() *[]byte {
	if contentWrapper.valueToBytes == nil && !contentWrapper.IsEmpty() {
		toBytes := contentWrapper.extractToBytes()
		contentWrapper.valueToBytes = &toBytes
	}

	return contentWrapper.valueToBytes
}

func (contentWrapper *ContentWrapper) ValueToBytes() []byte {
	bytes := contentWrapper.ValueToBytesPtr()

	if bytes != nil {
		return *bytes
	}

	return nil
}

func (contentWrapper *ContentWrapper) ValueToStringPtr() *string {
	if contentWrapper.valueToString == nil && !contentWrapper.IsEmpty() {
		contentWrapper.valueToString = contentWrapper.extractToString()
	}

	return contentWrapper.valueToString
}

func (contentWrapper *ContentWrapper) ValueToString() string {
	strPtr := contentWrapper.ValueToStringPtr()

	if strPtr == nil {
		return ""
	}

	return *strPtr
}

func (contentWrapper *ContentWrapper) ValueToStrings() []string {
	stringsPtr := contentWrapper.ValueToStringsPtr()

	if stringsPtr == nil {
		return nil
	}

	return *stringsPtr
}

func (contentWrapper *ContentWrapper) ValueToStringsPtr() *[]string {
	if contentWrapper.valueToStrings == nil && !contentWrapper.IsEmpty() {
		contentWrapper.valueToStrings = contentWrapper.extractToStrings()
	}

	return contentWrapper.valueToStrings
}

func (contentWrapper *ContentWrapper) ValueToBytesWithError() BytesWithError {
	bytesWithError := contentWrapper.ValueToBytesWithErrorPtr()

	if bytesWithError == nil {
		return emptyBytesError
	}

	return *bytesWithError
}

func (contentWrapper *ContentWrapper) ValueToBytesWithErrorPtr() *BytesWithError {
	if contentWrapper.valueToBytesWithError == nil && !contentWrapper.IsEmpty() {
		contentWrapper.valueToBytesWithError = contentWrapper.extractToBytesWithError()
	}

	return contentWrapper.valueToBytesWithError
}

func (contentWrapper *ContentWrapper) ValueToAnyItemsPtr() *AnyItems {
	if contentWrapper.valueToAnyItems == nil && !contentWrapper.IsEmpty() {
		contentWrapper.valueToAnyItems = contentWrapper.extractToAnyItems()
	}

	return contentWrapper.valueToAnyItems
}

func (contentWrapper *ContentWrapper) extractToString() *string {
	if contentWrapper.IsStringContent() {
		toStr := contentWrapper.String()
		return &toStr
	}

	toBytes := contentWrapper.ValueToBytes()
	toStr := string(toBytes)

	return &toStr
}

func (contentWrapper *ContentWrapper) extractToStrings() *[]string {
	if contentWrapper.IsStringsContent() {
		toStr := contentWrapper.Strings()

		if toStr != nil {
			return &toStr
		}

		return nil
	}

	toBytes := contentWrapper.ValueToBytes()

	if toBytes != nil {
		toStr := string(toBytes)

		return &[]string{toStr}
	}

	return nil
}

func (contentWrapper *ContentWrapper) extractToBytesWithError() *BytesWithError {
	if contentWrapper.IsBytesWithErrorContent() {
		bytesWithError := contentWrapper.BytesWithErrorPtr()

		return bytesWithError
	}

	toBytes := contentWrapper.ValueToBytesPtr()

	return NewBytesWithNoError(toBytes)
}

func (contentWrapper *ContentWrapper) extractToAnyItems() *AnyItems {
	if contentWrapper.IsAnyItemsContent() {
		bytesWithError := contentWrapper.AnyItemsPtr()

		return bytesWithError
	}

	toBytes := contentWrapper.ValueToBytes()
	anyItems := AnyItems{}
	anyItems.Add(toBytes, true)

	return &anyItems
}
