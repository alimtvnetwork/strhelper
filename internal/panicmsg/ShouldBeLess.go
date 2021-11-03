package panicmsg

// ShouldBeLess
//
// Returns SimpleValMsg(ShouldBeLessThan, variableName, numberValue)
func ShouldBeLess(variableName string, numberValue int) string {
	return SimpleValMsg(ShouldBeLessThanMessage, variableName, numberValue)
}
