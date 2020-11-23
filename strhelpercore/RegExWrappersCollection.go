package strhelpercore

import "gitlab.com/evatix-go/strhelper/internal/pkg/misc"

type RegExWrappersCollection struct {
	expressions *[]string
	regexesMap  *map[string]*RegExWrapper
	regexes     *[]*RegExWrapper
	hasAnyError *bool
}

func NewRegExWrappersCollection(expressions *[]string) *RegExWrappersCollection {
	return &RegExWrappersCollection{
		expressions: expressions,
	}
}

func NewRegExWrappersCollectionUsingStringPointer(expressions *[]*string) *RegExWrappersCollection {
	parsableExpression := misc.ConvertPtrStringArrayToStringArray(expressions)

	return &RegExWrappersCollection{
		expressions: parsableExpression,
	}
}

// AddExpression expensive operation, better to add all expression at New.
func (regExWrappersCollection *RegExWrappersCollection) AddExpression(expression string) {
	*regExWrappersCollection.expressions = append(*regExWrappersCollection.expressions, expression)
	length := len(*regExWrappersCollection.expressions)
	regexes := regExWrappersCollection.Value()
	regexWrapper := NewRegExWrapper(length-1, &expression)
	*regexes = append(*regexes, regexWrapper)
	(*regExWrappersCollection.regexesMap)[expression] = regexWrapper
}

func (regExWrappersCollection *RegExWrappersCollection) Expressions() *[]string {
	return regExWrappersCollection.expressions
}

func (regExWrappersCollection *RegExWrappersCollection) Regexes() *[]*RegExWrapper {
	return regExWrappersCollection.Value()
}

func (regExWrappersCollection *RegExWrappersCollection) IsEmpty() bool {
	return regExWrappersCollection.expressions == nil || len(*regExWrappersCollection.expressions) == 0
}

func (regExWrappersCollection *RegExWrappersCollection) IsDefined() bool {
	return !(regExWrappersCollection.expressions == nil || len(*regExWrappersCollection.expressions) == 0)
}

func (regExWrappersCollection *RegExWrappersCollection) Length() int {
	return len(*regExWrappersCollection.expressions)
}

func (regExWrappersCollection *RegExWrappersCollection) Value() *[]*RegExWrapper {
	if regExWrappersCollection.regexes == nil && regExWrappersCollection.IsDefined() {
		length := regExWrappersCollection.Length()
		regexes := make([]*RegExWrapper, length, length*2)

		for i := 0; i < length; i++ {
			valueAt := (*regExWrappersCollection.expressions)[i]
			regexes[i] = NewRegExWrapper(i, &valueAt)
		}

		regExWrappersCollection.regexes = &regexes
	}

	return regExWrappersCollection.regexes
}

func (regExWrappersCollection *RegExWrappersCollection) RegexesMap() *map[string]*RegExWrapper {
	if regExWrappersCollection.regexesMap == nil && regExWrappersCollection.IsDefined() {
		length := regExWrappersCollection.Length()
		regexes := regExWrappersCollection.Value()
		regexesMap := make(map[string]*RegExWrapper, length)

		for i := 0; i < length; i++ {
			valueAt := (*regExWrappersCollection.expressions)[i]
			regexesMap[valueAt] = (*regexes)[i]
		}

		regExWrappersCollection.regexesMap = &regexesMap
	}

	return regExWrappersCollection.regexesMap
}
