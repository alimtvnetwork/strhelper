package panicmsg

import (
	"gitlab.com/evatix-go/strhelper/constants"
)

// Returns "Error : message reference ( variableName constants.SpaceColonSpace variableValue )"
func Msg(message, variableName, variableValue string) string {
	return errorStart +
		message +
		referenceStart +
		variableName +
		constants.SpaceColonSpace +
		variableValue +
		spaceParenthesisEnd
}
