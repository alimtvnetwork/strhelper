package panicmsg

// NullOrEmpty returns "Cannot be nil or null. Reference ( " + Var(variableName, "nil") + " )"
func NullOrEmpty(variableName string, value interface{}) string {
	return SimpleValMsg(CannotBeNilOrEmptyMessage, variableName, value)
}
