package strlines

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/strconcat"
)

// String join using Unix New Line "\n"
func UnixGetContent(lines []string) string {
	return strconcat.JoinPtr(&lines, constants.NewLineUnix)
}
