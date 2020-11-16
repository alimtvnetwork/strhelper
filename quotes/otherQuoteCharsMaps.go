package quotes

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

var otherQuoteCharsMaps = map[uint8]QuoteStatus{
	strconst.SingleQuoteSymbol: {
		IsQuoteFound: true,
		Found:        Single,
	},
	strconst.DoubleQuoteSymbol: {
		IsQuoteFound: true,
		Found:        Double,
	},
}
