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

func (it *StringResultsMap) Length() int {
	if it.items == nil {
		return constants.Zero
	}

	return len(*it.items)
}

func (it *StringResultsMap) Count() int {
	return it.Length()
}

func (it *StringResultsMap) IsEmpty() bool {
	return it.Length() == 0
}

func (it *StringResultsMap) HasAnyItem() bool {
	return it.Length() > 0
}

// LastIndex Could be misleading, it refers to the length - 1
func (it *StringResultsMap) LastIndex() int {
	return it.Length() - 1
}

func (it *StringResultsMap) AddFoundOnly(result *StringResult) *StringResultsMap {
	if result == nil || result.FoundIndex > constants.InvalidNotFoundCase {
		return it
	}

	(*it.items)[result.FoundIndex] = result

	return it
}

func (it *StringResultsMap) HasIndex(index int) bool {
	_, has := (*it.items)[index]

	return has
}

// Strings found strings
func (it *StringResultsMap) Strings() []string {
	collection := corestr.New.Collection.Cap(it.Length())

	for _, result := range *it.items {
		collection.Add(result.Line)
	}

	return collection.ListStrings()
}

func (it *StringResultsMap) ListStrings() []string {
	return it.Strings()
}

func (it *StringResultsMap) ListStringResults() []StringResult {
	slice := make([]StringResult, it.Length())

	for _, result := range *it.items {
		slice = append(slice, *result)
	}

	return slice
}

func (it *StringResultsMap) Items() *map[int]*StringResult {
	return it.items
}

func (it *StringResultsMap) String() string {
	return strings.Join(it.Strings(), constants.NewLineUnix)
}

func (it *StringResultsMap) AsBasicSlicer() coreinterface.BasicSlicer {
	return it
}
