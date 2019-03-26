package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceContains(t *testing.T) {
	// TEST1
	a := []string{}
	b := []string{}
	assert.True(t, SliceContains(a, b), "Should return true for two empty slices")

	// TEST2
	a = []string{"a", "b", "c"}
	b = []string{"a", "b"}
	assert.True(t, SliceContains(a, b), "Should return true when the second slice is contained in the first")

	// TEST3
	a = []string{"a", "b", "c"}
	b = []string{"c"}
	assert.True(t, SliceContains(a, b), "Should return true when the second slice is contained in the first")

	// TEST4
	a = []string{"c", "a", "b"}
	b = []string{"b", "a"}
	assert.True(t, SliceContains(a, b), "Should return true when the second slice is contained in the first, ignoring order")

	// TEST5
	a = []string{"a", "b"}
	b = []string{"a", "b", "c"}
	assert.False(t, SliceContains(a, b), "Should return false when the first slice is shorter than second")

	// TEST6
	a = []string{"a", "b", "c"}
	b = []string{"a", "e"}
	assert.False(t, SliceContains(a, b), "Should return false when the second slice is not contained in the first")
}
