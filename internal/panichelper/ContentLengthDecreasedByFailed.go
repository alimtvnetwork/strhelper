package panichelper

import "gitlab.com/evatix-go/strhelper/panicmsg"

func ContentLengthDecreasedByFailed(contentLengthDecreasedBy, wholeTextLength int) {
	message := panicmsg.SimpleValMsgsWithType(
		"contentLengthDecreasedBy cannot be negative or more than the length of content.",
		panicmsg.ReferenceValue{
			VariableName: "contentLengthDecreasedBy",
			Value:        contentLengthDecreasedBy,
		},
		panicmsg.ReferenceValue{
			VariableName: "wholeTextLength",
			Value:        wholeTextLength,
		},
		panicmsg.ReferenceValue{
			VariableName: "contentLastIndex",
			Value:        wholeTextLength - 1,
		})

	panic(message)
}
