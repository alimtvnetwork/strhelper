package strhelpercore

import (
	"regexp"
	"sync"
)

type RegExResultWrapper struct {
	Index int
	*regexp.Regexp
	content      *string
	lines        []string
	indexes      [][]int
	mutexLines   sync.Mutex
	mutexIndexes sync.Mutex
}

// Cached version, once populated it will not take any more resources to serve again.
func (regExResultWrapper *RegExResultWrapper) Lines() []string {
	regExResultWrapper.mutexLines.Lock()
	defer regExResultWrapper.mutexLines.Unlock()

	if isEmptyStringArrayPtr(&regExResultWrapper.lines) {
		regExResultWrapper.lines = regExResultWrapper.Regexp.FindAllString(*regExResultWrapper.content, -1)
	}

	return regExResultWrapper.lines
}

// Cached version, once populated it will not take any more resources to serve again.
func (regExResultWrapper *RegExResultWrapper) Indexes() [][]int {
	regExResultWrapper.mutexIndexes.Lock()
	defer regExResultWrapper.mutexIndexes.Unlock()

	if isEmptyIntArrayOfArrayPtr(&regExResultWrapper.indexes) {
		regExResultWrapper.indexes = regExResultWrapper.Regexp.FindAllIndex([]byte(*regExResultWrapper.content), -1)
	}

	return regExResultWrapper.indexes
}

func NewRegExResultWrapper(
	index int,
	content *string,
	regexp *regexp.Regexp,
) *RegExResultWrapper {
	return &RegExResultWrapper{
		Index:        index,
		Regexp:       regexp,
		content:      content,
		lines:        nil,
		indexes:      nil,
		mutexLines:   sync.Mutex{},
		mutexIndexes: sync.Mutex{},
	}
}
