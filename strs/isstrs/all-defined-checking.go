package isstrs

import "strings"

// DefinedAll
//
//  return true if all the lines are
//       defined (!="")
func DefinedAll(lines ...string) bool {
	if len(lines) == 0 {
		return false
	}

	for _, line := range lines {
		if line == "" {
			return false
		}
	}

	return true
}

// DefinedAny
//
//  return true if any line is
//       defined (!="")
func DefinedAny(lines ...string) bool {
	if len(lines) == 0 {
		return false
	}

	for _, line := range lines {
		if line != "" {
			return true
		}
	}

	return false
}

// TrimDefinedAll
//
//  return true if all the lines are
//          defined (!="")
//  after strings.TrimSpace
func TrimDefinedAll(
	lines ...string,
) bool {
	if len(lines) == 0 {
		return false
	}

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			return false
		}
	}

	return true
}

// TrimDefinedAny
//
//  return true if any lines are
//         defined (!="")
//  after strings.TrimSpace
func TrimDefinedAny(
	lines ...string,
) bool {
	if len(lines) == 0 {
		return false
	}

	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			return true
		}
	}

	return false
}
