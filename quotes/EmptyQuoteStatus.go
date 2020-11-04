package quotes

func EmptyQuoteStatus() QuoteStatus {
	return QuoteStatus{
		IsQuoteFound: false,
		Found:        nil,
	}
}
