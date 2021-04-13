package whitespace

import "gitlab.com/evatix-go/core/constants"

func IsAsciiWhitespace(char uint8) bool {
	return asciiSpaces[char] == constants.One
}
