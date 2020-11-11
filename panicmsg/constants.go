package panicmsg

const (
	errorStart          = "Error : "
	referenceStart      = " Reference ( "
	spaceParenthesisEnd = " )"
	squareBracketStart  = "["
	squareBracketEnd    = "]"
)

// messages
const (
	CannotBeNegativeMessage         = "Cannot be negative value."
	CannotBeNilOrEmptyMessage       = "Cannot be nil or null or empty."
	CannotBeNilMessage              = "Cannot be nil or null."
	ShouldBeLessThanMessage         = "Should be less than the reference."
	ShouldBeGreaterThanMessage      = "Should be greater than the reference."
	ShouldBeLessThanEqualMessage    = "Should be less or equal to the reference."
	ShouldBeEqualToMessage          = "Should be equal to the reference."
	ShouldBeGreaterThanEqualMessage = "Should be greater or equal to the reference."
)
