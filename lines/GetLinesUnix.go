package lines

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/strconst"
)

// Gets new line by \n
func GetLinesUnix(content *string) []string {
	return strings.Split(*content, strconst.NewLineUnix)
}
