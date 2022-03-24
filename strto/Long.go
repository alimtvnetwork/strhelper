package strto

import "strconv"

func Long(
	str string, defaultVal int64,
) int64 {
	toL, er := strconv.ParseInt(
		str, 10, 64)

	if er == nil {
		return toL
	}

	return defaultVal
}
