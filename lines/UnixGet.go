package lines

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// Gets new line by \n
func UnixGet(content *string) []string {
	return strings.Split(*content, constants.NewLineUnix)
}
