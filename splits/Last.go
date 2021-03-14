package splits

import (
	"strings"

	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/constants"
)

func Last(s, separator string) *[]string {
	if s == "" {
		return core.EmptyStringsPtr()
	}
	strings.Split()
	return LastByLimitPtr(
		&s,
		&separator,
		constants.MinusOne)
}

func LastByLimitPtr(s, separator *string, take int) *[]string {
	if s == nil || *s == "" {
		return core.EmptyStringsPtr()
	}

	isSepEmpty :=
		separator == nil ||
			*separator == ""

	if isSepEmpty && take == constants.MinusOne {
		emptySeparatorResults := strings.Split(*s, "")

		return &emptySeparatorResults
	}

	if isSepEmpty {
		// TODO every char in array
		emptySeparatorResults := strings.Split(*s, "")

		return &emptySeparatorResults
	}

	length := len(*s)
	separatorLength := len(*separator)

	if separatorLength > length {
		return core.EmptyStringsPtr()
	}

	limit := constants.ArbitraryCapacity5

	if take > 0 {
		limit = take
	}

	splitResults := make(
		[]string,
		0,
		limit)

	found := 0
	for i := length - 1; i >= 0; i-- {

	}
}
