package isstrs

import "gitlab.com/evatix-go/strhelper/internal/whitespacesinternal"

func EmptyOrFirstWhitespace(lines *[]string) bool {
	isEmpty := lines == nil || *lines == nil || len(*lines) == 0

	if !isEmpty && lines != nil {
		return whitespacesinternal.IsWhitespaces(&(*lines)[0])
	}

	return isEmpty
}
