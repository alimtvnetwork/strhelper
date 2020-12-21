package strswasync

import (
	"sync"

	"gitlab.com/evatix-go/strhelper/internal/misc"
)

func NewPtr(lines *[]string) *Wrapper {
	return &Wrapper{
		lines:         lines,
		linesWrappers: nil,
		lowerLines:    nil,
		upperLines:    nil,
		Mutex:         sync.Mutex{},
		length:        misc.LinesLength(lines),
	}
}

func New(lines []string) *Wrapper {
	if lines == nil {
		return &Wrapper{
			lines:         nil,
			linesWrappers: nil,
			lowerLines:    nil,
			upperLines:    nil,
			Mutex:         sync.Mutex{},
			length:        0,
		}
	}

	return &Wrapper{
		lines:         &lines,
		linesWrappers: nil,
		lowerLines:    nil,
		upperLines:    nil,
		Mutex:         sync.Mutex{},
		length:        len(lines),
	}
}
