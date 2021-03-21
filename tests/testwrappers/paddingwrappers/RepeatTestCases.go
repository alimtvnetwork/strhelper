package paddingwrappers

import "gitlab.com/evatix-go/core/constants"

var RepeatTestCases = []RepeatWrapper{
	{
		Content:     "[ab]",
		RepeatWidth: constants.N5,
		HasPanic:    false,
		funcName:    repeatPtr,
		expected:    "[ab][ab][ab][ab][ab]",
	},
	{
		Content:     "",
		RepeatWidth: constants.N5,
		HasPanic:    false,
		funcName:    repeatPtr,
		expected:    "",
	},
	{
		Content:     "[a]",
		RepeatWidth: constants.N4,
		HasPanic:    false,
		funcName:    repeatPtr,
		expected:    "[a][a][a][a]",
	},
	{
		Content:     " ",
		RepeatWidth: constants.N2,
		HasPanic:    false,
		funcName:    repeatPtr,
		expected:    "  ",
	},
	{
		Content:     "\n\t",
		RepeatWidth: constants.N2,
		HasPanic:    false,
		funcName:    repeatPtr,
		expected:    "\n\t\n\t",
	},
	{
		Content:     "\n\t\v",
		RepeatWidth: constants.N2,
		HasPanic:    false,
		funcName:    repeatPtr,
		expected:    "\n\t\v\n\t\v",
	},
}
