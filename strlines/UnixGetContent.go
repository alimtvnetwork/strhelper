package strlines

import (
	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/strconcat"
)

// UnixGetContent
//
// String join using Unix New Line "\n"
func UnixGetContent(lines ...string) string {
	return strconcat.JoinPtr(&lines, constants.NewLineUnix)
}
