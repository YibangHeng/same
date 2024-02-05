package file

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// List fills the DirEntries of required files
// under the specified path into fileList.
//
// Only these whose type is regular file will be
// filled into fileList.
func List(fileList *[]EntryInfo, paths ...string) {
	for _, path := range paths {
		newFileList, _ := os.ReadDir(path)

		for _, f := range newFileList {
			if f.Type().IsRegular() { // All non-regular files will be ignored.
				if IsHiddenFile(f.Name()) || !viper.GetBool("file.ignore-hidden-file") {
					*fileList = append(*fileList, EntryInfo{path, f})
				}
			} else if f.IsDir() && viper.GetBool("file.recursive") {
				// f.IsDir() will not follow a
				// symbol link links a directory.
				List(fileList, filepath.Join(path, f.Name()))
			}
		}
	}
}
