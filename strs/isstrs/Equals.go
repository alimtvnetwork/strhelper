package isstrs

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/internal/panichelper"
)

// Equals compares leftLines and rightLines and returns bool
//  - If both nil returns true.
//  - If one nil and another is not then returns false.
//  - If both lengths are not same returns false.
//  - If both pointers are same returns true.
//  - If all the lines are equals as per the case sensitivity then returns true or else false.
func Equals(
	leftLines []string,
	rightLines []string,
	startsAt int,
	isCaseSensitive bool,
) bool {
	isLeftEmpty := Empty(leftLines)
	isRightEmpty := Empty(rightLines)

	if isLeftEmpty == isRightEmpty && true == isLeftEmpty {
		return true
	}

	isLeftEmptyAndRightNot := isLeftEmpty && isRightEmpty == false

	if isLeftEmptyAndRightNot {
		return false
	}

	isRightEmptyAndLeftNot := !isLeftEmpty && isRightEmpty == true

	if isRightEmptyAndLeftNot {
		return false
	}

	leftLength := len(leftLines)
	rightLength := len(rightLines)

	if leftLength != rightLength {
		return false
	}

	if startsAt < 0 || startsAt > leftLength-1 {
		panichelper.StartAtIndexFailed(startsAt, leftLength)
	}

	if isCaseSensitive {
		return caseSensitiveEqual(
			leftLines,
			rightLines,
			startsAt,
		)
	}

	return caseInsensitiveEqual(
		leftLines,
		rightLines,
		startsAt,
	)
}

func caseSensitiveEqual(
	leftLines []string,
	rightLines []string,
	startsAt int,
) bool {
	leftLength := len(leftLines)

	for ; startsAt < leftLength; startsAt++ {
		left := leftLines[startsAt]
		right := rightLines[startsAt]

		if left != right {
			return false
		}
	}

	return true
}

func caseInsensitiveEqual(
	leftLines []string,
	rightLines []string,
	startsAt int,
) bool {
	leftLength := len(leftLines)

	for ; startsAt < leftLength; startsAt++ {
		left := strings.ToLower(leftLines[startsAt])
		right := strings.ToLower(rightLines[startsAt])

		if left != right {
			return false
		}
	}

	return true
}
