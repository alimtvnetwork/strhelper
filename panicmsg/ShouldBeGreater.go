package panicmsg

// Returns SimpleValMsg(ShouldBeGreaterThanMessage, variableName, numberValue)
func ShouldBeGreater(variableName string, numberValue int) string {
	return SimpleValMsg(ShouldBeGreaterThanMessage, variableName, numberValue)
}
