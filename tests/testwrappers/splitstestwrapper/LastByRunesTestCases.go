package splitstestwrapper

import "gitlab.com/auk-go/core/constants"

var LastByRunesTestCases = []LastByRunesWrapper{
	{
		Content: "\\found....1/...found...2/...found3",
		SearchingContents: []rune{
			'/',
			'\\',
		},
		IsCaseSensitive: true,
		Limits:          constants.MinusOne,
		HasPanic:        false,
		funcName:        lastByRunes,
		expected: []string{
			"...found3",
			"...found...2",
			"found....1",
			"",
		},
		actual: nil,
	},
	{
		Content: "/found....1/...found...2/...found3\\",
		SearchingContents: []rune{
			'/',
			'\\',
		},
		IsCaseSensitive: true,
		Limits:          constants.MinusOne,
		HasPanic:        false,
		funcName:        lastByRunes,
		expected: []string{
			"",
			"...found3",
			"...found...2",
			"found....1",
			"",
		},
	},
	{
		Content: "[ab]found....1[ab]...found...2[ab]...found3",
		SearchingContents: []rune{
			'/',
			'\\',
		},
		IsCaseSensitive: true,
		Limits:          constants.MinusOne,
		HasPanic:        false,
		funcName:        lastByRunes,
		expected:        []string{"[ab]found....1[ab]...found...2[ab]...found3"},
	},
	{
		Content: "[ab]found....1[ab]...found...2[ab]...found3[ab]",
		SearchingContents: []rune{
			'/',
			'\\',
		},
		IsCaseSensitive: true,
		Limits:          constants.MinusOne,
		HasPanic:        false,
		funcName:        lastByRunes,
		expected:        []string{"[ab]found....1[ab]...found...2[ab]...found3[ab]"},
	},
	{
		Content: "/found....1/...found...2\\...found3/",
		SearchingContents: []rune{
			'/',
			'\\',
		},
		IsCaseSensitive: true,
		Limits:          constants.Three,
		funcName:        lastByRunes,
		expected: []string{
			"",
			"...found3",
			"/found....1/...found...2",
		},
	},
	{
		Content: "\\found....1\\...found...2\\...found3/found4",
		SearchingContents: []rune{
			'/',
			'\\',
		},
		IsCaseSensitive: true,
		Limits:          constants.Three,
		funcName:        lastByRunes,
		expected: []string{
			"found4",
			"...found3",
			"\\found....1\\...found...2",
		},
	},
	{
		Content: "found 0/found....1\\...found...2/...found3/found4",
		SearchingContents: []rune{
			'/',
			'\\',
		},
		IsCaseSensitive: true,
		Limits:          constants.MinusOne,
		HasPanic:        false,
		funcName:        lastByRunes,
		expected: []string{
			"found4",
			"...found3",
			"...found...2",
			"found....1",
			"found 0",
		},
	},
	{
		Content: "found 0\\found....1/...found...2\\...found3/found4",
		SearchingContents: []rune{
			'/',
			'\\',
		},
		IsCaseSensitive: true,
		Limits:          constants.Two,
		HasPanic:        false,
		funcName:        lastByRunes,
		expected: []string{
			"found4",
			"found 0\\found....1/...found...2\\...found3",
		},
	},
	{
		Content: "বাংলাদেশ1/বাংলাদেশ2\\বাংলাদেশ3\\বাংলাদেশ4/বাংলাদেশ5",
		SearchingContents: []rune{
			'/',
			'\\',
		},
		IsCaseSensitive: true,
		Limits:          constants.Four,
		HasPanic:        false,
		funcName:        lastByRune,
		expected: []string{
			"বাংলাদেশ5",
			"বাংলাদেশ4",
			"বাংলাদেশ3",
			"বাংলাদেশ1/বাংলাদেশ2",
		},
	},
}
