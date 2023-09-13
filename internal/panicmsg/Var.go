package panicmsg

import (
	"gitlab.com/auk-go/core/constants"
)

// Var
//
// Returns variableName + constants.SpaceColonSpace + value
func Var(variableName, value string) string {
	return variableName + constants.SpaceColonSpace + value
}
