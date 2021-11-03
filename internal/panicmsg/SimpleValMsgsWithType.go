package panicmsg

// SimpleValMsgsWithType
//
// Returns msg + referenceStart + VarWithType(typeName, variableName, printVal) + spaceParenthesisEnd
// Type name included
func SimpleValMsgsWithType(msg string, referenceValues ...ReferenceValue) string {
	return SimpleValMsgsUsingReferencesWithType(msg, &referenceValues)
}
