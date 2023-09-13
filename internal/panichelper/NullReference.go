package panichelper

import (
	"gitlab.com/auk-go/strhelper/internal/panicmsg"
)

func NullReference(nullReferenceName string) {
	message := panicmsg.SimpleValMsgsWithType(
		"Cannot be nil. ",
		panicmsg.ReferenceValue{
			VariableName: nullReferenceName,
			Value:        nullReferenceName,
		},
	)

	panic(message)
}
