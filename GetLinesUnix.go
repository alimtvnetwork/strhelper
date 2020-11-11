package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

// Gets new line by \n
func GetLinesUnix(content string) []string {
	return strings.Split(content, constants.NewLineUnix)
}
