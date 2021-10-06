package strhashset

import (
	"fmt"
	"strings"
	"sync"
)

const (
	defaultItems = 10
)

type Hashset struct {
	hashset       map[string]bool
	hasMapUpdated bool
	cachedList    []string
	length        int
	isEmptySet    bool
	sync.Mutex
}

func (it *Hashset) IsEmptySet() bool {
	if it.hasMapUpdated {
		it.isEmptySet = len(it.hashset) == 0
	}

	return it.isEmptySet
}

func (it *Hashset) Lock() {
	it.Mutex.Lock()
	fmt.Println("locked")
}

func (it *Hashset) Unlock() {
	it.Mutex.Unlock()
	// TODO remove msg.
	fmt.Println("unlocked")
}

func (it *Hashset) Add(key string) {
	it.hashset[key] = true
	it.hasMapUpdated = true
}

func (it *Hashset) AddWithLock(key string) {
	it.Lock()
	defer it.Unlock()

	it.hashset[key] = true
	it.hasMapUpdated = true
}

func (it *Hashset) Has(key string) bool {
	isSet, isFound := it.hashset[key]

	return isFound && isSet
}

func (it *Hashset) HasAll(keys ...string) bool {
	for _, key := range keys {
		isSet, isFound := it.hashset[key]

		if !(isFound && isSet) {
			// not found
			return false
		}
	}

	// all found.
	return true
}

func (it *Hashset) HasAny(keys ...string) bool {
	for _, key := range keys {
		isSet, isFound := it.hashset[key]

		if isFound && isSet {
			// any found
			return true
		}
	}

	// all not found.
	return false
}

func (it *Hashset) HasWithLock(key string) bool {
	it.Lock()
	defer it.Unlock()

	isSet, isFound := it.hashset[key]

	return isFound && isSet
}

func (it *Hashset) UnsetWithLock(key string) {
	it.Lock()
	defer it.Unlock()

	it.hashset[key] = false
	it.hasMapUpdated = true
}

func (it *Hashset) List() []string {
	if it.hasMapUpdated || it.cachedList == nil {
		it.setCached()
	}

	return it.cachedList
}

func (it *Hashset) ListWithLock() []string {
	it.Lock()
	defer it.Unlock()

	return it.List()
}

func (it *Hashset) setCached() {
	length := it.Length()
	list := make([]string, length)

	i := 0

	for key, isEnabled := range it.hashset {
		if isEnabled {
			list[i] = key
		}
	}

	it.hasMapUpdated = false
	it.cachedList = list
}

// Create a new hashset with all lower strings
func (it *Hashset) ToLowerSet() *Hashset {
	newMap := make(map[string]bool, it.Length())

	var toLower string
	for key, isEnabled := range it.hashset {
		toLower = strings.ToLower(key)
		newMap[toLower] = isEnabled
	}

	return NewUsingMap(newMap)
}

func (it *Hashset) Length() int {
	if it.hasMapUpdated {
		it.length = len(it.hashset)
	}

	return it.length
}

func (it *Hashset) Remove(key string) {
	delete(it.hashset, key)
	it.hasMapUpdated = true
}

func (it *Hashset) RemoveWithLock(key string) {
	it.Lock()
	defer it.Unlock()

	it.Remove(key)
}
