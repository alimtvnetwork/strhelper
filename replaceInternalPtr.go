package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func replaceInternalPtr(request *strhelpercore.ReplaceRequest) string {
	if request.StartsAt == 0 && IsEmptyPtr(request.Text) && IsEmptyPtr(&request.Search) {
		return request.ReplaceWith
	}

	if request.StartsAt == 0 && (*request).IsCaseSensitive {
		return strings.Replace(
			*request.Text,
			request.Search,
			request.ReplaceWith,
			request.HowManyReplace)
	}

	foundIndexes := IndexesOfAllPtr(
		request.Text,
		&request.Search,
		request.StartsAt,
		(*request).IsCaseSensitive)

	if foundIndexes == nil {
		// returns as is
		return *request.Text
	}

	textLength := len(*request.Text)
	replaceCount := len(foundIndexes)
	isHowManyReplaceSet := (*request).HowManyReplace > -1

	if !isHowManyReplaceSet || isHowManyReplaceSet && replaceCount > (*request).HowManyReplace {
		replaceCount = request.HowManyReplace
	}

	// not found case
	if request.HowManyReplace == 0 {
		return *request.Text
	}

	newWordLength := len(request.ReplaceWith)
	searchLength := len(request.Search)

	// Apply replacements to buffer.
	chars := make([]byte, textLength+replaceCount*(newWordLength-searchLength))
	wordIndex := 0
	for i := 0; i < textLength; i++ {
		if indexOfInts(&foundIndexes, i) > -1 && replaceCount > 0 {
			// found modify
			wordIndex += copy(chars[wordIndex:], request.ReplaceWith)
			i += searchLength - 1 // we should skip the search text since already replaced.
			replaceCount--
			continue
		}

		// not found existing, keep as is
		chars[wordIndex] = (*(*request).Text)[i]
		wordIndex++
	}

	return string(chars)
}
