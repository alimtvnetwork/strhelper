package splits

import "gitlab.com/evatix-go/core/constants"

var (
	bothSlashesMap = &map[rune]bool{
		constants.ForwardRune : true,
		constants.BackwardRune : true,
	}
)
