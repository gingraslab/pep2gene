package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceUnique(t *testing.T) {
	slice := []string{"a", "b", "c", "a", "a", "b"}
	expected := []string{"a", "b", "c"}
	assert.Equal(t, expected, SliceUnique(slice), "Should produce slice with unique values")
}
