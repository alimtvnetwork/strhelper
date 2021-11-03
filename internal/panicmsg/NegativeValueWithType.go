package panicmsg

// Returns "Cannot be negative value. Reference ( " + VarWithType(typeName, variableName, printVal) + " )"
// Type name included
func NegativeValueWithType(variableName string, value interface{}) string {
	return SimpleValMsgWithType(CannotBeNegativeMessage, variableName, value)
}
