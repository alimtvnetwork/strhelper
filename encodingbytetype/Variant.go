package encodingbytetype

type Variant uint8

var (
	byteTypesStrings = []string{
		"Unknown", "Unsafe", "Encoding", "JsonParsing", "AnyToValueStringBytes", "AnyToFullStringBytes",
	}
)

const (
	Unknown Variant = iota
	Unsafe
	Encoding
	JsonParsing
	AnyToValueStringBytes
	AnyToFullStringBytes
)

func (byteType Variant) String() string {
	return byteTypesStrings[byteType]
}
