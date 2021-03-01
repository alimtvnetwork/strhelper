package strcompare

import (
	"strings"
	"sync"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/strswasync"
)

type Result struct {
	equalityChecker *LineEqual
	diffLines       *DiffLines
	leftText        *string
	rightText       *string
	leftLines       *strswasync.Wrapper
	rightLines      *strswasync.Wrapper
	createType      CreateType
	sync.Mutex
}

func (result *Result) DiffLinesWithLock() *DiffLines {
	result.Lock()
	defer result.Unlock()

	return result.DiffLines()
}

func (result *Result) DiffLines() *DiffLines {
	if result.diffLines == nil {
		result.diffLines = NewDiffLines(result.LeftLines(), result.RightLines())
	}

	return result.diffLines
}

func (result *Result) EqualityCheckerWithLock() *LineEqual {
	result.Lock()
	defer result.Unlock()

	return result.EqualityChecker()
}

func (result *Result) EqualityChecker() *LineEqual {
	if result.equalityChecker != nil {
		return result.equalityChecker
	}

	if result.createType.IsPlainText() {
		return result.plainTextEqualityChecker()
	}

	if result.createType.IsLines() {
		return result.linesEqualityChecker()
	}

	panic("Equality check doesn't support other than plain text and lines.")
}

// returns confirm strswasync.Wrapper
func (result *Result) LeftLines() *strswasync.Wrapper {
	if result.leftLines != nil {
		return result.leftLines
	}

	if result.createType.IsPlainText() &&
		result.leftText != nil &&
		result.leftLines == nil {
		lines := strings.Split(*result.leftText, newLine)
		result.leftLines = strswasync.NewPtr(&lines)
	}

	if result.leftLines == nil {
		result.leftLines = strswasync.NewPtr(nil)
	}

	return result.leftLines
}

// returns confirm strswasync.Wrapper
func (result *Result) RightLines() *strswasync.Wrapper {
	if result.rightLines != nil {
		return result.rightLines
	}

	if result.createType.IsPlainText() &&
		result.rightText != nil &&
		result.rightLines == nil {
		lines := strings.Split(*result.rightText, newLine)
		result.rightLines = strswasync.NewPtr(&lines)
	}

	if result.rightLines == nil {
		result.rightLines = strswasync.NewPtr(nil)
	}

	return result.rightLines
}

func (result *Result) LeftLinesWithLock() *strswasync.Wrapper {
	result.Lock()
	defer result.Unlock()

	return result.LeftLines()
}

func (result *Result) RightLinesWithLock() *strswasync.Wrapper {
	result.Lock()
	defer result.Unlock()

	return result.RightLines()
}

// returns confirm string ptr
func (result *Result) LeftText() *string {
	if result.leftText != nil {
		return result.leftText
	}

	if result.createType.IsLines() &&
		result.leftLines != nil &&
		result.leftText == nil {
		result.leftText = result.
			leftLines.
			AsString(newLine, false)
	}

	if result.leftText == nil {
		result.leftText = constants.EmptyStringPtr
	}

	return result.leftText
}

// returns confirm string ptr
func (result *Result) RightText() *string {
	if result.rightText != nil {
		return result.rightText
	}

	if result.createType.IsLines() &&
		result.rightLines != nil &&
		result.rightText == nil {
		result.rightText = result.
			rightLines.
			AsString(newLine, false)
	}

	if result.rightText == nil {
		result.rightText = constants.EmptyStringPtr
	}

	return result.rightText
}

func (result *Result) LeftTextWithLock() *string {
	result.Lock()
	defer result.Unlock()

	return result.LeftText()
}

func (result *Result) RightTextWithLock() *string {
	result.Lock()
	defer result.Unlock()

	return result.RightText()
}

func (result *Result) plainTextEqualityChecker() *LineEqual {
	if result.equalityChecker != nil {
		return result.equalityChecker
	}

	leftLine := Line{
		Number: 0,
		text:   result.leftText,
	}

	rightLine := Line{
		Number: 0,
		text:   result.rightText,
	}

	lineEqual := NewLineEqual(
		&leftLine,
		&rightLine,
		true)

	result.equalityChecker = &lineEqual

	return result.equalityChecker
}

func (result *Result) linesEqualityChecker() *LineEqual {
	if result.equalityChecker == nil {
		linesEqual := result.
			DiffLines().
			LinesEqual()

		result.equalityChecker = &linesEqual
	}

	return result.equalityChecker
}
