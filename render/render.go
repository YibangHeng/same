package render

import (
	"github.com/spf13/viper"
	"github.com/yibangheng/same/file"
)

// Render renders groupedFiles in specified
// format and print it on stdout.
func Render(groupedFiles map[interface{}][]file.EntryInfoType) {
	if viper.GetBool("format.json") {
		JSON(groupedFiles)
	} else {
		Table(groupedFiles, "MD5")
	}
}
