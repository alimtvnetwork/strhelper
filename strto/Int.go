package strto

import "strconv"

func Int(str string, defaultInt int) int {
	toInt, er := strconv.Atoi(str)

	if er == nil {
		return toInt
	}

	return defaultInt
}
