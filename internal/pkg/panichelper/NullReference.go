package panichelper

import "gitlab.com/evatix-go/strhelper/panicmsg"

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
