package strlines

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/misc"
	"gitlab.com/evatix-go/strhelper/internal/panichelper"
	"gitlab.com/evatix-go/strhelper/strs/isstrs"
)

// Compare
//
// returns similar to strings.Compare(...), all cumulated sum values of strings.Compare
// Here it returns for multiple lines.
//
//  Having 0 doesn't confirm lines are equal. It can be only useful for sorting only ([][]str).
//
// Expression / Logic:
//  - Takes each item in lines and compare using strings.Compare() and stores += strings.Compare()
//  - Inside strings.Compare
//      - left(a) == right(b) return 0
//      - left(a) < right(b) return -1
//      - left(a) > right(b) return +1
//  - Notes: if there are 2 lines (left, right passed), one is -1 and in another is +1
//           the ultimate compare value will be 0 that doesn't mean both lines are same.
//
// Returns:
//  - leftLines nil or empty or rightLines has fewer lines than rightLines returns -1
//  - rightLines nil or empty or rightLines has fewer lines than leftLines returns +1
//  - if both lines nil or empty returns 0
//  - Cumulated Values == 0 return 0, doesn't confirm lines are equal.
//  - Cumulated Values > 0  return +1
//  - Cumulated Values < 0  return -1
func Compare(
	leftLines *[]string,
	rightLines *[]string,
	startsAt int,
	isPanicOnLengthDifferent bool,
	isCaseSensitive bool,
) int {
	isLeftEmpty := isstrs.EmptyPtr(leftLines)
	isRightEmpty := isstrs.EmptyPtr(rightLines)

	if isRightEmpty == isLeftEmpty && isLeftEmpty == true {
		return 0
	}

	isLeftEmptyAndRightNot := isLeftEmpty && isRightEmpty == false

	if isLeftEmptyAndRightNot {
		return constants.InvalidNotFoundCase
	}

	isRightEmptyAndLeftNot := !isLeftEmpty && isRightEmpty == true

	if isRightEmptyAndLeftNot {
		return 1
	}

	leftLength := len(*leftLines)
	rightLength := len(*rightLines)
	minLength := misc.MinInt(leftLength, rightLength)

	if startsAt < 0 || startsAt > minLength-1 {
		panichelper.StartAtIndexFailed(startsAt, minLength)
	}

	isPanicSatisfied := isPanicOnLengthDifferent && leftLength != rightLength
	panichelper.SimplePanic(
		isPanicSatisfied,
		"isPanicOnLengthDifferent : left and right lines lengths are not equal.")

	if isCaseSensitive {
		return caseSensitiveCompare(
			leftLines,
			rightLines,
			startsAt,
		)
	}

	return caseInsensitiveCompare(
		leftLines,
		rightLines,
		startsAt,
	)
}

func caseSensitiveCompare(
	leftLines *[]string,
	rightLines *[]string,
	startsAt int,
) int {
	leftLength := len(*leftLines)
	rightLength := len(*rightLines)
	minLength := misc.MinInt(leftLength, rightLength)
	resultSum := 0

	for ; startsAt < minLength; startsAt++ {
		left := (*leftLines)[startsAt]
		right := (*rightLines)[startsAt]
		resultSum += strings.Compare(left, right)
	}

	return simplifiedFinalCompareResult(
		resultSum,
		leftLength,
		rightLength)
}

func caseInsensitiveCompare(
	leftLines *[]string,
	rightLines *[]string,
	startsAt int,
) int {
	leftLength := len(*leftLines)
	rightLength := len(*rightLines)
	minLength := misc.MinInt(leftLength, rightLength)

	resultSum := 0

	for ; startsAt < minLength; startsAt++ {
		left := strings.ToLower((*leftLines)[startsAt])
		right := strings.ToLower((*rightLines)[startsAt])
		resultSum += strings.Compare(left, right)
	}

	return simplifiedFinalCompareResult(
		resultSum,
		leftLength,
		rightLength)
}

func simplifiedFinalCompareResult(
	resultSum,
	leftLength,
	rightLength int,
) int {
	minLength := misc.MinInt(leftLength, rightLength)
	maxLength := misc.MaxInt(leftLength, rightLength)
	diff := maxLength - minLength
	if diff > 0 && rightLength > leftLength {
		resultSum += -1 * diff
	} else if diff > 0 {
		resultSum += 1 * diff
	}

	if resultSum > 0 {
		return 1
	}

	if resultSum < 0 {
		return constants.InvalidNotFoundCase
	}

	return resultSum
}
