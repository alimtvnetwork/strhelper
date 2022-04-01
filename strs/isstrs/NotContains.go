package isstrs

func NotContains(lines []string, line string) bool {
	return !Contains(lines, line)
}
