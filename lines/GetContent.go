package lines

import (
	"gitlab.com/evatix-go/strhelper/concat"
	"gitlab.com/evatix-go/strhelper/strconst"
)

// String join using Unix New Line operating system newline (For windows it is \r\n and for unix it is \n)
func GetContent(lines []string) string {
	return concat.JoinPtr(&lines, strconst.NewLinePtr)
}
