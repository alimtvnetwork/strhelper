package charhelper

// Make str to rune array
// if str == nil or length == 0 then returns nil
func ToRunesPtr(str *string) *[]rune {
	length := len(*str)

	if length == 0 {
		return nil
	}

	newChars := make([]rune, length)

	for i := 0; i < length; i++ {
		newChars[i] = rune((*str)[i])
	}

	return &newChars
}
