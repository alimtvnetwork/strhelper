package panicmsg

import "gitlab.com/evatix-go/strhelper/constants"

// Returns variableName + "[" + typeName + "]" + constants.SpaceColonSpace + value
func TVar(typeName, variableName, value string) string {
	return variableName +
		squareBracketStart +
		typeName +
		squareBracketEnd +
		constants.SpaceColonSpace +
		value
}
