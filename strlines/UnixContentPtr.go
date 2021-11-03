package strlines

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/strconcat"
)

// UnixContentPtr
//
// String join using Unix New Line "\n"
func UnixContentPtr(lines *[]string) string {
	return strconcat.JoinPtr(lines, constants.NewLineUnix)
}
