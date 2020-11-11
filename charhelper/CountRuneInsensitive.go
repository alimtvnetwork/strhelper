package charhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/runehelper"
)

func CountRuneInsensitive(
	str *string,
	r rune,
	at int,
) int {
	length := len(*str)
	found := 0

	if length == 0 {
		return found
	}

	runes := []rune(strings.ToLower(*str))
	length = len(runes) // it needs to be updated because previous length was for byte/asc it changes.
	r = runehelper.ToLowerRune(r)

	for ; at < length; at++ {
		if runes[at] == r {
			found++
		}
	}

	return found
}
