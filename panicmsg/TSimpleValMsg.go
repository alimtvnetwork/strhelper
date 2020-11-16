package panicmsg

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns msg + referenceStart + TVar(typeName, variableName, printVal) + spaceParenthesisEnd
// Type name included
func TSimpleValMsg(msg, variableName string, value interface{}) string {
	var printVal string
	typeName := fmt.Sprintf(strconst.SprintTypeFormat, value)

	if value == nil {
		printVal = strconst.NilString
	} else {
		printVal = fmt.Sprintf(strconst.SprintValueFormat, value)
	}

	typedVariableReference := TVar(typeName, variableName, printVal)

	return msg +
		referenceStart +
		typedVariableReference +
		spaceParenthesisEnd
}
