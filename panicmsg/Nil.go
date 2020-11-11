package panicmsg

import "gitlab.com/evatix-go/strhelper/constants"

// Returns "Cannot be nil or null. Reference ( " + Var(variableName, "nil") + " )"
func Nil(variableName string) string {
	return SimpleValMsg(CannotBeNilMessage, variableName, constants.NilString)
}
