package lines

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/internal/pkg/misc"
	"gitlab.com/evatix-go/strhelper/internal/pkg/panichelper"
	"gitlab.com/evatix-go/strhelper/strconst"
	"gitlab.com/evatix-go/strhelper/strs/isstrs"
)

func Compare(
	leftLines *[]string,
	rightLines *[]string,
	startsAt int,
	isPanicOnLengthDifferent bool,
	isCaseSensitive bool,
) int {
	isLeftEmpty := isstrs.Empty(leftLines)
	isRightEmpty := isstrs.Empty(rightLines)

	if isLeftEmpty == isRightEmpty && isLeftEmpty == true {
		return 0
	}

	isLeftEmptyAndRightNot := isLeftEmpty == true && isRightEmpty == false

	if isLeftEmptyAndRightNot {
		return strconst.InvalidNotFoundCase
	}

	isRightEmptyAndLeftNot := isLeftEmpty == false && isRightEmpty == true

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
		return strconst.InvalidNotFoundCase
	}

	return resultSum
}
