package group

import (
	"sync"

	"github.com/yibangheng/same/file"
)

// muAppendToMap protects map in appendToMap.
var muAppendToMap sync.Mutex

// appendToMap appends de to slice associated
// with key in m.
//
// If m is nil, appendToMap do nothing.
func appendToMap(m map[file.Any][]file.EntryInfoType, key file.Any, t file.EntryInfoType) {
	if m == nil {
		return
	}

	muAppendToMap.Lock()
	defer muAppendToMap.Unlock()

	s, ok := m[key]
	if !ok {
		s = make([]file.EntryInfoType, 0)
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
func dedup(m map[file.Any][]file.EntryInfoType) map[file.Any][]file.EntryInfoType {
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
