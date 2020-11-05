package panicmsg

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/constants"
)

// Returns msg + referenceStart + TVar(typeName, variableName, printVal) + spaceParenthesisEnd
// Type name included
func TSimpleValMsg(msg, variableName string, value interface{}) string {
	var printVal string
	typeName := fmt.Sprintf(constants.SprintTypeFormat, value)

	if value == nil {
		printVal = constants.NilString
	} else {
		printVal = fmt.Sprintf(constants.SprintValueFormat, value)
	}

	typedVariableReference := TVar(typeName, variableName, printVal)

	return errorStart +
		msg +
		referenceStart +
		typedVariableReference +
		spaceParenthesisEnd
}
