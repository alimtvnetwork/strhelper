package panichelper

import "gitlab.com/evatix-go/strhelper/panicmsg"

func LastIndexIncreasedByFailed(lastIndexIncreasedBy, contentLength int) {
	message := panicmsg.TSimpleValMsgs(
		"lastIndexIncreasedBy cannot be negative or more than the length of content.",
		panicmsg.ReferenceValue{
			VariableName: "lastIndexIncreasedBy",
			Value:        lastIndexIncreasedBy,
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
