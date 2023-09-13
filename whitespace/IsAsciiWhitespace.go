package whitespace

import "gitlab.com/auk-go/core/constants"

func IsAsciiWhitespace(char uint8) bool {
	return asciiSpaces[char] == constants.One
}
