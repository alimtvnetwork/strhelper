package strlines

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// UnixGet
//
// Gets new line by \n
func UnixGet(content string) []string {
	return strings.Split(
		content, constants.NewLineUnix)
}
