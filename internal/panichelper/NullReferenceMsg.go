package panichelper

import (
	"gitlab.com/evatix-go/strhelper/internal/panicmsg"
)

func NullReferenceMsg(msg, nullReferenceName string) {
	message := panicmsg.SimpleValMsgsWithType(
		msg,
		panicmsg.ReferenceValue{
			VariableName: nullReferenceName,
			Value:        nullReferenceName,
		},
	)

	panic(message)
}
