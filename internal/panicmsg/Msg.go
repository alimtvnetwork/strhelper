package panicmsg

import (
	"gitlab.com/evatix-go/core/constants"
)

// Msg
//
// Returns "Error : message reference ( variableName constants.SpaceColonSpace variableValue )"
func Msg(message, variableName, variableValue string) string {
	return message +
		referenceStart +
		variableName +
		constants.SpaceColonSpace +
		variableValue +
		spaceParenthesisEnd
}
