package isstrs

import (
	"sort"
	"strings"

	"gitlab.com/auk-go/core/corecmp"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coredata/stringslice"
)

func Equal(leftLines, rightLines []string) bool {
	return corecmp.IsStringsEqual(leftLines, rightLines)
}

func EqualPtr(leftLines, rightLines *[]string) bool {
	return corecmp.IsStringsEqualPtr(leftLines, rightLines)
}

// DistinctEqual
//
//  both created to map and then compare from there
func DistinctEqual(leftLines, rightLines []string) bool {
	leftMap := corestr.New.Hashset.Strings(leftLines)
	rightMap := corestr.New.Hashset.Strings(rightLines)

	return leftMap.IsEqualsPtr(rightMap)
}

// UnorderedEqual
//
//  sort both and then compare, if length are not equal then mismatch
//
// Don't mutate, create copy and then sort and then returns the final result.
func UnorderedEqual(
	leftLines, rightLines []string,
) bool {
	if len(leftLines) != len(rightLines) {
		return false
	}

	leftSored := stringslice.Clone(leftLines)
	sort.Strings(leftSored)

	rightSort := stringslice.Clone(rightLines)
	sort.Strings(rightSort)

	return Equal(leftSored, rightSort)
}

// EqualByFunc
//
//  checks comparison by the given function.
func EqualByFunc(
	leftLines, rightLines []string,
	isMatchCheckerFunc func(index int, left, right string) (isMatch bool),
) bool {
	if len(leftLines) != len(rightLines) {
		return false
	}

	for i, leftLine := range leftLines {
		rightLine := rightLines[i]

		if !isMatchCheckerFunc(i, leftLine, rightLine) {
			return false
		}
	}

	return true
}

// EqualByFuncLinesSplit
//
//  first splits the line and then takes
//  lines and process same as EqualByFunc
func EqualByFuncLinesSplit(
	isTrim bool,
	splitter string,
	leftLine, rightLine string,
	isMatchCheckerFunc func(index int, left, right string) (isMatch bool),
) bool {
	leftLines := strings.Split(leftLine, splitter)
	rightLines := strings.Split(rightLine, splitter)

	for i, curLeftLine := range leftLines {
		curRightLine := rightLines[i]

		if isTrim {
			curLeftLine = strings.TrimSpace(curLeftLine)
			curRightLine = strings.TrimSpace(curRightLine)
		}

		if !isMatchCheckerFunc(i, curLeftLine, curRightLine) {
			return false
		}
	}

	return true
}
