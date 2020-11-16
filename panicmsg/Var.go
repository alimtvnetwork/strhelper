package panicmsg

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns variableName + constants.SpaceColonSpace + value
func Var(variableName, value string) string {
	return variableName + strconst.SpaceColonSpace + value
}
