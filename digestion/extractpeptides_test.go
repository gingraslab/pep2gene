package digestion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractPeptides(t *testing.T) {
	// TEST1: interal cut sites
	matches := []int{3, 12, 16, 17}
	sequence := "ABCKDEFRPGHIRJLLKKXYZ"
	wanted := []string{"ABCK", "DEFRPGHIR", "JLLK", "K", "XYZ"}
	assert.Equal(t, wanted, extractPeptides(sequence, "c", matches), "Should extract peptides from internal cut sites")

	// TEST2: leading cut site
	matches = []int{0, 12, 16}
	sequence = "KABCDEFRPGHIRJLLKXYZ"
	wanted = []string{"K", "ABCDEFRPGHIR", "JLLK", "XYZ"}
	assert.Equal(t, wanted, extractPeptides(sequence, "c", matches), "Should extract peptides with leading cut site")

	// TEST2: leading cut site
	matches = []int{3, 12, 19}
	sequence = "ABCKDEFRPGHIRJLLXYZK"
	wanted = []string{"ABCK", "DEFRPGHIR", "JLLXYZK"}
	assert.Equal(t, wanted, extractPeptides(sequence, "c", matches), "Should extract peptides with trailing cut site")

	// TEST4: N-terminal cut sites
	matches = []int{3, 12, 16}
	sequence = "ABCKDEFRPGHIRJLLKXYZ"
	wanted = []string{"ABC", "KDEFRPGHI", "RJLL", "KXYZ"}
	assert.Equal(t, wanted, extractPeptides(sequence, "n", matches), "Should extract peptides for N-terminal cleavage")
}
