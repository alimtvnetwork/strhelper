package whitespace

import "unicode"

func IsCharButWhitespace(char uint8) bool {
	return !(asciiSpaces[char] == 1 || unicode.IsSpace(rune(char)))
}
