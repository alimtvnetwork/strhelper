package panicmsg

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// SimpleValMsgsUsingReferencesWithType
//
// Returns msg + referenceStart + VarWithType(typeName, variableName, printVal) + spaceParenthesisEnd
// Type name included
func SimpleValMsgsUsingReferencesWithType(msg string, referenceValues *[]ReferenceValue) string {
	var printVal string

	if referenceValues == nil || len(*referenceValues) == 0 {
		printVal = constants.NilString
	} else {
		stringsArray := make([]string, len(*referenceValues))

		for i, value := range *referenceValues {
			stringsArray[i] = value.TypeString()
		}

		printVal = strings.Join(
			stringsArray,
			constants.CommaSpace)
	}

	return msg +
		referenceStart +
		printVal +
		spaceParenthesisEnd
}
