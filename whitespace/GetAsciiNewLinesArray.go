package whitespace

// FormFeed \f is also marked as newline here.
func GetAsciiNewLinesArray() [256]uint8 {
	return asciiNewLinesChars
}
