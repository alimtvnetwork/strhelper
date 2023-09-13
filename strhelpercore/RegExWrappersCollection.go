package strhelpercore

import (
	"gitlab.com/auk-go/core/issetter"
)

type RegExWrappersCollection struct {
	expressions []string
	regexesMap  map[string]*RegExWrapper
	regexes     []*RegExWrapper
	hasAnyError issetter.Value
}

func NewRegExWrappersCollection(expressions ...string) *RegExWrappersCollection {
	return &RegExWrappersCollection{
		expressions: expressions,
	}
}

// AddExpression expensive operation, better to add all expression at New.
func (it *RegExWrappersCollection) AddExpression(expression string) {
	// This is correct, as the item is set on the New
	it.expressions = append(
		it.expressions,
		expression)
	length := len(it.expressions)
	regexes := it.Value()
	regexWrapper := NewRegExWrapper(length-1, expression)
	regexes = append(regexes, regexWrapper)
	it.regexesMap[expression] = regexWrapper
}

func (it *RegExWrappersCollection) Expressions() []string {
	return it.expressions
}

func (it *RegExWrappersCollection) Regexes() []*RegExWrapper {
	return it.Value()
}

func (it *RegExWrappersCollection) IsExpressionEmpty() bool {
	return it.ExpressionLength() == 0
}

func (it *RegExWrappersCollection) IsExpressionDefined() bool {
	return it.ExpressionLength() > 0
}

func (it *RegExWrappersCollection) ExpressionLength() int {
	if it == nil || it.expressions == nil {
		return 0
	}

	return len(it.expressions)
}

func (it *RegExWrappersCollection) IsEquals(
	anotherRegExCollection *RegExWrappersCollection,
) bool {
	if anotherRegExCollection == nil && it == nil {
		return true
	}

	if anotherRegExCollection == nil || it == nil {
		return false
	}

	if anotherRegExCollection == it {
		return true
	}

	if anotherRegExCollection.IsExpressionEmpty() && it.IsExpressionEmpty() {
		return true
	}

	if anotherRegExCollection.IsExpressionEmpty() || it.IsExpressionEmpty() {
		return false
	}

	if anotherRegExCollection.ExpressionLength() != it.ExpressionLength() {
		return false
	}

	for i, regExWrapper := range anotherRegExCollection.regexes {
		current := it.regexes[i]

		if !current.IsEquals(regExWrapper) {
			return false
		}
	}

	return true
}

func (it *RegExWrappersCollection) Value() []*RegExWrapper {
	if it.regexes == nil && it.IsExpressionDefined() {
		length := it.ExpressionLength()
		regexes := make([]*RegExWrapper, length, length*2)

		for i := 0; i < length; i++ {
			valueAt := it.expressions[i]
			regexes[i] = NewRegExWrapper(i, valueAt)
		}

		it.regexes = regexes
	}

	return it.regexes
}

func (it *RegExWrappersCollection) RegexesMap() map[string]*RegExWrapper {
	if it == nil {
		return nil
	}

	if it.regexesMap != nil {
		return it.regexesMap
	}

	if it.IsExpressionDefined() {
		length := it.ExpressionLength()
		regexes := it.Value()
		regexesMap := make(map[string]*RegExWrapper, length)

		for i := 0; i < length; i++ {
			valueAt := it.expressions[i]
			regexesMap[valueAt] = regexes[i]
		}

		it.regexesMap = regexesMap
	} else {
		it.regexesMap = map[string]*RegExWrapper{}
	}

	return it.regexesMap
}
