package panicmsg

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns variableName + constants.SpaceColonSpace + value
func AnyVar(variableName string, value interface{}) string {
	return variableName + strconst.SpaceColonSpace + fmt.Sprintf(strconst.SprintValueFormat, value)
}
