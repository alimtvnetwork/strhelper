package stringnumtests

import "gitlab.com/auk-go/strhelper/stringnum"

var unicodeIntegerUnsignedTestCases = &numberVerifyGroupWrapper{
	methodName:   "IsUnicode.Integer.Unsigned",
	verifierFunc: stringnum.IsUnicode.Integer.Unsigned,
	testCases: []numberVerifyWrapper{
		{
			input:         "-১২৩৪৫",
			isValidNumber: false,
		},
		{
			input:         "+১২৩৪৫",
			isValidNumber: true,
		},
		{
			input:         "১২৩৪৫",
			isValidNumber: true,
		},
		{
			input:         "+১২৩.৪৫",
			isValidNumber: false,
		},
		{
			input:         "+১২৩.৪.৫",
			isValidNumber: false,
		},
		{
			input:         "-১২৩.৪৫000000000000000000000000000000000000",
			isValidNumber: false,
		},
		{
			input:         "১২৩.৪৫000000000000000000000000000000000000",
			isValidNumber: false,
		},
		{
			input:         "-১২৩.৪.৫",
			isValidNumber: false,
		},
		{
			input:         "+222111",
			isValidNumber: true,
		},
		{
			input:         "-222111",
			isValidNumber: false,
		},
		{
			input:         "222111",
			isValidNumber: true,
		},
		{
			input:         "22211100000000000000000000000000000000000000000000000000000000000000000000",
			isValidNumber: true,
		},
		{
			input:         "+222.111",
			isValidNumber: false,
		},
		{
			input:         "-222111.",
			isValidNumber: false,
		},
		{
			input:         ".222111",
			isValidNumber: false,
		},
		{
			input:         ".222.111",
			isValidNumber: false,
		},
		{
			input:         "",
			isValidNumber: false,
		},
		{
			input:         "-",
			isValidNumber: false,
		},
		{
			input:         "+",
			isValidNumber: false,
		},
		{
			input:         "-1",
			isValidNumber: false,
		},
		{
			input:         "+1",
			isValidNumber: true,
		},
		{
			input:         "0",
			isValidNumber: true,
		},
		{
			input:         "a",
			isValidNumber: false,
		},
	},
}
