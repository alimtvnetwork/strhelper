package strs

import "fmt"

func StringersToSlice(stringers ...fmt.Stringer) []string {
	if len(stringers) == 0 {
		return []string{}
	}

	lines := make([]string, len(stringers))

	for i, stringer := range stringers {
		lines[i] = stringer.String()
	}

	return lines
}
