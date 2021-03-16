package splitstestwrapper

import "gitlab.com/evatix-go/core/constants"

var LastCases = &[]GenericSplit{
	{
		Content:          "[ab]found....1[ab]...found...2[ab]...found3",
		SearchingContent: "[ab]",
		Limits:           constants.MinusOne,
		funcName:         last,
		expected: &[]string{
			"...found3",
			"...found...2",
			"found....1",
			"",
		},
	},
	{
		Content:          "[ab]found....1[ab]...found...2[ab]...found3",
		SearchingContent: "[ab]*",
		Limits:           constants.MinusOne,
		funcName:         last,
		expected:         &[]string{},
	},
	{
		Content:          "[ab]found....1[ab]...found...2[ab]...found3[ab]",
		SearchingContent: "[aB]",
		Limits:           constants.MinusOne,
		funcName:         last,
		expected: &[]string{
			"...found3",
			"...found...2",
			"found....1",
			"",
		},
	},
}
