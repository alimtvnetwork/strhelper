package panichelper

import "gitlab.com/evatix-go/strhelper/panicmsg"

func NullReferences(nullReferenceNames ...string) {
	references := panicmsg.NilValueVariableNamesReference(nullReferenceNames...)

	message := panicmsg.TSimpleValMsgsUsingArray(
		"Cannot be nil. ",
		references,
	)

	panic(message)
}
