package file

import (
	"io/fs"
	"path/filepath"
)

// EntryInfo contains the info of file, such as
// file type, full path, etc.
type EntryInfo struct {
	rootFolder string
	dirEntry   fs.DirEntry
}

func (fi *EntryInfo) GetName() string {
	return fi.dirEntry.Name()
}

func (fi *EntryInfo) GetRelativeName() string {
	return filepath.Join(fi.rootFolder, fi.dirEntry.Name())
}

func (fi *EntryInfo) GetFullName() string {
	return filepath.Join(fi.rootFolder, fi.dirEntry.Name())
}

func (fi *EntryInfo) GetSize() int64 {
	i, _ := fi.dirEntry.Info()
	return i.Size()
}

type Any = interface{}
type EntryInfoType = EntryInfo
