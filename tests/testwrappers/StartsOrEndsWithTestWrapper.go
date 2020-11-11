package testwrappers

type funcName string

const (
	isEndsWith   funcName = "IsEndsWith"
	isStartsWith funcName = "IsStartsWith"
)

func (funcName funcName) Value() string {
	return string(funcName)
}

type StartsOrEndsWithTestWrapper struct {
	WholeText       string
	Search          string
	StartsAt        int
	IsCaseSensitive bool
	funcName        funcName
	expected        bool
	actual          bool
}

func (startsOrEndsWithTestWrapper StartsOrEndsWithTestWrapper) Actual() interface{} {
	return startsOrEndsWithTestWrapper.actual
}

func (startsOrEndsWithTestWrapper *StartsOrEndsWithTestWrapper) SetActual(actual bool) {
	startsOrEndsWithTestWrapper.actual = actual
}

func (startsOrEndsWithTestWrapper StartsOrEndsWithTestWrapper) FuncName() string {
	return startsOrEndsWithTestWrapper.funcName.Value()
}

func (startsOrEndsWithTestWrapper StartsOrEndsWithTestWrapper) Value() interface{} {
	return startsOrEndsWithTestWrapper
}

func (startsOrEndsWithTestWrapper StartsOrEndsWithTestWrapper) Expected() interface{} {
	return startsOrEndsWithTestWrapper.expected
}

var EndsWithTestCases = []StartsOrEndsWithTestWrapper{
	{
		WholeText:       "Alim Ul Karim",
		Search:          "Karim",
		StartsAt:        0,
		IsCaseSensitive: true,
		expected:        true,
		funcName:        isEndsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "karim",
		StartsAt:        0,
		IsCaseSensitive: false,
		expected:        true,
		funcName:        isEndsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "kKarim",
		StartsAt:        0,
		IsCaseSensitive: true,
		expected:        false,
		funcName:        isEndsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "NothingNothingNothingNothing    %&^@(*)_)+",
		StartsAt:        0,
		IsCaseSensitive: true,
		expected:        false,
		funcName:        isEndsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          " ",
		StartsAt:        0,
		IsCaseSensitive: true,
		expected:        false,
		funcName:        isEndsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "kar",
		StartsAt:        2,
		IsCaseSensitive: false,
		expected:        true,
		funcName:        isEndsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "",
		StartsAt:        5,
		IsCaseSensitive: false,
		expected:        true,
		funcName:        isEndsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "",
		StartsAt:        15,
		IsCaseSensitive: false,
		expected:        false,
		funcName:        isEndsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "",
		StartsAt:        2,
		IsCaseSensitive: false,
		expected:        true,
	},
	{
		WholeText:       "",
		Search:          "",
		StartsAt:        0,
		IsCaseSensitive: false,
		expected:        true,
		funcName:        isEndsWith,
	},
	{
		WholeText:       "",
		Search:          "",
		StartsAt:        2,
		IsCaseSensitive: false,
		expected:        false,
		funcName:        isEndsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "karim",
		StartsAt:        0,
		IsCaseSensitive: true,
		expected:        false,
		funcName:        isEndsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "AAlim",
		StartsAt:        0,
		IsCaseSensitive: true,
		expected:        false,
		funcName:        isEndsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "Naureen",
		StartsAt:        0,
		IsCaseSensitive: true,
		expected:        false,
		funcName:        isEndsWith,
	},
}

var StartsWithTestCases = []StartsOrEndsWithTestWrapper{
	{
		WholeText:       "Alim Ul Karim",
		Search:          "Karim",
		StartsAt:        0,
		IsCaseSensitive: true,
		expected:        false,
		funcName:        isStartsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "alim",
		StartsAt:        0,
		IsCaseSensitive: false,
		expected:        true,
		funcName:        isStartsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "aalim",
		StartsAt:        0,
		IsCaseSensitive: false,
		expected:        false,
		funcName:        isStartsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "Alim",
		StartsAt:        0,
		IsCaseSensitive: true,
		expected:        true,
		funcName:        isStartsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "Alim ul Karim",
		StartsAt:        0,
		IsCaseSensitive: false,
		expected:        true,
		funcName:        isStartsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "Alim Ul Karim",
		StartsAt:        0,
		IsCaseSensitive: true,
		expected:        true,
		funcName:        isStartsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "Alim ",
		StartsAt:        0,
		IsCaseSensitive: true,
		expected:        true,
		funcName:        isStartsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "NothingNothingNothingNothing    %&^@(*)_)+",
		StartsAt:        0,
		IsCaseSensitive: true,
		expected:        false,
		funcName:        isStartsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          " ",
		StartsAt:        0,
		IsCaseSensitive: true,
		expected:        false,
		funcName:        isStartsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "Im uL",
		StartsAt:        2,
		IsCaseSensitive: false,
		expected:        true,
		funcName:        isStartsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "",
		StartsAt:        2,
		IsCaseSensitive: false,
		expected:        true,
		funcName:        isStartsWith,
	},
	{
		WholeText:       "Alim Ul Karim",
		Search:          "",
		StartsAt:        15,
		IsCaseSensitive: false,
		expected:        false,
		funcName:        isStartsWith,
	},
	{
		WholeText:       "",
		Search:          "",
		StartsAt:        0,
		IsCaseSensitive: false,
		expected:        true,
		funcName:        isStartsWith,
	},
	{
		WholeText:       "     ",
		Search:          "",
		StartsAt:        2,
		IsCaseSensitive: false,
		expected:        true,
		funcName:        isStartsWith,
	},
	{
		WholeText:       "Naureen Afrose review the code",
		Search:          "did Alim reviewed?",
		StartsAt:        0,
		IsCaseSensitive: true,
		expected:        false,
		funcName:        isStartsWith,
	},
}
