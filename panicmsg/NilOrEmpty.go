package panicmsg

// Returns "Cannot be nil or null. Reference ( " + Var(variableName, "nil") + " )"
func NilOrEmpty(variableName string, value interface{}) string {
	return SimpleValMsg(CannotBeNilOrEmptyMessage, variableName, value)
}
