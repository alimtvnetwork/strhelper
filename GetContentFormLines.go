package strhelper

import (
	"gitlab.com/evatix-go/strhelper/concat"
	"gitlab.com/evatix-go/strhelper/constants"
)

// String join using Unix New Line operating system newline (For windows it is \r\n and for unix it is \n)
func GetContentFormLines(lines []string) string {
	return concat.JoinPtr(&lines, constants.NewLinePtr)
}
