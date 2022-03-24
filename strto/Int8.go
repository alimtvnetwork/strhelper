package strto

import "strconv"

// Int8
//
// returns defaultVal if any conversion error
func Int8(str string, defaultVal int8) int8 {
	result, er := strconv.ParseInt(str, 10, 8)

	if er == nil {
		rs := int8(result)

		return rs
	}

	return defaultVal
}
