package panichelper

import "gitlab.com/evatix-go/strhelper/panicmsg"

func LastIndexIncreasedByFailed(contentLengthDecreasedBy, contentLength int) {
	message := panicmsg.SimpleValMsgsWithType(
		"contentLengthDecreasedBy cannot be negative or more than the length of content.",
		panicmsg.ReferenceValue{
			VariableName: "contentLengthDecreasedBy",
			Value:        contentLengthDecreasedBy,
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
