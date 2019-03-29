package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilename(t *testing.T) {
	path := "/a/b/c/file.txt"
	wanted := "file"
	assert.Equal(t, wanted, Filename(path), "Should return the filename without its extension")
}
