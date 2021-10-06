package strhelpercore

import (
	"sync"
)

// IndexesResultSet
//
// Search term with indexes where found.
type IndexesResultSet struct {
	// Key : Search term
	//
	// Values : Indexes where the search terms are found.
	StringKeyAsIndexesMap map[string][]int
	// represents the last index where search key is possible
	LastIndexFound  int
	indexAsKeysMap  map[int]string
	hasResult       bool
	isEmptySet      bool
	totalFound      *int
	indexesKeyMutex sync.Mutex
	sync.Mutex
}

func (indexesResultSet *IndexesResultSet) GetIndexesMapWhereIndexAsKeyWithLock() map[int]string {
	indexesResultSet.indexesKeyMutex.Lock()
	defer indexesResultSet.indexesKeyMutex.Unlock()

	return indexesResultSet.GetIndexesMapWhereIndexAsKey()
}

func (indexesResultSet *IndexesResultSet) GetIndexesMapWhereIndexAsKey() map[int]string {
	if indexesResultSet.hasResult && indexesResultSet.indexAsKeysMap == nil {
		indexAsKeyMap := make(map[int]string, indexesResultSet.CountOfAllFoundIndexes())
		length := 0
		for key, indexes := range indexesResultSet.StringKeyAsIndexesMap {
			length = len(indexes)

			if length > 0 {
				for _, valueIndex := range indexes {
					indexAsKeyMap[valueIndex] = key
				}
			}
		}

		indexesResultSet.indexAsKeysMap = indexAsKeyMap
	}

	return indexesResultSet.indexAsKeysMap
}

func (indexesResultSet *IndexesResultSet) CountOfAllFoundIndexesWithLock() int {
	indexesResultSet.Mutex.Lock()
	defer indexesResultSet.Mutex.Unlock()

	return indexesResultSet.CountOfAllFoundIndexes()
}

func (indexesResultSet *IndexesResultSet) CountOfAllFoundIndexes() int {
	length := 0
	if indexesResultSet.totalFound == nil {
		counter := 0
		for _, values := range indexesResultSet.StringKeyAsIndexesMap {
			length = len(values)
			if length > 0 {
				counter += length
			}
		}

		indexesResultSet.totalFound = &counter
	}

	return *indexesResultSet.totalFound
}

func (indexesResultSet *IndexesResultSet) HasResultBy(key string) bool {
	results, has := (indexesResultSet.StringKeyAsIndexesMap)[key]

	if has && len(results) > 0 {
		return true
	}

	return false
}

// GetIndexes Returns indexes if exists or else returns nil.
func (indexesResultSet *IndexesResultSet) GetIndexes(key string) []int {
	results, has := (indexesResultSet.StringKeyAsIndexesMap)[key]

	if has && len(results) > 0 {
		return results
	}

	return nil
}

func (indexesResultSet *IndexesResultSet) SetHasResult(hasResult bool) {
	indexesResultSet.hasResult = hasResult
}

func (indexesResultSet *IndexesResultSet) HasResult() bool {
	return indexesResultSet.hasResult
}

func (indexesResultSet *IndexesResultSet) SetEmpty(isEmpty bool) {
	indexesResultSet.isEmptySet = isEmpty
}

func (indexesResultSet *IndexesResultSet) HasResultWithLock() bool {
	indexesResultSet.Lock()
	defer indexesResultSet.Unlock()

	return indexesResultSet.hasResult
}

func (indexesResultSet *IndexesResultSet) IsEmpty() bool {
	return indexesResultSet.isEmptySet
}

func NewIndexesResultSet(
	indexesMap map[string][]int,
	hasFoundAny bool,
	maxIndexWhereFound int,
) *IndexesResultSet {
	return &IndexesResultSet{
		StringKeyAsIndexesMap: indexesMap,
		LastIndexFound:        maxIndexWhereFound,
		hasResult:             hasFoundAny,
		isEmptySet:            !hasFoundAny,
	}
}

func NewEmptyIndexesResultSet() *IndexesResultSet {
	return &IndexesResultSet{
		hasResult:  false,
		isEmptySet: true,
	}
}
