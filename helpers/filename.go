package helpers

import "path/filepath"

// Filename return the name of a file without its extension
func Filename(path string) string {
	filename := filepath.Base(path)
	extension := filepath.Ext(filename)
	return filename[0 : len(filename)-len(extension)]
}
