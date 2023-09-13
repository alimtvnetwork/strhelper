package panicmsg

import (
	"gitlab.com/auk-go/core/constants"
)

// Returns Path : path + GetMsg(message, variableName, variableValue string)
func MsgForPath(path, message, variableName, variableValue string) string {
	return "Path : " +
		path +
		constants.CommaSpace +
		Msg(
			message,
			variableName,
			variableValue)
}
