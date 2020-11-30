package panicmsg

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns variableName + "[" + typeName + "]" + constants.SpaceColonSpace + value
func VarWithType(typeName, variableName, value string) string {
	return variableName +
		squareBracketStart +
		typeName +
		squareBracketEnd +
		strconst.SpaceColonSpace +
		value
}
