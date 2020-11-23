package panicmsg

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/strconst"
)

// Returns msg + referenceStart + TVar(typeName, variableName, printVal) + spaceParenthesisEnd
// Type name included
func TSimpleValMsgsUsingArray(msg string, referenceValues *[]ReferenceValue) string {
	var printVal string

	if referenceValues == nil || len(*referenceValues) == 0 {
		printVal = strconst.NilString
	} else {
		stringsArray := make([]string, len(*referenceValues))

		for i, value := range *referenceValues {
			stringsArray[i] = value.TypeString()
		}

		printVal = strings.Join(
			stringsArray,
			strconst.CommaSpace)
	}

	return msg +
		referenceStart +
		printVal +
		spaceParenthesisEnd
}
