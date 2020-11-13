package charhelper

// Returns Upper case runes by creating new runes. (Don't modify in place, thus requires more memory consumption)
//
// Invalid case (returns nil)
//  - if inputs == nil
func ToUpperRunes(inputs *[]rune) *[]rune {
	if inputs == nil {
		return nil
	}

	// Copying, example reference : https://play.golang.org/p/r65MrCg86YH
	newRunes := *inputs

	return ToUpperRunesInPlace(&newRunes)
}
