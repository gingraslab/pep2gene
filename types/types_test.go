package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneCopy(t *testing.T) {
	originalGene := &Gene{
		Count:      5,
		IsSubsumed: false,
		PeptideCount: map[string]float64{
			"ABC": 1,
		},
		Peptides:     []string{"ABC", "DEF"},
		Shared:       []string{"b"},
		Subsumed:     []string{"c"},
		Unique:       3,
		UniqueShared: 5,
	}

	// TEST1
	copiedGene := originalGene.Copy()
	assert.Equal(t, copiedGene, originalGene, "Should copy Gene type")

	// TEST2
	copiedGene.Count = 10
	copiedGene.IsSubsumed = true
	copiedGene.PeptideCount = map[string]float64{
		"ABC": 1,
		"DEF": 2,
	}
	copiedGene.Peptides = []string{"ABC", "DEF", "GHI"}
	copiedGene.Shared = []string{"b", "c"}
	copiedGene.Subsumed = []string{}
	copiedGene.Unique = 4
	copiedGene.UniqueShared = 6
	wanted := &Gene{
		Count:      5,
		IsSubsumed: false,
		PeptideCount: map[string]float64{
			"ABC": 1,
		},
		Peptides:     []string{"ABC", "DEF"},
		Shared:       []string{"b"},
		Subsumed:     []string{"c"},
		Unique:       3,
		UniqueShared: 5,
	}
	assert.Equal(t, wanted, originalGene, "Should not modify original Gene when copy changes")
}

func TestPeptideStatCopy(t *testing.T) {
	originalPeptideStat := &PeptideStat{
		Count: 5,
		Genes: []string{"a", "b"},
		Modified: map[string]int{
			"ABC": 1,
			"DEF": 2,
		},
	}

	// TEST1
	copiedPeptideStat := originalPeptideStat.Copy()
	assert.Equal(t, copiedPeptideStat, originalPeptideStat, "Should copy PeptideStat type")

	// TEST2
	copiedPeptideStat.Count = 10
	copiedPeptideStat.Genes = []string{"b", "c"}
	copiedPeptideStat.Modified = map[string]int{
		"ABC": 2,
		"DEF": 4,
	}
	wanted := &PeptideStat{
		Count: 5,
		Genes: []string{"a", "b"},
		Modified: map[string]int{
			"ABC": 1,
			"DEF": 2,
		},
	}
	assert.Equal(t, wanted, originalPeptideStat, "Should not modify original Peptide when copy changes")
}
