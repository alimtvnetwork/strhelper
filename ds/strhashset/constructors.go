package strhashset

import (
	"gitlab.com/evatix-go/core/converters"
)

func NewEmpty() *Hashset {
	return New(0)
}

func New(length int) *Hashset {
	hashset := make(map[string]bool, length)

	return &Hashset{
		hashset:    hashset,
		length:     length,
		isEmptySet: length == 0,
	}
}

func NewUsingStringPointersArray(inputArray []*string) *Hashset {
	if len(inputArray) == 0 {
		return New(defaultItems)
	}

	maps := converters.StringsPointersToStringBoolMap(&inputArray)

	return NewUsingMap(*maps)
}

func NewUsingArray(
	inputArray ...string,
) *Hashset {
	if len(inputArray) == 0 {
		return New(defaultItems)
	}

	maps := converters.
		StringsTo.
		Hashset(inputArray)

	return NewUsingMap(maps)
}

func NewUsingMap(mapString map[string]bool) *Hashset {
	if len(mapString) == 0 {
		return New(defaultItems)
	}

	length := len(mapString)

	return &Hashset{
		hashset:    mapString,
		length:     length,
		isEmptySet: length == 0,
	}
}
