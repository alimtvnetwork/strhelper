package lines

import (
	"gitlab.com/evatix-go/strhelper/concat"
	"gitlab.com/evatix-go/strhelper/strconst"
)

// String join using Unix New Line "\n"
func UnixGetContent(lines []string) string {
	return concat.JoinPtr(&lines, strconst.NewLineUnixPtr)
}
