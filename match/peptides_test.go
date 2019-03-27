package match

import (
	"sort"
	"testing"

	"github.com/knightjdr/gene-peptide/types"
	"github.com/stretchr/testify/assert"
)

func TestPeptides(t *testing.T) {
	db := []types.Protein{
		{GeneID: "123", Sequence: "XXXKABCKXXX"},
		{GeneID: "456", Sequence: "DEFKABCKXXX"},
		{GeneID: "789", Sequence: "GHIKGHIK"},
		{GeneID: "101112", Sequence: "XXXKABCK"},
		{GeneID: "131415", Sequence: "DEFKXXX"},
	}
	peptides := types.Peptides{
		"ABCK": &types.PeptideStat{},
		"DEFK": &types.PeptideStat{},
		"GHIK": &types.PeptideStat{},
		"JLLK": &types.PeptideStat{},
	}

	// TEST1: match against full sequence
	wantedGenes := types.Genes{
		"123": &types.Gene{
			Peptides: []string{"ABCK"},
		},
		"456": &types.Gene{
			Peptides: []string{"ABCK", "DEFK"},
		},
		"789": &types.Gene{
			Peptides: []string{"GHIK"},
		},
		"101112": &types.Gene{
			Peptides: []string{"ABCK"},
		},
		"131415": &types.Gene{
			Peptides: []string{"DEFK"},
		},
	}
	wantedPeptides := types.Peptides{
		"ABCK": &types.PeptideStat{
			Genes: []string{"101112", "123", "456"},
		},
		"DEFK": &types.PeptideStat{
			Genes: []string{"131415", "456"},
		},
		"GHIK": &types.PeptideStat{
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

	// TEST2: match against tryptic sequence
	matchedPeptides, matchedGenes = Peptides(peptides, db, "trypsin", 1)
	for gene := range matchedGenes {
		sort.Strings(matchedGenes[gene].Peptides)
	}
	for peptide := range matchedPeptides {
		sort.Strings(matchedPeptides[peptide].Genes)
	}
	assert.Equal(t, wantedPeptides, matchedPeptides, "Should match genes to peptides")
	assert.Equal(t, wantedGenes, matchedGenes, "Should match peptides to genes")
}
