package panicmsg

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns "Error : message reference ( variableName constants.SpaceColonSpace variableValue )"
func Msg(message, variableName, variableValue string) string {
	return message +
		referenceStart +
		variableName +
		strconst.SpaceColonSpace +
		variableValue +
		spaceParenthesisEnd
}
