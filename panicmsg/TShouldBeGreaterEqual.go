package panicmsg

// Returns TSimpleValMsg(ShouldBeGreaterThanEqualMessage, variableName, numberValue)
// Type name included
func TShouldBeGreaterEqual(variableName string, numberValue interface{}) string {
	return TSimpleValMsg(ShouldBeGreaterThanEqualMessage, variableName, numberValue)
}
