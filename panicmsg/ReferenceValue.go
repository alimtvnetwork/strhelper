package panicmsg

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/strconst"
)

type ReferenceValue struct {
	VariableName string
	Value        interface{}
}

func (referenceValue *ReferenceValue) String() string {
	return (*referenceValue).VariableName +
		strconst.SpaceColonSpace +
		fmt.Sprintf(strconst.SprintValueFormat, (*referenceValue).Value)
}

func (referenceValue *ReferenceValue) TypeString() string {
	return referenceValue.VariableName +
		squareBracketStart +
		fmt.Sprintf(strconst.SprintTypeFormat, referenceValue.Value) +
		squareBracketEnd +
		strconst.SpaceColonSpace +
		fmt.Sprintf(strconst.SprintValueFormat, referenceValue.Value)
}
