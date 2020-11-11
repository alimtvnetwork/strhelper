package panicmsg

// Returns "Cannot be negative value. Reference ( " + TVar(typeName, variableName, printVal) + " )"
// Type name included
func TNegativeValue(variableName string, value interface{}) string {
	return TSimpleValMsg(CannotBeNegativeMessage, variableName, value)
}
