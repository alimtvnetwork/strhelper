package panicmsg

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
)

// SimpleValMsg
//
// Returns errorStart + msg + referenceStart + Var(variableName, printVal) + spaceParenthesisEnd
// Type name NOT included
func SimpleValMsg(msg, variableName string, value interface{}) string {
	var printVal string

	if value == nil {
		printVal = constants.NilString
	} else {
		printVal = fmt.Sprintf(constants.SprintValueFormat, value)
	}

	return errorStart + msg + referenceStart + Var(variableName, printVal) + spaceParenthesisEnd
}
