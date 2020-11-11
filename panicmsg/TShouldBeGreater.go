package panicmsg

// Returns TSimpleValMsg(ShouldBeGreaterThanMessage, variableName, numberValue)
// Type name included
func TShouldBeGreater(variableName string, numberValue interface{}) string {
	return TSimpleValMsg(ShouldBeGreaterThanMessage, variableName, numberValue)
}
