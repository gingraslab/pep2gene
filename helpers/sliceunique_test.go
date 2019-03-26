package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceUnique(t *testing.T) {
	slice := []string{"a", "b", "c", "a", "a", "b"}
	wanted := []string{"a", "b", "c"}
	assert.Equal(t, wanted, SliceUnique(slice), "Should produce slice with unique values")
}
