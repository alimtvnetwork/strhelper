package strhelper

import (
	"gitlab.com/evatix-go/strhelper/concat"
	"gitlab.com/evatix-go/strhelper/constants"
)

// String join using Unix New Line "\n"
func GetContentFormLinesUnix(lines []string) string {
	return concat.JoinPtr(&lines, constants.NewLineUnixPtr)
}
