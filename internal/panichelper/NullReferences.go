package panichelper

import (
	"gitlab.com/evatix-go/strhelper/internal/panicmsg"
)

func NullReferences(nullReferenceNames ...string) {
	references := panicmsg.NullValueVariableNamesReference(nullReferenceNames...)

	message := panicmsg.SimpleValMsgsUsingReferencesWithType(
		"Cannot be nil. ",
		references,
	)

	panic(message)
}
