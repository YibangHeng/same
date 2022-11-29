package render

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/yibangheng/same/file"
)

// muAppendToMap protects map in appendToMap.
var muAppendToMap sync.Mutex

// appendToMap appends de to slice associated
// with key in m.
//
// If m is nil, appendToMap do nothing.
func appendToMap(m map[string][]string, key string, t string) {
	if m == nil {
		return
	}

	muAppendToMap.Lock()
	defer muAppendToMap.Unlock()

	s, ok := m[key]
	if !ok {
		s = make([]string, 0)
	}

	s = append(s, t)
	m[key] = s // If a new underlying array is allocated.
}

func JSON(m map[file.Any][]file.EntryInfoType) {
	// Convert map[Any][]Type to
	// map[string][]string since type
	// map[Any][]Type is unsupported in
	// json.Marshal.
	mss := make(map[string][]string, len(m))
	for k, v := range m {
		for _, vv := range v {
			appendToMap(mss, fmt.Sprintf("%v", k), vv.GetRelativeName())
		}
	}

	b, err := json.Marshal(mss)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", b)
}
