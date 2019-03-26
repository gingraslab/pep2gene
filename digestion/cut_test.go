package digestion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCut(t *testing.T) {
	re := "([KR])[^P]"
	// TEST1: no missed cleavages
	sequence := "ABCKDEFRPGHIRJLLKXYZ"
	wanted := map[string]bool{
		"ABCK":      true,
		"DEFRPGHIR": true,
		"JLLK":      true,
		"XYZ":       true,
	}
	assert.Equal(t, wanted, cut(sequence, re, "c", 0), "Should produce a slice of single cleavage peptides")

	// TEST2: 1 missed cleavage
	sequence = "ABCKDEFRPGHIRJLLKXYZ"
	wanted = map[string]bool{
		"ABCK":          true,
		"DEFRPGHIR":     true,
		"JLLK":          true,
		"XYZ":           true,
		"ABCKDEFRPGHIR": true,
		"DEFRPGHIRJLLK": true,
		"JLLKXYZ":       true,
	}
	assert.Equal(t, wanted, cut(sequence, re, "c", 1), "Should produce a slice including single missed cleavage peptides")

	// TEST3: 2 missed cleavages
	sequence = "ABCKDEFRPGHIRJLLKXYZ"
	wanted = map[string]bool{
		"ABCK":              true,
		"DEFRPGHIR":         true,
		"JLLK":              true,
		"XYZ":               true,
		"ABCKDEFRPGHIR":     true,
		"DEFRPGHIRJLLK":     true,
		"JLLKXYZ":           true,
		"ABCKDEFRPGHIRJLLK": true,
		"DEFRPGHIRJLLKXYZ":  true,
	}
	assert.Equal(t, wanted, cut(sequence, re, "c", 2), "Should produce a slice including single and double missed cleavage peptides")
}
