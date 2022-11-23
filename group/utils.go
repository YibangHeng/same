package group

import (
	"sort"
	"sync"
)

// muAppendToMap protects map in appendToMap.
var muAppendToMap sync.Mutex

// appendToMap appends de to slice associated
// with key in m.
//
// If m is nil, appendToMap do nothing.
func appendToMap(m map[Any][]Type, key Any, t Type) {
	if m == nil {
		return
	}

	muAppendToMap.Lock()
	defer muAppendToMap.Unlock()

	s, ok := m[key]
	if !ok {
		s = make([]Type, 0)
	}

	s = append(s, t)
	m[key] = s // If a new underlying array is allocated.
}

// muDedup protects map in dedup.
var muDedup sync.Mutex

// dedup removes pairs whose value contains one
// or less elements and return itself.
//
// If m is nil, dedup do nothing.
func dedup(m map[Any][]Type) map[Any][]Type {
	if m == nil {
		return nil
	}

	muDedup.Lock()
	defer muDedup.Unlock()

	for k, v := range m {
		if len(v) < 2 {
			delete(m, k)
		}
	}

	return m
}

// sortBySize sorts fileList by file size in descending
// order.
func sortBySize(fileList *[]Type) {
	sort.SliceStable(*fileList, func(i, j int) bool {
		return (*fileList)[i].GetSize() > (*fileList)[j].GetSize()
	})
}
