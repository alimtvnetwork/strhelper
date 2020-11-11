package quotes

import "gitlab.com/evatix-go/strhelper/constants"

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
