package quotes

type Quote uint8

const (
	UnknownQuote Quote = iota
	Double       Quote = '"'
	Single       Quote = '\''
)

func (quote Quote) IsEqual(char uint8) bool {
	return quote.Value() == char
}

func (quote Quote) String() string {
	return string(quote)
}

func (quote Quote) Value() uint8 {
	return uint8(quote)
}

func (quote Quote) ValuePtr() *uint8 {
	val := uint8(quote)

	return &val
}

func (quote Quote) Wrap(str *string) string {
	if isEmptyStringPtr(str) {
		return quote.String() + quote.String()
	}

	return quote.String() +
		*str +
		quote.String()
}

func (quote Quote) GetTheOther() Quote {
	other, _ := otherQuoteMaps[quote]

	return other
}
