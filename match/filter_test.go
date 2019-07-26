package match

import (
	"testing"

	"github.com/knightjdr/pep2gene/types"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	genes := types.Genes{
		"1": &types.Gene{},
		"2": &types.Gene{},
		"5": &types.Gene{},
	}
	peptides := types.Peptides{
		"ABC": &types.PeptideStat{
			Genes: []string{"1", "2", "3"},
		},
		"DEF": &types.PeptideStat{
			Genes: []string{"4", "5"},
		},
		"GHI": &types.PeptideStat{
			Genes: []string{"6", "7", "8"},
		},
	}
	expected := types.Peptides{
		"ABC": &types.PeptideStat{
			Genes: []string{"1", "2"},
		},
		"DEF": &types.PeptideStat{
			Genes: []string{"5"},
		},
		"GHI": &types.PeptideStat{
			Genes: []string{},
		},
	}
	assert.Equal(t, expected, Filter(peptides, genes), "Should filter out genes not present in gene map")
}
