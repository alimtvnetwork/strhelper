package panicmsg

import (
	"fmt"

	"gitlab.com/evatix-go/core/constants"
)

// Returns variableName + constants.SpaceColonSpace + value
func AnyVar(variableName string, value interface{}) string {
	return variableName + constants.SpaceColonSpace + fmt.Sprintf(constants.SprintValueFormat, value)
}
