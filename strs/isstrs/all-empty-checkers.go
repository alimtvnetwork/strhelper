package isstrs

import (
	"strings"
)

// EmptyAll
//
//  return true if all the lines are empty
func EmptyAll(lines ...string) bool {
	if len(lines) == 0 {
		return true
	}

	for _, line := range lines {
		if line != "" {
			return false
		}
	}

	return true
}

// TrimEmptyAll
//
//  return true if all the lines are empty
//  after strings.TrimSpace
func TrimEmptyAll(lines ...string) bool {
	if len(lines) == 0 {
		return true
	}

	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			return false
		}
	}

	return true
}

// TrimEmptyAny
//
//  return true if any line is empty
//  after strings.TrimSpace
func TrimEmptyAny(lines ...string) bool {
	if len(lines) == 0 {
		return true
	}

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			return true
		}
	}

	return false
}
