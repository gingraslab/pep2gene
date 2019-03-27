package match

import (
	"testing"

	"github.com/knightjdr/gene-peptide/types"
	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	genes := types.Genes{
		"1": &types.Gene{
			Peptides: []string{"ABC", "DEF"},
		},
		"2": &types.Gene{
			Peptides: []string{"ABC"},
		},
		"3": &types.Gene{
			Peptides: []string{"ABC", "GHI", "JKL"},
		},
	}
	peptides := types.Peptides{
		"ABC": &types.PeptideStat{
			Genes: []string{"1", "2", "3"},
		},
		"DEF": &types.PeptideStat{
			Genes: []string{"1"},
		},
		"GHI": &types.PeptideStat{
			Genes: []string{"3"},
		},
		"JKL": &types.PeptideStat{
			Genes: []string{"3"},
		},
	}
	wanted := types.Genes{
		"1": &types.Gene{
			Peptides: []string{"ABC", "DEF"},
			Unique:   1,
		},
		"2": &types.Gene{
			Peptides: []string{"ABC"},
			Unique:   0,
		},
		"3": &types.Gene{
			Peptides: []string{"ABC", "GHI", "JKL"},
			Unique:   2,
		},
	}
	assert.Equal(t, wanted, Unique(peptides, genes), "Should count unique peptides per gene")
}
