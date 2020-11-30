package content

type Type byte

//goland:noinspection ALL
const (
	String Type = iota
	Strings
	Map
	Bytes
	AnyItems
	BytesWithError
	AnyWrappers
)
