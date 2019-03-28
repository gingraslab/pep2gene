package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyStringFloatMap(t *testing.T) {
	originalMap := map[string]float64{
		"a": 1,
		"b": 2,
	}

	// TEST1
	copiedMap := CopyStringFloatMap(originalMap)
	assert.Equal(t, originalMap, copiedMap, "Should copy original map")

	// TEST2
	copiedMap["c"] = 3
	wanted := map[string]float64{
		"a": 1,
		"b": 2,
	}
	assert.Equal(t, wanted, originalMap, "Should not modify original map when copy changes")
}
