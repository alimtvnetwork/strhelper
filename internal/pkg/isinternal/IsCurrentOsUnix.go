package isinternal

import "gitlab.com/evatix-go/strhelper/strconst"

func IsCurrentOsUnix() bool {
	//goland:noinspection ALL
	return strconst.NewLine == strconst.NewLineUnix
}
