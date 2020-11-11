package strhelpercore

type SplitResult struct {
	SplitPrev *string
	// or the splitter
	Separator *string
	Index     int
	IsEmpty   bool
}
