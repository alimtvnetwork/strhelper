package strto

import "strconv"

// Returns defaultVal if any conversion error
func Bool(str string, defaultVal bool) bool {
	result, er := strconv.ParseBool(str)

	if er == nil {
		return result
	}

	return defaultVal
}

// Returns defaultVal if any conversion error
func BoolPtr(str *string, defaultVal bool) bool {
	result, er := strconv.ParseBool(*str)

	if er == nil {
		return result
	}

	return defaultVal
}
