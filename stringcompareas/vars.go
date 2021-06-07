package stringcompareas

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreimpl/enumimpl"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/strhelper/isstr"
)

var (
	stringRanges = [...]string{
		Equal:      "Equal",
		StartsWith: "StartsWith",
		EndsWith:   "EndsWith",
		AnyWhere:   "AnyWhere",
		NotEqual:   "NotEqual",
	}

	typeRanges = [...]Variant{
		Equal:      Equal,
		StartsWith: StartsWith,
		EndsWith:   EndsWith,
		AnyWhere:   AnyWhere,
		NotEqual:   NotEqual,
	}

	basicEnumImpl = enumimpl.
			NewBasicByteUsingIndexedSlice(stringRanges[:])

	RangesInvalidError = errnew.NewPtr(
		errtype.OutOfRangeValue,
		basicEnumImpl.RangesInvalidErr())

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

	rangesMap = map[Variant]IsLineCompareFunc{
		Equal:      isEqualFunc,
		StartsWith: isStartsWithFunc,
		EndsWith:   isEndsWithFunc,
		AnyWhere:   isAnywhereFunc,
		NotEqual:   isNotEqualFunc,
	}
)
