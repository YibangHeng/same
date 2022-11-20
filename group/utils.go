package group

import (
	"sort"
	"sync"
)

// mu protects map in appendToMap.
var mu sync.Mutex

// appendToMap appends de to slice associated
// with key in m.
//
// If m is nil, appendToMap do nothing.
func appendToMap(m map[Any][]Type, key Any, t Type) {
	if m == nil {
		return
	}

	mu.Lock()
	defer mu.Unlock()

	s, ok := m[key]
	if !ok {
		s = make([]Type, 0)
	}

	s = append(s, t)
	m[key] = s // If a new underlying array is allocated.
}

// sortBySize sorts fileList by file size in descending
// order.
func sortBySize(fileList *[]Type) {
	sort.SliceStable(*fileList, func(i, j int) bool {
		iInfo, _ := (*fileList)[i].Info()
		jInfo, _ := (*fileList)[j].Info()

		return iInfo.Size() > jInfo.Size()
	})
}
