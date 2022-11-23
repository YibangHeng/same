package group

import (
	"github.com/YibangHeng/same/file"
)

type Any = interface{}
type Type = file.FileInfo

type Grouper interface {
	// Group groups elements in slice by keys.
	//
	// The type of elements must be regular file
	// but the type of keys is implementation
	// dependent.
	//
	// Group should return nil only if the slice
	// is nil or has no element.
	Group([]Type) map[Any][]Type
}

// Belows are implementations of Grouper:
var _ Grouper = new(SizeGrouper)
var _ Grouper = new(MD5Grouper)
