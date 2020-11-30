package chars

import "unicode/utf8"

func IsUnicodeRune(char rune) bool {
	return char >= utf8.RuneSelf
}
