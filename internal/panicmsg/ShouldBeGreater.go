package panicmsg

// ShouldBeGreater
//
// Returns SimpleValMsg(ShouldBeGreaterThanMessage, variableName, numberValue)
func ShouldBeGreater(variableName string, numberValue int) string {
	return SimpleValMsg(ShouldBeGreaterThanMessage, variableName, numberValue)
}
