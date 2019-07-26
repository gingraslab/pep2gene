package stats

import (
	"testing"

	"github.com/knightjdr/pep2gene/types"
	"github.com/stretchr/testify/assert"
)

func TestQuantPeptides(t *testing.T) {
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
			Modified: map[string]int{
				"ABC":      1,
				"AB[mod]C": 2,
			},
		},
		"DEF": &types.PeptideStat{
			Count: 5,
			Modified: map[string]int{
				"DEF":           2,
				"DE[mod]F":      1,
				"DE[mod]F[mod]": 2,
			},
		},
	}
	assert.Equal(t, QuantifyPeptides(peptides), expected, "Should generate spectral counts for peptides")
}
