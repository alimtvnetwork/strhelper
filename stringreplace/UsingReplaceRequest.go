package stringreplace

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/strhelpercore"
	"gitlab.com/evatix-go/strhelper/stringindex"
)

func UsingReplaceRequest(request *strhelpercore.ReplaceRequest) string {
	if request.StartsAt == 0 && request.Text == "" && request.Search == "" {
		return request.ReplaceWith
	}

	if request.StartsAt == 0 && request.IsCaseSensitive {
		return strings.Replace(
			request.Text,
			request.Search,
			request.ReplaceWith,
			request.HowManyReplace)
	}

	foundIndexesMap := stringindex.OfAllAsKeyMap(
		request.Text,
		request.Search,
		request.StartsAt,
		request.HowManyReplace,
		request.IsCaseSensitive)

	if len(foundIndexesMap) == 0 {
		// returns as is
		return request.Text
	}

	textLength := len(request.Text)
	replaceCount := len(foundIndexesMap)
	isHowManyReplaceSet := request.HowManyReplace > -1

	if isHowManyReplaceSet && replaceCount > request.HowManyReplace {
		replaceCount = request.HowManyReplace
	}

	// not found or nothing to replace case
	if request.HowManyReplace == 0 || replaceCount == 0 {
		return request.Text
	}

	newWordLength := len(request.ReplaceWith)
	searchLength := len(request.Search)

	// Apply replacements to buffer.
	chars := make([]byte, textLength+replaceCount*(newWordLength-searchLength))
	wordIndex := 0
	for i := 0; i < textLength; i++ {
		if foundIndexesMap[i] == true && replaceCount > 0 {
			// found modify
			wordIndex += copy(chars[wordIndex:], request.ReplaceWith)
			i += searchLength - 1 // we should skip the search text since already replaced.
			replaceCount--
			continue
		}

		// not found existing, keep as is
		chars[wordIndex] = request.Text[i]
		wordIndex++
	}

	return string(chars)
}
