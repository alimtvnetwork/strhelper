package stringnum

type ascInteger struct{}


func (it *ascInteger) Positive(input string) bool {
	return it.Unsigned(input)
}

func (it *ascInteger) Negative(input string) bool {
	if input == "" || input[0] != '-' {
		return false
	}

	return it.Signed(input)
}

func (it *ascInteger) Signed(str string) bool {
	if str == "" {
		return false
	}

	firstChar := str[0]
	isFirstCharSign := firstChar == '-' || firstChar == '+'

	if isFirstCharSign {
		if len(str) <= 1 {
			return false
		}

		for _, c := range str[1:] {
			if !('0' <= c && c <= '9') {
				return false
			}
		}
	} else {
		for _, c := range str {
			if !('0' <= c && c <= '9') {
				return false
			}
		}
	}

	return true
}

// Unsigned
//
// Returns true for positive integer and `+anyAscIINumber`
func (it *ascInteger) Unsigned(str string) bool {
	if str == "" || str[0] == '-' {
		return false
	}

	firstChar := str[0]
	isFirstCharSign := firstChar == '+'

	if isFirstCharSign {
		if len(str) <= 1 {
			return false
		}

		for _, c := range str[1:] {
			if !('0' <= c && c <= '9') {
				return false
			}
		}
	} else {
		for _, c := range str {
			if !('0' <= c && c <= '9') {
				return false
			}
		}
	}

	return true
}

// UnsignedWithoutPlus
//
// Returns true for positive integer without `+` symbol as well.
func (it *ascInteger) UnsignedWithoutPlus(str string) bool {
	if str == "" {
		return false
	}

	for _, c := range str {
		if !('0' <= c && c <= '9') {
			return false
		}
	}

	return true
}
