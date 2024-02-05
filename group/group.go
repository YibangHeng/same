package group

import (
	"github.com/yibangheng/same/file"
)

type Grouper interface {
	// Group groups elements in slice by keys.
	//
	// The type of elements must be regular file
	// but the type of keys is implementation
	// dependent.
	//
	// Group should return nil only if the slice
	// is nil or has no element.
	Group([]file.EntryInfoType) map[file.Any][]file.EntryInfoType
}
