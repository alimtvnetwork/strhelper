package panicmsg

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/constants"
)

type ReferenceValue struct {
	VariableName string
	Value        interface{}
}

func (referenceValue *ReferenceValue) String() string {
	return (*referenceValue).VariableName +
		constants.SpaceColonSpace +
		fmt.Sprintf(constants.SprintValueFormat, (*referenceValue).Value)
}

func (referenceValue *ReferenceValue) TypeString() string {
	return referenceValue.VariableName +
		squareBracketStart +
		fmt.Sprintf(constants.SprintTypeFormat, referenceValue.Value) +
		squareBracketEnd +
		constants.SpaceColonSpace +
		fmt.Sprintf(constants.SprintValueFormat, referenceValue.Value)
}
