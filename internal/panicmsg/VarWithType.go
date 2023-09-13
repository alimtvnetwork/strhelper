package panicmsg

import (
	"gitlab.com/auk-go/core/constants"
)

// Returns variableName + "[" + typeName + "]" + constants.SpaceColonSpace + value
func VarWithType(typeName, variableName, value string) string {
	return variableName +
		squareBracketStart +
		typeName +
		squareBracketEnd +
		constants.SpaceColonSpace +
		value
}
