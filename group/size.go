package group

type SizeGrouper struct{}

// Group groups elements by file size. The keys
// will be int64s stands for length of files in
// bytes.
//
// s should not contain any non-regular files.
// Any files whose info cannot be retrieved will
// be considered 0 byte.
func (_ *SizeGrouper) Group(s []Type) (m map[Any][]Type) {
	if len(s) == 0 {
		return nil
	}

	m = make(map[Any][]Type)

	for _, t := range s {
		deInfo, _ := t.Info()
		appendToMap(m, deInfo.Size(), t)
	}

	return m
}
