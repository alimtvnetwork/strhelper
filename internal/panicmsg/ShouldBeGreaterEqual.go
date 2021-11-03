package panicmsg

// ShouldBeGreaterEqual
//
// Returns SimpleValMsg(ShouldBeGreaterThanEqualMessage, variableName, numberValue)
func ShouldBeGreaterEqual(variableName string, numberValue int) string {
	return SimpleValMsg(ShouldBeGreaterThanEqualMessage, variableName, numberValue)
}
