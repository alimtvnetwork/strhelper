package lines

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/concat"
)

// String join using Unix New Line "\n"
func UnixContentPtr(lines *[]string) string {
	return concat.JoinPtr(lines, constants.NewLineUnixPtr)
}
