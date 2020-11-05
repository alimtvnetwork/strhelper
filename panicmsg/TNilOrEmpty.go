package panicmsg

// Returns "Cannot be nil or null. Reference ( " + TVar(typename, variableName, "nil") + " )"
// Type name included
func TNilOrEmpty(variableName string, value interface{}) string {
	return TSimpleValMsg(CannotBeNilOrEmptyMessage, variableName, value)
}
