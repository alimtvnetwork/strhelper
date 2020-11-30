package panicmsg

// Returns "Cannot be nil or null. Reference ( " + VarWithType(typename, variableName, "nil") + " )"
// Type name included
func NullOrEmptyWithType(variableName string, value interface{}) string {
	return SimpleValMsgWithType(CannotBeNilOrEmptyMessage, variableName, value)
}
