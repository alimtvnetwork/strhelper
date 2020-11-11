package panicmsg

// Returns "Cannot be negative value. " + referenceStart + Var(variableName, printVal) + spaceParenthesisEnd
func NegativeValue(variableName string, value interface{}) string {
	return SimpleValMsg(CannotBeNegativeMessage, variableName, value)
}
