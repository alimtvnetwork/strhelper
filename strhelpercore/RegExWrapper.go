package strhelpercore

import (
	"regexp"
)

type RegExWrapper struct {
	index              int
	request            *string
	regex              *regexp.Regexp
	errorWrapper       *ErrorWrapper
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

func (regExWrapper *RegExWrapper) Request() *string {
	return regExWrapper.request
}

// Requires to compile regex to get the currentError
func (regExWrapper *RegExWrapper) ErrorWrapper() *ErrorWrapper {
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
		regExWrapper.errorWrapper = NewErrorWrapper(er)
	}
}

func (regExWrapper *RegExWrapper) Value() *regexp.Regexp {
	return regExWrapper.RegEx()
}
