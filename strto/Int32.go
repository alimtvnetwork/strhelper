package strto

import "strconv"

// Returns defaultVal if any conversion error
func Int32(str string, defaultVal int32) int32 {
	result, er := strconv.ParseInt(str, 10, 32)

	if er == nil {
		rs := int32(result)

		return rs
	}

	return defaultVal
}

// Returns defaultVal if any conversion error
func Int32Ptr(str *string, defaultVal int32) int32 {
	result, er := strconv.ParseInt(*str, 10, 32)

	if er == nil {
		rs := int32(result)

		return rs
	}

	return defaultVal
}
