package panicmsg

import (
	"gitlab.com/evatix-go/strhelper/constants"
)

// Returns variableName + constants.SpaceColonSpace + value
func Var(variableName, value string) string {
	return variableName + constants.SpaceColonSpace + value
}
