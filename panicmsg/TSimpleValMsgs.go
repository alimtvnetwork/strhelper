package panicmsg

import (
	"gitlab.com/evatix-go/strhelper/concat"
	"gitlab.com/evatix-go/strhelper/constants"
)

// Returns msg + referenceStart + TVar(typeName, variableName, printVal) + spaceParenthesisEnd
// Type name included
func TSimpleValMsgs(msg string, referenceValues ...ReferenceValue) string {
	var printVal string

	if referenceValues == nil || len(referenceValues) == 0 {
		printVal = constants.NilString
	} else {
		stringsArray := make([]string, len(referenceValues))

		for i, value := range referenceValues {
			stringsArray[i] = value.TypeString()
		}

		printVal = concat.JoinPtr(
			&stringsArray,
			constants.CommaSpacePtr)
	}

	return msg +
		referenceStart +
		printVal +
		spaceParenthesisEnd
}
