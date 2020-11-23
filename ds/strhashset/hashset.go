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
	hashset       *map[string]bool
	hasMapUpdated bool
	cachedList    *[]string
	length        int
	isEmptySet    bool
	sync.Mutex
}

func (hashset *Hashset) IsEmptySet() bool {
	if hashset.hasMapUpdated {
		hashset.isEmptySet = len(*hashset.hashset) == 0
	}

	return hashset.isEmptySet
}

func (hashset *Hashset) Lock() {
	hashset.Mutex.Lock()
	fmt.Println("locked")
}

func (hashset *Hashset) Unlock() {
	hashset.Mutex.Unlock()
	// TODO remove msg.
	fmt.Println("unlocked")
}

func (hashset *Hashset) Add(key string) {
	(*hashset.hashset)[key] = true
	hashset.hasMapUpdated = true
}

func (hashset *Hashset) AddWithLock(key string) {
	hashset.Lock()
	defer hashset.Unlock()

	(*hashset.hashset)[key] = true
	hashset.hasMapUpdated = true
}

func (hashset *Hashset) Has(key string) bool {
	isSet, isFound := (*hashset.hashset)[key]

	return isFound && isSet
}

func (hashset *Hashset) HasAll(keys ...string) bool {
	for _, key := range keys {
		isSet, isFound := (*hashset.hashset)[key]

		if !(isFound && isSet) {
			// not found
			return false
		}
	}

	// all found.
	return true
}

func (hashset *Hashset) HasAny(keys ...string) bool {
	for _, key := range keys {
		isSet, isFound := (*hashset.hashset)[key]

		if isFound && isSet {
			// any found
			return true
		}
	}

	// all not found.
	return false
}

func (hashset *Hashset) HasWithLock(key string) bool {
	hashset.Lock()
	defer hashset.Unlock()

	isSet, isFound := (*hashset.hashset)[key]

	return isFound && isSet
}

func (hashset *Hashset) UnsetWithLock(key string) {
	hashset.Lock()
	defer hashset.Unlock()

	(*hashset.hashset)[key] = false
	hashset.hasMapUpdated = true
}

func (hashset *Hashset) List() *[]string {
	if hashset.hasMapUpdated || hashset.cachedList == nil {
		hashset.setCached()
	}

	return hashset.cachedList
}

func (hashset *Hashset) ListWithLock() *[]string {
	hashset.Lock()
	defer hashset.Unlock()

	return hashset.List()
}

func (hashset *Hashset) setCached() {
	length := hashset.Length()
	list := make([]string, length)

	i := 0

	for key, isEnabled := range *hashset.hashset {
		if isEnabled {
			list[i] = key
		}
	}

	hashset.hasMapUpdated = false
	hashset.cachedList = &list
}

// Create a new hashset with all lower strings
func (hashset *Hashset) ToLowerSet() *Hashset {
	newMap := make(map[string]bool, hashset.Length())

	var toLower string
	for key, isEnabled := range *hashset.hashset {
		toLower = strings.ToLower(key)
		newMap[toLower] = isEnabled
	}

	return NewUsingMap(&newMap)
}

func (hashset *Hashset) Length() int {
	if hashset.hasMapUpdated {
		hashset.length = len(*hashset.hashset)
	}

	return hashset.length
}

func (hashset *Hashset) Remove(key string) {
	delete(*hashset.hashset, key)
	hashset.hasMapUpdated = true
}

func (hashset *Hashset) RemoveWithLock(key string) {
	hashset.Lock()
	defer hashset.Unlock()

	hashset.Remove(key)
}
