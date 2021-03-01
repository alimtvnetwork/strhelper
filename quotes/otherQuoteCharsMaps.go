package quotes

import (
	"gitlab.com/evatix-go/core/constants"
)

var otherQuoteCharsMaps = map[uint8]QuoteStatus{
	constants.SingleQuoteSymbol: {
		IsQuoteFound: true,
		Found:        Single,
	},
	constants.DoubleQuoteSymbol: {
		IsQuoteFound: true,
		Found:        Double,
	},
}
