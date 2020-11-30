package parsingtype

type ByteType uint8

var (
	byteTypesStrings = []string{
		"Unknown", "Unsafe", "Encoding", "JsonParsing", "AnyToValueStringBytes", "AnyToFullStringBytes",
	}
)

const (
	Unknown ByteType = iota
	Unsafe
	Encoding
	JsonParsing
	AnyToValueStringBytes
	AnyToFullStringBytes
)

func (byteType ByteType) String() string {
	return byteTypesStrings[byteType]
}
