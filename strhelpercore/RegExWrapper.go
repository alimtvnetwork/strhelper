package strhelpercore

import (
	"regexp"

	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/core/regexnew"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

type RegExWrapper struct {
	index              int
	request            string
	regex              *regexp.Regexp
	errorWrapper       *errorwrapper.Wrapper
	regExResultWrapper *RegExResultWrapper
}

func NewRegExWrapper(index int, request string) *RegExWrapper {
	return &RegExWrapper{
		index:   index,
		request: request,
	}
}

// RegExResultWrapper Requires to compile regex
func (it *RegExWrapper) RegExResultWrapper(content string) *RegExResultWrapper {
	if it.regExResultWrapper == nil && it.ErrorWrapper().IsEmpty() {
		it.regExResultWrapper = NewRegExResultWrapper(it.index, content, it.regex)
	}

	return it.regExResultWrapper
}

func (it *RegExWrapper) Request() string {
	return it.request
}

func (it *RegExWrapper) IsEquals(another *RegExWrapper) bool {
	if another == nil && it == nil {
		return true
	}

	if another == nil || it == nil {
		return false
	}

	if another == it {
		return true
	}

	if another == nil {
		return false
	}

	if another == it {
		return true
	}

	return it.request == another.request
}

func (it *RegExWrapper) IsEqualsString(str string) bool {
	return it.request == str
}

// ErrorWrapper
//
// Requires to compile regex to get the currentError
func (it *RegExWrapper) ErrorWrapper() *errorwrapper.Wrapper {
	it.initializeRegex()

	return it.errorWrapper
}

// RegEx
//
// Compile request and return the regExWrapper.regex and cache currentError and regex to return next time.
func (it *RegExWrapper) RegEx() *regexp.Regexp {
	it.initializeRegex()

	return it.regex
}

func (it *RegExWrapper) initializeRegex() {
	if it.regex == nil && it.errorWrapper == nil {
		regexCompiled, err := regexnew.NewLock(it.request)
		it.regex = regexCompiled
		it.errorWrapper = errnew.UsingStackSkip(
			codestack.Skip1,
			errtype.RegexCompiledFailed,
			err)
	}
}

func (it *RegExWrapper) Value() *regexp.Regexp {
	return it.RegEx()
}
