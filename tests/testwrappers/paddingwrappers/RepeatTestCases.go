package paddingwrappers

import "gitlab.com/evatix-go/core/constants"

var RepeatTestCases = []RepeatWrapper{
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
		Content:     "[a]",
		RepeatWidth: constants.N7,
		HasPanic:    false,
		funcName:    repeatPtr,
		expected:    "[a][a][a][a][a][a][a]",
	},
	{
		Content:     "([a]*2)",
		RepeatWidth: constants.N10,
		HasPanic:    false,
		funcName:    repeatPtr,
		expected:    "([a]*2)([a]*2)([a]*2)([a]*2)([a]*2)([a]*2)([a]*2)([a]*2)([a]*2)([a]*2)",
	},
	{
		Content:     " ",
		RepeatWidth: constants.N10,
		HasPanic:    false,
		funcName:    repeatPtr,
		expected:    "          ",
	},
	{
		Content:     " \t\n\v",
		RepeatWidth: constants.N10,
		HasPanic:    false,
		funcName:    repeatPtr,
		expected:    " \t\n\v \t\n\v \t\n\v \t\n\v \t\n\v \t\n\v \t\n\v \t\n\v \t\n\v \t\n\v",
	},
	{
		Content:     "\nবাংলাদেশ",
		RepeatWidth: constants.N6,
		HasPanic:    false,
		funcName:    repeatPtr,
		expected:    "\nবাংলাদেশ\nবাংলাদেশ\nবাংলাদেশ\nবাংলাদেশ\nবাংলাদেশ\nবাংলাদেশ",
	},
}
