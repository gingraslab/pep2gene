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
		"4": &types.Gene{
			Peptides: []string{"MNO", "PQR"},
			Shared:   []string{"5"},
		},
		"5": &types.Gene{
			Peptides: []string{"MNO", "PQR"},
			Shared:   []string{"4"},
		},
	}
	peptides := types.Peptides{
		"ABC": &types.PeptideStat{
			Genes: []string{"1", "2", "3"},
		},
		"DEF": &types.PeptideStat{
			Genes:  []string{"1"},
			Unique: true,
		},
		"GHI": &types.PeptideStat{
			Genes:  []string{"3"},
			Unique: true,
		},
		"JKL": &types.PeptideStat{
			Genes:  []string{"3"},
			Unique: true,
		},
		"MNO": &types.PeptideStat{
			Genes: []string{"4", "5"},
		},
		"PQR": &types.PeptideStat{
			Genes: []string{"4", "5", "6"},
		},
	}
	expected := types.Genes{
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
		"4": &types.Gene{
			Peptides:     []string{"MNO", "PQR"},
			Shared:       []string{"5"},
			Unique:       0.5,
			UniqueShared: 1,
		},
		"5": &types.Gene{
			Peptides:     []string{"MNO", "PQR"},
			Shared:       []string{"4"},
			Unique:       0.5,
			UniqueShared: 1,
		},
	}
	assert.Equal(t, expected, Unique(peptides, genes), "should count unique peptides per gene")
}
