package stringsearch

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/strhelper/isstr"
)

var (
	isStartsWithFunc = func(
		contentLine,
		searchComparingLine string,
		isCaseSensitive bool,
	) bool {
		return isstr.StartsWith(
			contentLine,
			searchComparingLine,
			constants.Zero,
			isCaseSensitive)
	}

	isEndsWithFunc = func(
		contentLine,
		searchComparingLine string,
		isCaseSensitive bool,
	) bool {
		return isstr.EndsWithPtr(
			&contentLine,
			&searchComparingLine,
			constants.Zero,
			isCaseSensitive)
	}

	isAnywhereFunc = func(
		contentLine,
		searchComparingLine string,
		isCaseSensitive bool,
	) bool {
		return isstr.ContainsPtr(
			&contentLine,
			&searchComparingLine,
			constants.Zero,
			isCaseSensitive)
	}

	isEqualFunc = func(
		contentLine,
		searchComparingLine string,
		isCaseSensitive bool,
	) bool {
		if isCaseSensitive {
			return contentLine == searchComparingLine
		}

		return strings.EqualFold(searchComparingLine, contentLine)
	}

	isNotEqualFunc = func(
		contentLine,
		searchComparingLine string,
		isCaseSensitive bool,
	) bool {
		if isCaseSensitive {
			return contentLine != searchComparingLine
		}

		return !strings.EqualFold(searchComparingLine, contentLine)
	}
)
