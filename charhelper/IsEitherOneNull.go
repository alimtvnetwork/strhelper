package charhelper

func IsEitherOneNull(char1 *uint8, char2 *uint8) bool {
	return (char1 == nil && char2 != nil) ||
		(char2 == nil && char1 != nil)
}
