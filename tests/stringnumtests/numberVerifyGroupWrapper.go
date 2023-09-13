package stringnumtests

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/simplewrap"
)

type numberVerifyGroupWrapper struct {
	methodName   string
	verifierFunc func(input string) bool
	testCases    []numberVerifyWrapper
}

func (n numberVerifyGroupWrapper) MethodName() string {
	return simplewrap.WithBrackets(n.methodName) + constants.Space
}
