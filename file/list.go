package file

import (
	"io/fs"
	"os"

	"github.com/spf13/viper"
)

// List fills the DirEntries of required files
// under the specified path into fileList.
//
// Only these whose type is regular file will be
// filled into fileList.
func List(path string, fileList *[]fs.DirEntry) {
	newFileList, _ := os.ReadDir(path)

	for _, f := range newFileList {
		if f.Type().IsRegular() { // All symbol links will be ignored.
			*fileList = append(*fileList, f)
		} else if f.IsDir() && viper.GetBool("file.recursive") {
			// f.IsDir() will not follow a
			// symbol link links a directory.
			List(f.Name(), fileList)
		}
	}
}
