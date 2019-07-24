package digestion

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCut(t *testing.T) {
	afterRegex, _ := regexp.Compile("^P")
	cutRegex, _ := regexp.Compile("([KR])")
	// TEST1: no missed cleavages
	sequence := "ABCKDEFRPGHIRJLLKKXYZ"
	expected := map[string]bool{
		"ABCK":      true,
		"DEFRPGHIR": true,
		"JLLK":      true,
		"K":         true,
		"XYZ":       true,
	}
	assert.Equal(t, expected, cut(sequence, "c", cutRegex, afterRegex, 0), "Should produce a slice of single cleavage peptides")

	// TEST2: 1 missed cleavage
	sequence = "ABCKDEFRPGHIRJLLKXYZ"
	expected = map[string]bool{
		"ABCK":          true,
		"DEFRPGHIR":     true,
		"JLLK":          true,
		"XYZ":           true,
		"ABCKDEFRPGHIR": true,
		"DEFRPGHIRJLLK": true,
		"JLLKXYZ":       true,
	}
	assert.Equal(t, expected, cut(sequence, "c", cutRegex, afterRegex, 1), "Should produce a slice including single missed cleavage peptides")

	// TEST3: 2 missed cleavages
	sequence = "ABCKDEFRPGHIRJLLKXYZ"
	expected = map[string]bool{
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
	assert.Equal(t, expected, cut(sequence, "c", cutRegex, afterRegex, 2), "Should produce a slice including single and double missed cleavage peptides")
}
