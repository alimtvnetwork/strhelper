package strhelpercore

import (
	"regexp"

	"gitlab.com/auk-go/core/constants"
)

type RegExResultWrapper struct {
	Index int
	*regexp.Regexp
	content string
	lines   []string
	indexes [][]int
}

// Lines
//
// Cached version, once populated it will not take any more resources to serve again.
func (regExResultWrapper *RegExResultWrapper) Lines() []string {
	if isEmptyStringArrayPtr(&regExResultWrapper.lines) {
		regExResultWrapper.lines = regExResultWrapper.
			Regexp.
			FindAllString(
				regExResultWrapper.content,
				constants.TakeAllMinusOne)
	}

	return regExResultWrapper.lines
}

// Indexes
//
// Cached version, once populated it will not take any more resources to serve again.
func (regExResultWrapper *RegExResultWrapper) Indexes() [][]int {
	if isEmptyIntArrayOfArrayPtr(&regExResultWrapper.indexes) {
		allBytes := []byte(regExResultWrapper.content)

		regExResultWrapper.indexes = regExResultWrapper.
			Regexp.
			FindAllIndex(allBytes, constants.TakeAllMinusOne)
	}

	return regExResultWrapper.indexes
}

func NewRegExResultWrapper(
	index int,
	content string,
	regexp *regexp.Regexp,
) *RegExResultWrapper {
	return &RegExResultWrapper{
		Index:   index,
		Regexp:  regexp,
		content: content,
	}
}
