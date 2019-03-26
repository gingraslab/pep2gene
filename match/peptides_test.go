package match

import (
	"sort"
	"testing"

	"github.com/knightjdr/gene-peptide/types"
	"github.com/stretchr/testify/assert"
)

func TestPeptides(t *testing.T) {
	db := []types.Protein{
		{GeneID: "123", Sequence: "XXXABCXXX"},
		{GeneID: "456", Sequence: "DEFABCXXX"},
		{GeneID: "789", Sequence: "GHIGHI"},
		{GeneID: "101112", Sequence: "XXXABC"},
		{GeneID: "131415", Sequence: "DEFXXX"},
	}
	peptides := types.Peptides{
		"ABC": &types.PeptideStat{},
		"DEF": &types.PeptideStat{},
		"GHI": &types.PeptideStat{},
	}

	// TEST1: match against full sequence
	wantedGenes := types.Genes{
		"123": &types.Gene{
			Peptides: []string{"ABC"},
		},
		"456": &types.Gene{
			Peptides: []string{"ABC", "DEF"},
		},
		"789": &types.Gene{
			Peptides: []string{"GHI"},
		},
		"101112": &types.Gene{
			Peptides: []string{"ABC"},
		},
		"131415": &types.Gene{
			Peptides: []string{"DEF"},
		},
	}
	wantedPeptides := types.Peptides{
		"ABC": &types.PeptideStat{
			Genes: []string{"101112", "123", "456"},
		},
		"DEF": &types.PeptideStat{
			Genes: []string{"131415", "456"},
		},
		"GHI": &types.PeptideStat{
			Genes: []string{"789"},
		},
	}
	matchedPeptides, matchedGenes := Peptides(peptides, db, "", 0)
	for gene := range matchedGenes {
		sort.Strings(matchedGenes[gene].Peptides)
	}
	for peptide := range matchedPeptides {
		sort.Strings(matchedPeptides[peptide].Genes)
	}
	assert.Equal(t, wantedPeptides, matchedPeptides, "Should match genes to peptides")
	assert.Equal(t, wantedGenes, matchedGenes, "Should match peptides to genes")
}
