package strcompare

type CreateType byte

const (
	PlainText CreateType = 0
	Lines     CreateType = 1
)

func (createType CreateType) IsPlainText() bool {
	return createType == PlainText
}

func (createType CreateType) IsLines() bool {
	return createType == Lines
}
