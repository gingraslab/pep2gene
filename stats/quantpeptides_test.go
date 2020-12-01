package stats

import (
	"testing"

	"github.com/knightjdr/pep2gene/types"
	"github.com/stretchr/testify/assert"
)

func TestQuantPeptides(t *testing.T) {
	// TEST1: peptides have no intensity information
	peptides := []types.Peptide{
		{Modified: "ABC", Sequence: "ABC"},
		{Modified: "AB[mod]C", Sequence: "ABC"},
		{Modified: "AB[mod]C", Sequence: "ABC"},
		{Modified: "DE[mod]F", Sequence: "DEF"},
		{Modified: "DEF", Sequence: "DEF"},
		{Modified: "DEF", Sequence: "DEF"},
		{Modified: "DE[mod]F[mod]", Sequence: "DEF"},
		{Modified: "DE[mod]F[mod]", Sequence: "DEF"},
	}
	expected := types.Peptides{
		"ABC": &types.PeptideStat{
			Count: 3,
			Modified: map[string]float64{
				"ABC":      1,
				"AB[mod]C": 2,
			},
		},
		"DEF": &types.PeptideStat{
			Count: 5,
			Modified: map[string]float64{
				"DEF":           2,
				"DE[mod]F":      1,
				"DE[mod]F[mod]": 2,
			},
		},
	}
	assert.Equal(t, QuantifyPeptides(peptides), expected, "Should sum spectral counts for peptides")

	// TEST2: peptides have no intensity information
	peptides = []types.Peptide{
		{Intensity: 1, Modified: "ABC", Sequence: "ABC"},
		{Intensity: 2, Modified: "AB[mod]C", Sequence: "ABC"},
		{Intensity: 3, Modified: "AB[mod]C", Sequence: "ABC"},
		{Intensity: 4, Modified: "DE[mod]F", Sequence: "DEF"},
		{Intensity: 5, Modified: "DEF", Sequence: "DEF"},
		{Intensity: 6, Modified: "DEF", Sequence: "DEF"},
		{Intensity: 7, Modified: "DE[mod]F[mod]", Sequence: "DEF"},
		{Intensity: 8, Modified: "DE[mod]F[mod]", Sequence: "DEF"},
	}
	expected = types.Peptides{
		"ABC": &types.PeptideStat{
			Count: 6,
			Modified: map[string]float64{
				"ABC":      1,
				"AB[mod]C": 5,
			},
		},
		"DEF": &types.PeptideStat{
			Count: 30,
			Modified: map[string]float64{
				"DEF":           11,
				"DE[mod]F":      4,
				"DE[mod]F[mod]": 15,
			},
		},
	}
	assert.Equal(t, QuantifyPeptides(peptides), expected, "Should sum intensities for peptides")
}
