package panicmsg

import (
	"gitlab.com/auk-go/core/constants"
)

// Not implemented message
// if url empty then returns constants.NotImplemented
// else returns constants.NotImplemented + " : [TODO] Will be solved at (" + url + ")"
func NonImplMsg(url string) string {
	if len(url) == 0 {
		return constants.NotImplemented
	}

	return constants.NotImplemented + " : [TODO] Will be solved at (" + url + ")"
}
