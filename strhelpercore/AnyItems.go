package strhelpercore

import (
	"sync"

	"gitlab.com/evatix-go/strhelper/internal/pkg/isanyinternal"
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
	Items *[]*interface{}
	mutex sync.Mutex
}

func NewAnyItems(items *[]*interface{}) *AnyItems {
	return &AnyItems{Items: items}
}

func NewAnyItemsEmpty(capacity int) *AnyItems {
	items := make([]*interface{}, 0, capacity)

	return &AnyItems{Items: &items}
}

func (anyItems *AnyItems) IsNull() bool {
	return anyItems.Items == nil || *anyItems.Items == nil
}

func (anyItems *AnyItems) Lock() {
	anyItems.mutex.Lock()
}

func (anyItems *AnyItems) Unlock() {
	anyItems.mutex.Unlock()
}

func (anyItems *AnyItems) Length() int {
	return len(*anyItems.Items)
}

func (anyItems *AnyItems) LengthLock() int {
	anyItems.Lock()
	defer anyItems.Unlock()

	return len(*anyItems.Items)
}

// Add returns true upon add item.
func (anyItems *AnyItems) Add(any interface{}, isSkipOnNil bool) bool {
	if any != nil {
		*anyItems.Items = append(*anyItems.Items, &any)

		return true
	}

	if !isSkipOnNil {
		*anyItems.Items = append(*anyItems.Items, nil)

		return true
	}

	return false
}

// Add returns true upon add item.
func (anyItems *AnyItems) AddPtr(anyPtr *interface{}, isSkipOnNil bool) bool {
	if anyPtr != nil && *anyPtr != nil {
		*anyItems.Items = append(*anyItems.Items, anyPtr)

		return true
	}

	if !isSkipOnNil {
		*anyItems.Items = append(*anyItems.Items, anyPtr)

		return true
	}

	return false
}

// Add returns true upon add item.
func (anyItems *AnyItems) AddPtrLock(anyPtr *interface{}, isSkipOnNil bool) bool {
	anyItems.Lock()
	defer anyItems.Unlock()

	return anyItems.AddPtr(anyPtr, isSkipOnNil)
}

func (anyItems *AnyItems) IsNullOrEmpty() bool {
	return anyItems.Items == nil || len(*anyItems.Items) == 0
}

func (anyItems *AnyItems) IsNullOrEmptyLock() bool {
	anyItems.Lock()
	defer anyItems.Unlock()

	return anyItems.Items == nil || len(*anyItems.Items) == 0
}

func (anyItems *AnyItems) IsEquals(another *AnyItems) bool {
	if another == nil {
		return false
	}

	if anyItems.IsNull() == another.IsNull() {
		return true
	}

	return isanyinternal.PointersOfPointersAnyItemsEquals(
		anyItems.Items,
		another.Items,
		0,
		false)
}

func (anyItems *AnyItems) IsEqualsLock(another *AnyItems) bool {
	anyItems.Lock()
	defer anyItems.Unlock()

	return anyItems.IsEquals(another)
}

// IsAnyItemsEquals returns true if both items byte level is same.
//
// @isContinueOnBothItemParseError
//  - if true then if at the same index both item has parse error then continue that means
//      assuming both are same based on error.
//  - if false then if at the same index any parse error from binary then returns false no panic.
func (anyItems *AnyItems) IsAnyItemsEquals(anyItemsPtr *[]*interface{}, isContinueOnBothItemParseError bool) bool {
	return isanyinternal.PointersOfPointersAnyItemsEquals(
		anyItems.Items,
		anyItemsPtr,
		0,
		isContinueOnBothItemParseError,
	)
}

// IsAnyItemsEqualsLock returns true if both items byte level is same.
//
// @isContinueOnBothItemParseError
//  - if true then if at the same index both item has parse error then continue that means
//      assuming both are same based on error.
//  - if false then if at the same index any parse error from binary then returns false no panic.
func (anyItems *AnyItems) IsAnyItemsEqualsLock(anyItemsPtr *[]*interface{}, isContinueOnBothItemParseError bool) bool {
	anyItems.Lock()
	defer anyItems.Unlock()

	return isanyinternal.PointersOfPointersAnyItemsEquals(
		anyItems.Items,
		anyItemsPtr,
		0,
		isContinueOnBothItemParseError)
}

// IsAnyItemsWithoutPointersEquals returns true if both items byte level is same.
//
// @isContinueOnBothItemParseError
//  - if true then if at the same index both item has parse error then continue that means
//      assuming both are same based on error.
//  - if false then if at the same index any parse error from binary then returns false no panic.
func (anyItems *AnyItems) IsAnyItemsWithoutPointersEquals(
	anyItemsValues *[]interface{},
	isContinueOnBothItemParseError bool,
) bool {
	return isanyinternal.ItemsEqualsWhereOnePointersOfPointersAnyItems(
		anyItems.Items,
		anyItemsValues,
		0,
		isContinueOnBothItemParseError)
}

// IsAnyItemsWithoutPointersEqualsLock returns true if both items byte level is same.
//
// @isContinueOnBothItemParseError
//  - if true then if at the same index both item has parse error then continue that means
//      assuming both are same based on error.
//  - if false then if at the same index any parse error from binary then returns false no panic.
func (anyItems *AnyItems) IsAnyItemsWithoutPointersEqualsLock(
	anyItemsValues *[]interface{},
	isContinueOnBothItemParseError bool,
) bool {
	anyItems.Lock()
	defer anyItems.Unlock()

	return isanyinternal.ItemsEqualsWhereOnePointersOfPointersAnyItems(
		anyItems.Items,
		anyItemsValues,
		0,
		isContinueOnBothItemParseError)
}

// ToBytesWithError creates BytesWithError pointer using *AnyItems.Items
func (anyItems *AnyItems) ToBytesWithError() *BytesWithError {
	return NewBytesWithErrorUsingAny(*anyItems.Items)
}

// ToBytesWithError creates BytesWithError pointer using *AnyItems.Items
func (anyItems *AnyItems) ToBytesWithErrorLock() *BytesWithError {
	anyItems.Lock()
	defer anyItems.Unlock()

	return NewBytesWithErrorUsingAny(*anyItems.Items)
}
