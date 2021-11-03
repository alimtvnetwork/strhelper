package stringnum

// ascNumber
//
// represents verification of ASCII to float numbers
type ascNumber struct{}

func (it *ascNumber) Positive(input string) bool {
	return it.Unsigned(input)
}

func (it *ascNumber) Negative(input string) bool {
	if input == "" || input[0] != '-' {
		return false
	}

	return it.Signed(input)
}

// Signed
//
// Example : https://play.golang.org/p/_314bvo6TDk
func (it *ascNumber) Signed(str string) bool {
	if str == "" {
		return false
	}

	firstChar := str[0]
	isFirstCharSign := firstChar == '-' || firstChar == '+'
	isSingleDotFound := false

	if isFirstCharSign {
		if len(str) <= 1 {
			return false
		}

		for _, c := range str[1:] {
			if !isSingleDotFound && c == '.' {
				isSingleDotFound = true
				continue
			}

			if !(('0' <= c && c <= '9') || (!isSingleDotFound && c == '.')) {
				return false
			}
		}
	} else {
		for _, c := range str {
			if !isSingleDotFound && c == '.' {
				isSingleDotFound = true
				continue
			}

			if !(('0' <= c && c <= '9') || (!isSingleDotFound && c == '.')) {
				return false
			}
		}
	}

	return true
}

// Unsigned
//
// Example : https://play.golang.org/p/NohlCrIj77r
func (it *ascNumber) Unsigned(str string) bool {
	if str == "" || str[0] == '-' {
		return false
	}

	processingPointer := str
	firstChar := str[0]
	isPositiveSign := firstChar == '+'

	if isPositiveSign && len(str) <= 1 {
		return false
	}

	if isPositiveSign {
		// Reference : https://blog.golang.org/slices-intro | https://i.imgur.com/O3Hlmac.png
		// no copy just points
		newPointers := str[1:]
		processingPointer = newPointers
	}

	isSingleDotFound := false

	for _, c := range processingPointer {
		if !isSingleDotFound && c == '.' {
			isSingleDotFound = true
			continue
		}

		if !(('0' <= c && c <= '9') || (!isSingleDotFound && c == '.')) {
			return false
		}
	}

	return true
}

// UnsignedWithoutPlus
//
// Returns true for positive number without `+` symbol as well.
func (it *ascNumber) UnsignedWithoutPlus(str string) bool {
	if str == "" || str[0] == '-' || str[0] == '+' {
		return false
	}

	isSingleDotFound := false

	for _, c := range str {
		if !isSingleDotFound && c == '.' {
			isSingleDotFound = true
			continue
		}

		if !(('0' <= c && c <= '9') || (!isSingleDotFound && c == '.')) {
			return false
		}
	}

	return true
}
