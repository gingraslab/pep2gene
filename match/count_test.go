package match

import (
	"testing"

	"github.com/knightjdr/pep2gene/types"
	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	genes := types.Genes{
		"1": &types.Gene{
			Unique: 1,
		},
		"2": &types.Gene{
			Unique: 0,
		},
		"3": &types.Gene{
			Unique: 2,
		},
		"4": &types.Gene{
			Unique: 0,
		},
	}
	peptides := types.Peptides{
		"ABC": &types.PeptideStat{
			Count: 15,
			Genes: []string{"1", "3"},
			Modified: map[string]float64{
				"ABC": 15,
			},
		},
		"DEF": &types.PeptideStat{
			Count: 20,
			Genes: []string{"1"},
			Modified: map[string]float64{
				"DEF":      15,
				"DE[mod]F": 5,
			},
		},
		"GHI": &types.PeptideStat{
			Count: 30,
			Genes: []string{"1", "2", "3"},
			Modified: map[string]float64{
				"GHI":      15,
				"GH[mod]I": 15,
			},
		},
		"JKL": &types.PeptideStat{
			Count: 20,
			Genes: []string{"2", "4"},
			Modified: map[string]float64{
				"JKL": 20,
			},
		},
	}
	expected := types.Genes{
		"1": &types.Gene{
			Count: float64(35),
			PeptideCount: map[string]float64{
				"ABC":      float64(5),
				"DEF":      float64(15),
				"DE[mod]F": float64(5),
				"GHI":      float64(5),
				"GH[mod]I": float64(5),
			},
			Unique: 1,
		},
		"2": &types.Gene{
			Count: float64(10),
			PeptideCount: map[string]float64{
				"JKL": float64(10),
			},
			Unique: 0,
		},
		"3": &types.Gene{
			Count: float64(30),
			PeptideCount: map[string]float64{
				"ABC":      float64(10),
				"GHI":      float64(10),
				"GH[mod]I": float64(10),
			},
			Unique: 2,
		},
		"4": &types.Gene{
			Count: float64(10),
			PeptideCount: map[string]float64{
				"JKL": float64(10),
			},
			Unique: 0,
		},
	}
	assert.Equal(t, expected, Count(peptides, genes), "Should sum spectral counts")
}
