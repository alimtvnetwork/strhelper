package strto

import "strconv"

// Returns defaultVal if any conversion error
func Uint64(str string, defaultVal uint64) uint64 {
	result, er := strconv.ParseUint(str, 10, 64)

	if er == nil {
		return result
	}

	return defaultVal
}

// Returns defaultVal if any conversion error
func Uint64Ptr(str *string, defaultVal uint64) *uint64 {
	result, er := strconv.ParseUint(*str, 10, 64)

	if er == nil {
		return &result
	}

	return &defaultVal
}
