package strto

import "strconv"

// Returns defaultVal if any conversion error
func Uint8(str string, defaultVal uint8) uint8 {
	result, er := strconv.ParseUint(str, 10, 8)

	if er == nil {
		rs := uint8(result)

		return rs
	}

	return defaultVal
}

// Returns defaultVal if any conversion error
func Uint8Ptr(str *string, defaultVal uint8) uint8 {
	result, er := strconv.ParseUint(*str, 10, 8)

	if er == nil {
		rs := uint8(result)

		return rs
	}

	return defaultVal
}
