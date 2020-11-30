package whitespace

import "gitlab.com/evatix-go/strhelper/strconst"

const (
	maxUnit8 = 255
)

var (
	asciiSpaces = strconst.AsciiSpace
	// Reference :
	// - https://en.wikipedia.org/wiki/Newline,
	// - https://en.wikipedia.org/wiki/Whitespace_character
	// - https://en.wikipedia.org/wiki/Regular_expression#Character_classes
	asciiNewLinesChars = strconst.AsciiNewLinesChars
)
