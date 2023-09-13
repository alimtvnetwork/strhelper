package panichelper

import (
	"gitlab.com/auk-go/strhelper/internal/panicmsg"
)

func StartAtIndexFailed(startsAtIndex, contentLength int) {
	message := panicmsg.SimpleValMsgsWithType(
		"startsAtIndex cannot be negative or more than the length of content.",
		panicmsg.ReferenceValue{
			VariableName: "startsAtIndex",
			Value:        startsAtIndex,
		},
		panicmsg.ReferenceValue{
			VariableName: "contentLength",
			Value:        contentLength,
		},
		panicmsg.ReferenceValue{
			VariableName: "contentLastIndex",
			Value:        contentLength - 1,
		})

	panic(message)
}
