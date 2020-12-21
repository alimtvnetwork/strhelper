package strcompare

type DiffLine struct {
	LeftLine  Line
	RightLine Line
	LineEqual
}

func NewDiffLine(left, right Line) DiffLine {
	return DiffLine{
		LeftLine:  left,
		RightLine: right,
		LineEqual: NewLineEqual(
			&left,
			&right,
			true),
	}
}

func NewDiffLineUsingString(
	index int,
	left, right *string,
) DiffLine {
	leftLine := Line{
		Number: index,
		text:   left,
	}

	rightLine := Line{
		Number: index,
		text:   right,
	}

	return DiffLine{
		LeftLine:  leftLine,
		RightLine: rightLine,
		LineEqual: NewLineEqual(
			&leftLine,
			&rightLine,
			true),
	}
}
