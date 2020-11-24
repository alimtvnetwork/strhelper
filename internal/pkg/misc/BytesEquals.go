package misc

// Equals compares leftLines and rightLines and returns bool
//  - If both nil returns true.
//  - If one nil and another is not then returns false.
//  - If both lengths are not same returns false.
//  - If both pointers are same returns true.
//  - If all the lines are equals as per the case sensitivity then returns true or else false.
func BytesEquals(
	leftBytes *[]byte,
	rightBytes *[]byte,
	startsAt int,
) bool {
	isLeftEmpty := EmptyBytes(leftBytes)
	isRightEmpty := EmptyBytes(rightBytes)

	if isLeftEmpty == isRightEmpty && isLeftEmpty == true {
		return true
	}

	isLeftEmptyAndRightNot := isLeftEmpty == true && isRightEmpty == false

	if isLeftEmptyAndRightNot {
		return false
	}

	isRightEmptyAndLeftNot := isLeftEmpty == false && isRightEmpty == true

	if isRightEmptyAndLeftNot {
		return false
	}

	// if both pointers are same
	if leftBytes == rightBytes {
		return true
	}

	leftLength := len(*leftBytes)
	rightLength := len(*rightBytes)

	if leftLength != rightLength {
		return false
	}

	if startsAt < 0 || startsAt > leftLength-1 {
		panic("Start index has exceed length or negative.")
	}

	return bytesEqual(
		leftBytes,
		rightBytes,
		startsAt,
	)
}

func bytesEqual(
	leftBytes *[]byte,
	rightBytes *[]byte,
	startsAt int,
) bool {
	leftLength := len(*leftBytes)

	for ; startsAt < leftLength; startsAt++ {
		left := (*leftBytes)[startsAt]
		right := (*rightBytes)[startsAt]

		if left != right {
			return false
		}
	}

	return true
}
