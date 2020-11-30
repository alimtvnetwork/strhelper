package chars

// Make str to rune array
// if str == nil or length == 0 then returns nil
func ToRunes(str string) []rune {
	return *ToRunesPtr(&str)
}
