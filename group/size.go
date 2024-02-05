package group

import (
	"github.com/spf13/viper"
	"github.com/yibangheng/same/file"
)

var _ Grouper = (*SizeGrouper)(nil)

type SizeGrouper struct{}

// Group groups elements by file size. The keys
// will be int64s stands for length of files in
// bytes.
//
// s should not contain any non-regular files.
// Any files whose info cannot be retrieved will
// be considered 0 byte.
func (_ *SizeGrouper) Group(s []file.EntryInfoType) map[file.Any][]file.EntryInfoType {
	if len(s) == 0 {
		return nil
	}

	m := make(map[file.Any][]file.EntryInfoType)

	for _, t := range s {
		if t.GetSize() != 0 || !viper.GetBool("file.ignore-empty-file") {
			appendToMap(m, t.GetSize(), t)
		}
	}

	return dedup(m)
}
