package charhelper

func CountRuneSensitive(
	str *string,
	r rune,
	at int,
) int {
	length := len(*str)
	found := 0

	if length == 0 {
		return found
	}

	runes := []rune(*str)
	length = len(runes)

	for ; at < length; at++ {
		if runes[at] == r {
			found++
		}
	}

	return found
}
