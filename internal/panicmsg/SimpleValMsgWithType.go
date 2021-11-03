package panicmsg

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

// SimpleValMsgWithType
//
// Returns msg + referenceStart + VarWithType(typeName, variableName, printVal) + spaceParenthesisEnd
// Type name included
func SimpleValMsgWithType(msg, variableName string, value interface{}) string {
	var printVal string
	typeName := fmt.Sprintf(constants.SprintTypeFormat, value)

	if value == nil {
		printVal = constants.NilString
	} else {
		printVal = fmt.Sprintf(constants.SprintValueFormat, value)
	}

	typedVariableReference := VarWithType(typeName, variableName, printVal)

	return msg +
		referenceStart +
		typedVariableReference +
		spaceParenthesisEnd
}
