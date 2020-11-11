package charhelper

// Make str to uint8 (char) array
// if str == nil or length == 0 then returns nil
func ToChars(str string) []uint8 {
	return *ToCharsPtr(&str)
}
