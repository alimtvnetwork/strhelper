package substr

// Use language integrated one that will be faster. eg. str[startsAtIndex : startsAtIndex+length]
// If want to use by function then pointer one would be more efficient than this one.
func ByLength(str string, startsAtIndex, length int) string {
	return str[startsAtIndex : startsAtIndex+length]
}
