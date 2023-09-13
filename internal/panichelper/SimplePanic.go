package panichelper

import (
	"gitlab.com/auk-go/strhelper/internal/panicmsg"
)

func SimplePanic(isPanic bool, msg string, references ...panicmsg.ReferenceValue) {
	if !isPanic {
		return
	}

	message := panicmsg.SimpleValMsgsWithType(
		msg,
		references...)

	panic(message)
}
