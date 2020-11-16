package panicmsg

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Not implemented message
// if url empty then returns constants.NotImplemented
// else returns constants.NotImplemented + " : [TODO] Will be solved at (" + url + ")"
func NonImplMsg(url string) string {
	if len(url) == 0 {
		return strconst.NotImplemented
	}

	return strconst.NotImplemented + " : [TODO] Will be solved at (" + url + ")"
}
