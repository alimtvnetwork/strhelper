package strhelpercore

import "gitlab.com/evatix-go/strhelper/internal/misc"

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
	parsableExpression := misc.ConvertPointerStringsToStrings(expressions)

	return &RegExWrappersCollection{
		expressions: parsableExpression,
	}
}

// AddExpression expensive operation, better to add all expression at New.
func (regExWrappersCollection *RegExWrappersCollection) AddExpression(expression string) {
	// This is correct, as the item is set on the New
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

func (regExWrappersCollection *RegExWrappersCollection) IsEquals(
	anotherRegExCollection *RegExWrappersCollection,
) bool {
	if anotherRegExCollection == nil {
		return false
	}

	if anotherRegExCollection.IsEmpty() && regExWrappersCollection.IsEmpty() {
		return true
	}

	if anotherRegExCollection.IsEmpty() || regExWrappersCollection.IsEmpty() {
		return false
	}

	if anotherRegExCollection.Length() != regExWrappersCollection.Length() {
		return false
	}

	for i, regExWrapper := range *anotherRegExCollection.regexes {
		current := (*regExWrappersCollection.regexes)[i]

		if regExWrapper == nil && current == nil {
			continue
		}

		if regExWrapper == nil || current == nil {
			return false
		}

		if !current.IsEquals(regExWrapper) {
			return false
		}
	}

	return true
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
