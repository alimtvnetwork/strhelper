package strlines

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// GetPtr
//
// Gets new line by os specific new line
// (For windows it is \r\n and for unix it is \n)
func GetPtr(content *string) []string {
	allLines := strings.Split(*content, constants.NewLine)

	return allLines
}
