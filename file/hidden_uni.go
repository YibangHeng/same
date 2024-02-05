//go:build !windows

package file

func IsHiddenFile(filename string) bool {
	return filename[0] == '.'
}
