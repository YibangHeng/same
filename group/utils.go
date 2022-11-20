package group

// appendToMap appends de to slice associated
// with key in m.
//
// If m is nil, appendToMap do nothing.
func appendToMap(m map[Any][]Type, key Any, t Type) {
	if m == nil {
		return
	}

	s, ok := m[key]
	if !ok {
		s = make([]Type, 0)
	}

	s = append(s, t)
	m[key] = s // If a new underlying array is allocated.
}
