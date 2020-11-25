package panicmsg

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns msg + referenceStart + VarWithType(typeName, variableName, printVal) + spaceParenthesisEnd
// Type name included
func SimpleValMsgWithType(msg, variableName string, value interface{}) string {
	var printVal string
	typeName := fmt.Sprintf(strconst.SprintTypeFormat, value)

	if value == nil {
		printVal = strconst.NilString
	} else {
		printVal = fmt.Sprintf(strconst.SprintValueFormat, value)
	}

	typedVariableReference := VarWithType(typeName, variableName, printVal)

	return msg +
		referenceStart +
		typedVariableReference +
		spaceParenthesisEnd
}
