package strcompare

import (
	"sync"

	"gitlab.com/evatix-go/strhelper/internal/misc"
	"gitlab.com/evatix-go/strhelper/strswasync"
)

type DiffLines struct {
	Diffs      []DiffLine
	Length     int
	linesEqual *LineEqual
	sync.Mutex
}

func NewDiffLines(leftLines, rightLines *strswasync.Wrapper) *DiffLines {
	leftLength := leftLines.Length()
	rightLength := rightLines.Length()
	maxLength := misc.MaxInt(leftLength, rightLength)
	diffLines := make([]DiffLine, maxLength)
	equalityChecker := LineEqual{
		IsSame:                   true,
		IsSameIgnoringWhitespace: true,
	}

	for i := 0; i < maxLength; i++ {
		left := leftLines.GetSafeIndexAt(i, nil)
		right := rightLines.GetSafeIndexAt(i, nil)
		diffLine := NewDiffLineUsingString(i, left, right)
		diffLines[i] = diffLine

		if !diffLine.LineEqual.IsSameIgnoringWhitespace {
			equalityChecker = diffLine.LineEqual
		}

		if !diffLine.LineEqual.IsSame {
			equalityChecker.IsSame = false
		}
	}

	return &DiffLines{
		Diffs:      diffLines,
		Length:     maxLength,
		linesEqual: &equalityChecker,
	}
}

func NewDiffLinesUsingRawLines(leftLines, rightLines *[]string) *DiffLines {
	leftLength := misc.LinesLength(leftLines)
	rightLength := misc.LinesLength(rightLines)
	maxLength := misc.MaxInt(leftLength, rightLength)
	diffLines := make([]DiffLine, maxLength)
	equalityChecker := LineEqual{
		IsSame:                   true,
		IsSameIgnoringWhitespace: true,
	}

	leftLastIndex := leftLength - 1
	rightLengthIndex := rightLength - 1

	for i := 0; i < maxLength; i++ {
		left := misc.LinesSafeValuePtrAt(leftLines, leftLastIndex, i, nil)
		right := misc.LinesSafeValuePtrAt(rightLines, rightLengthIndex, i, nil)
		diffLine := NewDiffLineUsingString(i, left, right)
		diffLines[i] = diffLine

		if !diffLine.LineEqual.IsSameIgnoringWhitespace {
			equalityChecker = diffLine.LineEqual
		}

		if !diffLine.LineEqual.IsSame {
			equalityChecker.IsSame = false
		}
	}

	return &DiffLines{
		Diffs:      diffLines,
		Length:     maxLength,
		linesEqual: &equalityChecker,
	}
}

func (diffLines *DiffLines) LinesEqualWithLock() LineEqual {
	diffLines.Lock()
	defer diffLines.Unlock()

	return diffLines.LinesEqual()
}

func (diffLines *DiffLines) LinesEqual() LineEqual {
	if diffLines.linesEqual != nil {
		return *diffLines.linesEqual
	}

	// TODO fix logic, double check the logic in future.
	var isSame = true
	var diffLine *DiffLine
	for i := 0; i <= diffLines.Length; i++ {
		diffLine = &(diffLines.Diffs)[i]
		lineEqual := diffLine.LineEqual
		if !lineEqual.IsSameIgnoringWhitespace {

			diffLines.linesEqual = &lineEqual

			break
		}

		if !lineEqual.IsSame {
			isSame = false
		}
	}

	if diffLines.linesEqual == nil {
		lineEqual := LineEqual{
			IsSame:                   isSame,
			IsSameIgnoringWhitespace: true,
		}

		diffLines.linesEqual = &lineEqual
	}

	return *diffLines.linesEqual
}
