package stringsearch

type (
	IsLineContainsFunc func(
		index int,
		contentLine string,
	) bool
)
