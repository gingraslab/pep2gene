package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyStringIntMap(t *testing.T) {
	originalMap := map[string]int{
		"a": 1,
		"b": 2,
	}

	// TEST1
	copiedMap := CopyStringIntMap(originalMap)
	assert.Equal(t, originalMap, copiedMap, "Should copy original map")

	// TEST2
	copiedMap["c"] = 3
	expected := map[string]int{
		"a": 1,
		"b": 2,
	}
	assert.Equal(t, expected, originalMap, "Should not modify original map when copy changes")
}
