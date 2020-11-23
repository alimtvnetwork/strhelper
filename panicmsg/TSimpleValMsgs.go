package panicmsg

// Returns msg + referenceStart + TVar(typeName, variableName, printVal) + spaceParenthesisEnd
// Type name included
func TSimpleValMsgs(msg string, referenceValues ...ReferenceValue) string {
	return TSimpleValMsgsUsingArray(msg, &referenceValues)
}
