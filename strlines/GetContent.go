package strlines

import (
	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/strconcat"
)

// GetContent
//
// String join using Unix New Line operating system newline
// (For windows it is \r\n and for unix it is \n)
func GetContent(lines []string) string {
	return strconcat.JoinPtr(&lines, constants.NewLine)
}
