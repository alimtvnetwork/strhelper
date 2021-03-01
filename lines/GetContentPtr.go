package lines

import (
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/concat"
)

// String join using Unix New Line operating system newline (For windows it is \r\n and for unix it is \n)
func GetContentPtr(lines *[]string) string {
	return concat.JoinPtr(lines, constants.NewLinePtr)
}
