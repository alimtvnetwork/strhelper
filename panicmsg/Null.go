package panicmsg

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Null returns "Cannot be nil or null. Reference ( " + Var(variableName, "nil") + " )"
func Null(variableName string) string {
	return SimpleValMsg(CannotBeNilMessage, variableName, strconst.NilString)
}
