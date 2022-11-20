package group

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type MD5Grouper struct{}

func (_ *MD5Grouper) md5(t Type) string {
	fd, err := os.Open(filepath.Join(viper.GetString("file.directory"), t.Name()))
	if err != nil {
		return ""
	}
	defer fd.Close()

	md5Writer := md5.New()
	if _, err := io.Copy(md5Writer, fd); err != nil {
		return ""
	}

	return hex.EncodeToString(md5Writer.Sum(nil))
}

// Group groups elements by its' md5 value. The
// keys will be strings stands for the
// hexadecimal encoded md5 values.
//
// s should not contain any non-regular files.
// Any files cannot read will be considered
// empty files.
func (mg *MD5Grouper) Group(s []Type) (m map[Any][]Type) {
	if len(s) == 0 {
		return nil
	}

	m = make(map[Any][]Type)

	for _, t := range s {
		appendToMap(m, mg.md5(t), t)
	}

	return m
}
