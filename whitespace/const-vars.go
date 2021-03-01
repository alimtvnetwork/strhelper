package whitespace

import "gitlab.com/evatix-go/core/constants"

const (
	maxUnit8 = 255
)

var (
	asciiSpaces = constants.AsciiSpace
	// Reference :
	// - https://en.wikipedia.org/wiki/Newline,
	// - https://en.wikipedia.org/wiki/Whitespace_character
	// - https://en.wikipedia.org/wiki/Regular_expression#Character_classes
	asciiNewLinesChars = constants.AsciiNewLinesChars
)
