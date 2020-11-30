package panicmsg

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns errorStart + msg + referenceStart + Var(variableName, printVal) + spaceParenthesisEnd
// Type name NOT included
func SimpleValMsg(msg, variableName string, value interface{}) string {
	var printVal string

	if value == nil {
		printVal = strconst.NilString
	} else {
		printVal = fmt.Sprintf(strconst.SprintValueFormat, value)
	}

	return errorStart + msg + referenceStart + Var(variableName, printVal) + spaceParenthesisEnd
}
