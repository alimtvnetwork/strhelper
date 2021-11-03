package panicmsg

// ShouldBeGreaterWithType
//
// Returns SimpleValMsgWithType(ShouldBeGreaterThanMessage, variableName, numberValue)
// Type name included
func ShouldBeGreaterWithType(variableName string, numberValue interface{}) string {
	return SimpleValMsgWithType(ShouldBeGreaterThanMessage, variableName, numberValue)
}
