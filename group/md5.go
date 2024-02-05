package group

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"runtime"
	"sync"

	"github.com/yibangheng/same/file"
)

var _ Grouper = (*MD5Grouper)(nil)

type MD5Grouper struct {
	fileIn chan file.EntryInfoType
	wg     sync.WaitGroup
}

func (mg *MD5Grouper) md5(m map[file.Any][]file.EntryInfoType) {
	md5Writer := md5.New()
	for t := range mg.fileIn {
		md5Writer.Reset()

		fd, err := os.Open(t.GetFullName())
		if err == nil {
			if _, err := io.Copy(md5Writer, fd); err == nil {
				appendToMap(m, hex.EncodeToString(md5Writer.Sum(nil)), t)
			}
			fd.Close()
		}
	}
	mg.wg.Done()
}

// Group groups elements by its' md5 value. The
// keys will be strings stands for the
// hexadecimal encoded md5 values.
//
// s should not contain any non-regular files.
// Any files cannot read will be considered
// empty files.
func (mg *MD5Grouper) Group(s []file.EntryInfoType) map[file.Any][]file.EntryInfoType {
	if len(s) == 0 {
		return nil
	}

	mg.fileIn = make(chan file.EntryInfoType)

	m := make(map[file.Any][]file.EntryInfoType)

	// The most efficient number of goroutines
	// on machine.
	mg.wg.Add(runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		go mg.md5(m)
	}

	for _, t := range s {
		mg.fileIn <- t
	}

	close(mg.fileIn)

	mg.wg.Wait()

	return dedup(m)
}
