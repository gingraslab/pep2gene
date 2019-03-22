// Package fs creates a filesystem to use (for easy mocking)
package fs

import "github.com/spf13/afero"

// Instance contains the file system instance
var Instance = afero.NewOsFs()
