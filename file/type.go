package file

import (
	"io/fs"
	"path/filepath"
)

// FileInfo contains the info of file, such as
// file type, full path, etc.
type FileInfo struct {
	rootFolder string
	dirEntry   fs.DirEntry
}

func (fi *FileInfo) GetName() string {
	return fi.dirEntry.Name()
}

func (fi *FileInfo) GetRelativeName() string {
	return filepath.Join(fi.rootFolder, fi.dirEntry.Name())
}

func (fi *FileInfo) GetFullName() string {
	return filepath.Join(fi.rootFolder, fi.dirEntry.Name())
}

func (fi *FileInfo) GetSize() int64 {
	i, _ := fi.dirEntry.Info()
	return i.Size()
}
