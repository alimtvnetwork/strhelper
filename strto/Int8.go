package strto

import "strconv"

// Returns defaultVal if any conversion error
func Int8(str string, defaultVal int8) int8 {
	result, er := strconv.ParseInt(str, 10, 8)

	if er == nil {
		rs := int8(result)

		return rs
	}

	return defaultVal
}

// Returns defaultVal if any conversion error
func Int8Ptr(str *string, defaultVal int8) int8 {
	result, er := strconv.ParseInt(*str, 10, 8)

	if er == nil {
		rs := int8(result)

		return rs
	}

	return defaultVal
}
