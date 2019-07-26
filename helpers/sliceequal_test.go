package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceEqual(t *testing.T) {
	// TEST1
	a := []string{}
	b := []string{}
	assert.True(t, SliceEqual(a, b), "Should return true for two empty slices")

	// TEST2
	a = []string{"x", "y", "z"}
	b = []string{"x", "y", "z"}
	assert.True(t, SliceEqual(a, b), "Should return true for two slices with same elements in same order")

	// TEST3
	a = []string{"x", "y", "z"}
	b = []string{"x", "z", "y"}
	assert.True(t, SliceEqual(a, b), "Should return true for two slices with same elements in different order")

	// TEST4
	a = []string{"x", "y", "z"}
	b = []string{"x", "y"}
	assert.False(t, SliceEqual(a, b), "Should return false for two slices of different lengths")

	// TEST4
	a = []string{"x", "y", "z"}
	b = []string{"x", "y", "zz"}
	assert.False(t, SliceEqual(a, b), "Should return false for two slices with different elements")
}
