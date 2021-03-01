package strhelpercore

import (
	"regexp"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"

	"gitlab.com/evatix-go/strhelper/internal/isinternal"
)

type RegExWrapper struct {
	index              int
	request            *string
	regex              *regexp.Regexp
	errorWrapper       *errorwrapper.Wrapper
	regExResultWrapper *RegExResultWrapper
}

func NewRegExWrapper(index int, request *string) *RegExWrapper {
	return &RegExWrapper{
		index:   index,
		request: request,
	}
}

// Requires to compile regex
func (regExWrapper *RegExWrapper) RegExResultWrapper(content *string) *RegExResultWrapper {
	if regExWrapper.regExResultWrapper == nil && regExWrapper.ErrorWrapper().IsEmpty() {
		regExWrapper.regExResultWrapper = NewRegExResultWrapper(regExWrapper.index, content, regExWrapper.regex)
	}

	return regExWrapper.regExResultWrapper
}

func (regExWrapper *RegExWrapper) Request() string {
	return *regExWrapper.request
}

func (regExWrapper *RegExWrapper) IsEquals(another *RegExWrapper) bool {
	if another == nil {
		return false
	}

	if another == regExWrapper {
		return true
	}

	return isinternal.EqualsPtr(regExWrapper.request, another.request)
}

func (regExWrapper *RegExWrapper) IsEqualsString(str *string) bool {
	return isinternal.EqualsPtr(regExWrapper.request, str)
}

// Requires to compile regex to get the currentError
func (regExWrapper *RegExWrapper) ErrorWrapper() *errorwrapper.Wrapper {
	regExWrapper.initializeRegex()

	return regExWrapper.errorWrapper
}

// Compile request and return the regExWrapper.regex and cache currentError and regex to return next time.
func (regExWrapper *RegExWrapper) RegEx() *regexp.Regexp {
	regExWrapper.initializeRegex()

	return regExWrapper.regex
}

func (regExWrapper *RegExWrapper) initializeRegex() {
	if regExWrapper.regex == nil && regExWrapper.errorWrapper == nil {
		r, er := regexp.Compile(*regExWrapper.request)
		regExWrapper.regex = r
		errorWrapper := errnew.Err(er)
		regExWrapper.errorWrapper = &errorWrapper
	}
}

func (regExWrapper *RegExWrapper) Value() *regexp.Regexp {
	return regExWrapper.RegEx()
}
