package whitespace

func IsAsciiWhitespace(char uint8) bool {
	return asciiSpaces[char] == 1
}
