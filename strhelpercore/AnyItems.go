package strhelpercore

import (
	"sync"

	"gitlab.com/auk-go/core/defaulterr"
	"gitlab.com/auk-go/strhelper/anyto"
	"gitlab.com/auk-go/strhelper/internal/isanyinternal"
)

type AnyItems struct {
	// Items represents the pointer to optimize memory copying.
	//
	// Keeping as public reasoning: For json parsing.
	//
	// Warning:
	//  - Returns cached value from a field. Expects no modification in data.
	//  - Pointer returns can be abused by modifying the bytes outside and then this object will not behave as expected.
	//  - Reviewer should check the mutation of the pointers.
	Items []interface{}
	mutex sync.Mutex
}

func NewAnyItems(items ...interface{}) *AnyItems {
	return &AnyItems{Items: items}
}

func NewAnyItemsEmpty(capacity int) *AnyItems {
	items := make([]interface{}, 0, capacity)

	return &AnyItems{Items: items}
}

func (it *AnyItems) IsNull() bool {
	return it == nil || it.Items == nil
}

func (it *AnyItems) Lock() {
	it.mutex.Lock()
}

func (it *AnyItems) Unlock() {
	it.mutex.Unlock()
}

func (it *AnyItems) Length() int {
	if it == nil {
		return 0
	}

	return len(it.Items)
}

func (it *AnyItems) LengthLock() int {
	it.Lock()
	defer it.Unlock()

	return it.Length()
}

// Add returns true upon add item.
func (it *AnyItems) Add(any interface{}, isSkipOnNil bool) bool {
	// *it.Items using this is correct as it is set on New or creation time.
	if any != nil {
		it.Items = append(it.Items, &any)

		return true
	}

	if !isSkipOnNil {
		it.Items = append(it.Items, nil)

		return true
	}

	return false
}

// AddPtr Add returns true upon add item.
func (it *AnyItems) AddPtr(anyPtr *interface{}, isSkipOnNil bool) bool {
	if anyPtr != nil && *anyPtr != nil {
		it.Items = append(it.Items, anyPtr)

		return true
	}

	if !isSkipOnNil {
		it.Items = append(it.Items, anyPtr)

		return true
	}

	return false
}

// AddPtrLock Add returns true upon add item.
func (it *AnyItems) AddPtrLock(anyPtr *interface{}, isSkipOnNil bool) bool {
	it.Lock()
	defer it.Unlock()

	return it.AddPtr(anyPtr, isSkipOnNil)
}

func (it *AnyItems) IsNullOrEmpty() bool {
	return it.Length() == 0
}

func (it *AnyItems) IsNullOrEmptyLock() bool {
	it.Lock()
	defer it.Unlock()

	return it.Length() == 0
}

func (it *AnyItems) IsEquals(another *AnyItems) bool {
	if another == nil {
		return false
	}

	if it.IsNull() == another.IsNull() {
		return true
	}

	return isanyinternal.Equals(
		it.Items,
		another.Items,
		0,
		false)
}

func (it *AnyItems) IsEqualsLock(another *AnyItems) bool {
	it.Lock()
	defer it.Unlock()

	return it.IsEquals(another)
}

// IsAnyItemsEquals returns true if both items' byte level is same.
//
// @isContinueOnBothItemParseError
//  - if true then at the same index both item has parse error then continue that means
//      assuming both are same based on error.
//  - if false then at the same index any parse error from binary then returns false no panic.
func (it *AnyItems) IsAnyItemsEquals(
	anyItemsPtr []interface{},
	isContinueOnBothItemParseError bool,
) bool {
	return isanyinternal.Equals(
		it.Items,
		anyItemsPtr,
		0,
		isContinueOnBothItemParseError,
	)
}

// IsAnyItemsEqualsLock returns true if both items' byte level is same.
//
// @isContinueOnBothItemParseError
//  - if true then at the same index both item has parse error then continue that means
//      assuming both are same based on error.
//  - if false then at the same index any parse error from binary then returns false no panic.
func (it *AnyItems) IsAnyItemsEqualsLock(
	anyItemsPtr []interface{},
	isContinueOnBothItemParseError bool,
) bool {
	it.Lock()
	defer it.Unlock()

	return isanyinternal.Equals(
		it.Items,
		anyItemsPtr,
		0,
		isContinueOnBothItemParseError)
}

// ToBytesWithError creates BytesWithError pointer using *AnyItems.Items (no panic if nil)
func (it *AnyItems) ToBytesWithError() *BytesWithError {
	if it.IsNull() {
		return NewBytesWithErrorUsingAny(nil)
	}

	return NewBytesWithErrorUsingAny(it.Items)
}

// ToBytes creates []byte pointer using *AnyItems.Items
func (it *AnyItems) ToBytes() ([]byte, error) {
	if it.IsNullOrEmpty() {
		return []byte{}, defaulterr.CannotProcessNilOrEmpty
	}

	rawBytes, err := anyto.Bytes(it.Items)

	return rawBytes, err
}

// ToBytesWithErrorLock ToBytesWithError creates BytesWithError pointer using *AnyItems.Items
func (it *AnyItems) ToBytesWithErrorLock() *BytesWithError {
	it.Lock()
	defer it.Unlock()

	return it.ToBytesWithError()
}
