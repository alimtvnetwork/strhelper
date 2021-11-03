package stringnum

var (
	IsAscii = &isAsc{
		Integer: &ascInteger{},
		Number:  &ascNumber{},
	}

	IsUnicode = &isUnicode{
		Integer: &unicodeInteger{},
		Number:  &unicodeNumber{},
	}
)
