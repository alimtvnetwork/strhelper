package lines

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/strconst"
)

// Gets new line by os specific new line (For windows it is \r\n and for unix it is \n)
func GetLines(content *string) []string {
	return strings.Split(*content, strconst.NewLine)
}
