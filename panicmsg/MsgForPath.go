package panicmsg

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns Path : path + GetMsg(message, variableName, variableValue string)
func MsgForPath(path, message, variableName, variableValue string) string {
	return "Path : " +
		path +
		strconst.CommaSpace +
		Msg(
			message,
			variableName,
			variableValue)
}
