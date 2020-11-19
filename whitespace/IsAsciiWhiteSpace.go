package whitespace

func IsAsciiWhiteSpace(char uint8) bool {
	return asciiSpaces[char] == 1
}
