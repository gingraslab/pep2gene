package match

import (
	"sort"
	"testing"

	"github.com/knightjdr/gene-peptide/types"
	"github.com/stretchr/testify/assert"
)

func TestSharedSubsumed(t *testing.T) {
	genes := types.Genes{
		"1": &types.Gene{
			Peptides: []string{"ABC", "DEF"},
		},
		"2": &types.Gene{
			Peptides: []string{"ABC", "DEF"},
		},
		"3": &types.Gene{
			Peptides: []string{"ABC"},
		},
		"4": &types.Gene{
			Peptides: []string{"ABC"},
		},
		"5": &types.Gene{
			Peptides: []string{"ABC", "DEF", "PQR"},
		},
		"6": &types.Gene{
			Peptides: []string{"DEF"},
		},
		"7": &types.Gene{
			Peptides: []string{"DEF", "GHI"},
		},
		"8": &types.Gene{
			Peptides: []string{"JKL", "MNO"},
		},
	}
	expected := types.Genes{
		"5": &types.Gene{
			IsSubsumed: false,
			Peptides:   []string{"ABC", "DEF", "PQR"},
			Subsumed:   []string{"1", "2", "3", "4", "6"},
		},
		"7": &types.Gene{
			IsSubsumed: false,
			Peptides:   []string{"DEF", "GHI"},
			Subsumed:   []string{"6"},
		},
		"8": &types.Gene{
			IsSubsumed: false,
			Peptides:   []string{"JKL", "MNO"},
			Subsumed:   []string{},
		},
	}
	result := SharedSubsumed(genes)
	for geneID := range result {
		sort.Strings(result[geneID].Peptides)
		sort.Strings(result[geneID].Subsumed)
	}
	assert.Equal(t, expected, result, "Should determine which genes share peptides and which are subsumed by others")
}
