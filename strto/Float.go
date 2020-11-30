package strto

import "strconv"

func Float(str string, defaultVal float64) float64 {
	toF, er := strconv.ParseFloat(str, 64)

	if er == nil {
		return toF
	}

	return defaultVal
}

// Returns defaultVal if any conversion error
func FloatPtr(str *string, defaultVal float64) *float64 {
	toF, er := strconv.ParseFloat(*str, 64)

	if er == nil {
		return &toF
	}

	return &defaultVal
}
