package panichelper

import (
	"gitlab.com/auk-go/strhelper/internal/panicmsg"
)

func SimplePanicVariableNames(isPanic bool, msg string, referencesNames ...string) {
	if !isPanic {
		return
	}

	references := panicmsg.EmptyValueVariableNamesReference(referencesNames...)

	message := panicmsg.SimpleValMsgsWithType(
		msg,
		*references...)

	panic(message)
}
