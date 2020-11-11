package charhelper

// Make str to uint8 (char) array
// if str == nil or length == 0 then returns nil
//
// Reference : https://stackoverflow.com/a/28848879
func ToCharsPtr(str *string) *[]uint8 {
	newChars := []uint8(*str)

	return &newChars
}

