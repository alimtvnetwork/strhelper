package lines

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// Gets new line by os specific new line (For windows it is \r\n and for unix it is \n)
func Get(content *string) []string {
	return strings.Split(*content, constants.NewLine)
}

// Gets new line by os specific new line (For windows it is \r\n and for unix it is \n)
func GetPtr(content *string) *[]string {
	allLines := strings.Split(*content, constants.NewLine)

	return &allLines
}
