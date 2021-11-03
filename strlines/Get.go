package strlines

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// Get
//
// - Gets new line by os specific new line (For windows it is \r\n and for unix it is \n)
func Get(content string) []string {
	return strings.Split(content, constants.NewLine)
}
