package panicmsg

import (
	"gitlab.com/evatix-go/core/constants"
)

// Null returns "Cannot be nil or null. Reference ( " + Var(variableName, "nil") + " )"
func Null(variableName string) string {
	return SimpleValMsg(CannotBeNilMessage, variableName, constants.NilString)
}
