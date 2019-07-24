package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilename(t *testing.T) {
	path := "/a/b/c/file.txt"
	expected := "file"
	assert.Equal(t, expected, Filename(path), "Should return the filename without its extension")
}
