package file

import (
	"io/fs"
	"sort"
)

// Sort sorts fileList by file size in ascending
// order.
func Sort(fileList *[]fs.DirEntry) {
	sort.SliceStable(*fileList, func(i, j int) bool {
		iInfo, _ := (*fileList)[i].Info()
		jInfo, _ := (*fileList)[j].Info()

		return iInfo.Size() < jInfo.Size()
	})
}
