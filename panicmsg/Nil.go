package panicmsg

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns "Cannot be nil or null. Reference ( " + Var(variableName, "nil") + " )"
func Nil(variableName string) string {
	return SimpleValMsg(CannotBeNilMessage, variableName, strconst.NilString)
}
