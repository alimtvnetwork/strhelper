package strhashset

import (
	"sync"

	"gitlab.com/evatix-go/strhelper/converters"
	"gitlab.com/evatix-go/strhelper/internal/pkg/panichelper"
)

func NewEmpty() *Hashset {
	return New(0)
}

func New(length int) *Hashset {
	hashset := make(map[string]bool, length)

	return &Hashset{
		hashset:       &hashset,
		hasMapUpdated: false,
		cachedList:    nil,
		length:        length,
		isEmptySet:    true,
		Mutex:         sync.Mutex{},
	}
}

func NewWithValues(items ...string) *Hashset {
	if items == nil {
		panichelper.NullReferenceMsg("Items cannot be null.", "items")
	}

	return NewUsingArray(&items)
}

func NewUsingStringPointersArray(inputArray *[]*string) *Hashset {
	if inputArray == nil || *inputArray == nil {
		return New(defaultItems)
	}

	maps := converters.StringsPointersArrayToMap(inputArray)

	return NewUsingMap(maps)
}

func NewUsingArray(inputArray *[]string) *Hashset {
	if inputArray == nil || *inputArray == nil {
		return New(defaultItems)
	}

	maps := converters.StringArrayToMap(inputArray)

	return NewUsingMap(maps)
}

func NewUsingMap(mapString *map[string]bool) *Hashset {
	if mapString == nil || *mapString == nil {
		return New(defaultItems)
	}

	length := len(*mapString)

	return &Hashset{
		hashset:       mapString,
		hasMapUpdated: false,
		cachedList:    nil,
		length:        length,
		isEmptySet:    length == 0,
		Mutex:         sync.Mutex{},
	}
}
