package strhelpercore

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coreinterface"
)

type StringResultsMap struct {
	items *map[int]*StringResult
}

func EmptyStringResultsMap() *StringResultsMap {
	return NewStringResultsMap(constants.Zero)
}

func NewStringResultsMap(capacity int) *StringResultsMap {
	currentMap := make(map[int]*StringResult, capacity)

	return &StringResultsMap{
		items: &currentMap,
	}
}

func (receiver *StringResultsMap) Length() int {
	if receiver.items == nil {
		return constants.Zero
	}

	return len(*receiver.items)
}

func (receiver *StringResultsMap) Count() int {
	return receiver.Length()
}

func (receiver *StringResultsMap) IsEmpty() bool {
	return receiver.Length() == 0
}

func (receiver *StringResultsMap) HasAnyItem() bool {
	return receiver.Length() > 0
}

// LastIndex Could be misleading, it refers to the length - 1
func (receiver *StringResultsMap) LastIndex() int {
	return receiver.Length() - 1
}

func (receiver *StringResultsMap) Add(result *StringResult) *StringResultsMap {
	if result == nil || result.FoundIndex > constants.InvalidNotFoundCase {
		return receiver
	}

	(*receiver.items)[result.FoundIndex] = result

	return receiver
}

func (receiver *StringResultsMap) HasIndex(index int) bool {
	_, has := (*receiver.items)[index]

	return has
}

func (receiver *StringResultsMap) Strings() []string {
	collection := corestr.NewCollection(receiver.Length())

	for _, result := range *receiver.items {
		collection.Add(result.String())
	}

	return collection.ListStrings()
}

func (receiver *StringResultsMap) ListStrings() []string {
	return receiver.Strings()
}

func (receiver *StringResultsMap) Items() *map[int]*StringResult {
	return receiver.items
}

func (receiver *StringResultsMap) String() string {
	return strings.Join(receiver.Strings(), constants.NewLineUnix)
}

func (receiver *StringResultsMap) AsBasicSlicer() coreinterface.BasicSlicer {
	return receiver
}
