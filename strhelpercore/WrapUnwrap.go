package strhelpercore

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/internal/isinternal"
	"gitlab.com/evatix-go/strhelper/strconst"
)

// It is useful for generic wrap & unwrap tasks.
type WrapUnWrap struct {
	// Represents left side
	Start *string
	// Represents right side
	End *string
}

// @IsLeftFound:
//  - Indicates whether WrapUnWrap.Start is present in the left side as a start word
// @IsRightFound:
//  - Indicates whether WrapUnWrap.End is present in the right side as an ending word
type WrapStatus struct {
	// Indicates whether WrapUnWrap.Start is present in the left side as a start word
	IsLeftFound bool
	// Indicates whether WrapUnWrap.End is present in the right side as an ending word
	IsRightFound bool
}

func NewWrapUnWrap(start, end string) *WrapUnWrap {
	return &WrapUnWrap{
		Start: &start,
		End:   &end,
	}
}

func NewWrapUnWrapPtr(start, end *string) *WrapUnWrap {
	return &WrapUnWrap{
		Start: start,
		End:   end,
	}
}

func (wrapUnwrap *WrapUnWrap) IsNull() bool {
	return (wrapUnwrap.Start == nil) && (wrapUnwrap.End == nil)
}

func (wrapUnwrap *WrapUnWrap) IsEmpty() bool {
	return (wrapUnwrap.Start == nil || *wrapUnwrap.Start == "") && (wrapUnwrap.End == nil || *wrapUnwrap.End == "")
}

func (wrapUnwrap *WrapUnWrap) IsStartEmpty() bool {
	return wrapUnwrap.Start == nil || *wrapUnwrap.Start == ""
}

func (wrapUnwrap *WrapUnWrap) IsEndEmpty() bool {
	return wrapUnwrap.End == nil || *wrapUnwrap.End == ""
}

func (wrapUnwrap *WrapUnWrap) IsEquals(another *WrapUnWrap) bool {
	if another == nil {
		return false
	}

	if wrapUnwrap.IsNull() && another.IsNull() {
		return true
	}

	isStartWrapper := isinternal.PointerEqualBasedOnAddressDeduction(wrapUnwrap.Start, another.Start)

	if isStartWrapper.IsApplicableWithFalse() {
		return isStartWrapper.Result
	}

	isEndWrapper := isinternal.PointerEqualBasedOnAddressDeduction(wrapUnwrap.End, another.End)

	if isEndWrapper.IsApplicableWithFalse() {
		return isEndWrapper.Result
	}

	return isinternal.EqualsPtr(wrapUnwrap.Start, another.Start) &&
		isinternal.EqualsPtr(wrapUnwrap.End, another.End)
}

// Wraps the given input string with WrapUnWrap.Start & WrapUnWrap.End without checking anything.
func (wrapUnwrap *WrapUnWrap) Wrap(input string) string {
	wrapped := *wrapUnwrap.Start + input + *wrapUnwrap.End

	return wrapped
}

// Wraps the given input string with WrapUnWrap.Start & WrapUnWrap.End without checking anything.
func (wrapUnwrap *WrapUnWrap) WrapPtr(input *string) *string {
	wrapped := *wrapUnwrap.Start + *input + *wrapUnwrap.End

	return &wrapped
}

// Wraps the given input string with WrapUnWrap.Start & WrapUnWrap.End checking.
//
// Conditions:
//  - if input start has WrapUnWrap.Start then skips adding for WrapUnWrap.Start.
//  - if input end has WrapUnWrap.End then skips adding for WrapUnWrap.End.
//  - if both found then returns as is.
func (wrapUnwrap *WrapUnWrap) WrapByChecking(input string, isCaseSensitive bool) string {
	return *wrapUnwrap.WrapByCheckingPtr(&input, isCaseSensitive)
}

// Wraps the given input string with WrapUnWrap.Start & WrapUnWrap.End checking.
//
// Conditions:
//  - if input start has WrapUnWrap.Start then skips adding for WrapUnWrap.Start.
//  - if input end has WrapUnWrap.End then skips adding for WrapUnWrap.End.
//  - if both found then returns as is.
func (wrapUnwrap *WrapUnWrap) WrapByCheckingPtr(input *string, isCaseSensitive bool) *string {
	status := wrapUnwrap.WrapStatus(input, isCaseSensitive)
	isBothFound := status.IsLeftFound && status.IsRightFound

	if !isBothFound {
		return wrapUnwrap.WrapPtr(input)
	}

	if status.IsLeftFound {
		wrapped := *input + *wrapUnwrap.End

		return &wrapped
	}

	if status.IsRightFound {
		wrapped := *wrapUnwrap.Start + *input

		return &wrapped
	}

	// both found.
	return input
}

// returns true if input has WrapUnWrap.Start at the beginning and end contains end of WrapUnWrap.End
func (wrapUnwrap *WrapUnWrap) HasStartEndBoth(input *string, isCaseSensitive bool) bool {
	status := wrapUnwrap.WrapStatus(input, isCaseSensitive)

	return status.IsRightFound && status.IsLeftFound
}

// returns true if input has WrapUnWrap.Start at the beginning
func (wrapUnwrap *WrapUnWrap) HasStartAtBeginning(input *string, isCaseSensitive bool) bool {
	status := wrapUnwrap.WrapStatus(input, isCaseSensitive)

	return status.IsLeftFound
}

// returns true if input has WrapUnWrap.End at the end
func (wrapUnwrap *WrapUnWrap) HasEndAtEnding(input *string, isCaseSensitive bool) bool {
	status := wrapUnwrap.WrapStatus(input, isCaseSensitive)

	return status.IsLeftFound
}

// Checks first if the start exists on the left if so then unwraps it.
//
// Checks if the end exists on the right if so then unwraps it.
//
// If no wrap found then returns input as is.
func (wrapUnwrap *WrapUnWrap) Unwrap(input string, isCaseSensitive bool) string {
	return wrapUnwrap.UnwrapPtr(&input, isCaseSensitive)
}

// Checks first if the start exists on the left if so then unwraps it.
//
// Checks if the end exists on the right if so then unwraps it.
//
// If no wrap found then returns input as is.
func (wrapUnwrap *WrapUnWrap) UnwrapPtr(input *string, isCaseSensitive bool) string {
	status := wrapUnwrap.WrapStatus(input, isCaseSensitive)
	isBothFound := status.IsLeftFound && status.IsRightFound

	if isBothFound {
		return (*input)[len(*wrapUnwrap.Start)-1 : len(*input)-len(*wrapUnwrap.End)-1]
	}

	if status.IsLeftFound {
		return (*input)[len(*wrapUnwrap.Start)-1:]
	}

	if status.IsRightFound {
		return (*input)[:len(*input)-len(*wrapUnwrap.End)-1]
	}

	return *input
}

// Returns WrapStatus contains found information of the wrapper start and end.
//
// Where IsLeftFound if found as a starting word.
// Where IsRightFound if found as an ending word.
func (wrapUnwrap *WrapUnWrap) WrapStatus(input *string, isCaseSensitive bool) *WrapStatus {
	if input == nil || *input == strconst.EmptyString {
		return &WrapStatus{
			IsLeftFound:  false,
			IsRightFound: false,
		}
	}

	leftWord := (*input)[0 : len(*wrapUnwrap.Start)-1]
	rightWord := (*input)[len(*input)-len(*wrapUnwrap.End)-1:]
	comparingLeft := *wrapUnwrap.Start
	comparingRight := *wrapUnwrap.End

	if isCaseSensitive == false {
		leftWord = strings.ToLower(leftWord)
		rightWord = strings.ToLower(rightWord)
		comparingLeft = strings.ToLower(comparingLeft)
		comparingRight = strings.ToLower(comparingRight)
	}

	return &WrapStatus{
		IsLeftFound:  leftWord == comparingLeft,
		IsRightFound: rightWord == comparingRight,
	}
}
