package output

import (
	"testing"

	"github.com/knightjdr/gene-peptide/types"
	"github.com/stretchr/testify/assert"
)

func TestSummarizePeptides(t *testing.T) {
	// TEST1: single gene.
	genes := []string{"1"}
	peptideCount := map[string]float64{
		"AAA":      3,
		"BBB":      1,
		"CCC[155]": 2,
	}
	peptideMap := map[string]string{
		"AAA":      "AAA",
		"BBB":      "BBB",
		"CCC[155]": "CCC",
	}
	peptides := types.Peptides{
		"AAA": &types.PeptideStat{
			Genes: []string{"1"},
			Modified: map[string]int{
				"AAA": 3,
			},
			Unique: true,
		},
		"BBB": &types.PeptideStat{
			Genes: []string{"1"},
			Modified: map[string]int{
				"BBB": 2,
			},
			Unique: false,
		},
		"CCC": &types.PeptideStat{
			Genes: []string{"1"},
			Modified: map[string]int{
				"CCC[155]": 2,
			},
			Unique: true,
		},
	}

	actual := summarizePeptides(genes, peptideCount, peptides, peptideMap)
	expected := map[string]Peptide{
		"AAA": Peptide{
			AllottedSpectralCount: 3,
			TotalSpectralCount:    3,
			Unique:                true,
			UniqueShared:          false,
		},
		"BBB": Peptide{
			AllottedSpectralCount: 1,
			TotalSpectralCount:    2,
			Unique:                false,
			UniqueShared:          false,
		},
		"CCC[155]": Peptide{
			AllottedSpectralCount: 2,
			TotalSpectralCount:    2,
			Unique:                true,
			UniqueShared:          false,
		},
	}
	assert.Equal(t, expected, actual, "Should format peptides for output for a single gene")

	// TEST2: two genes sharing peptides.
	genes = []string{"1", "2"}
	peptideCount = map[string]float64{
		"AAA":      1.5,
		"BBB":      1,
		"CCC[155]": 1,
	}
	peptideMap = map[string]string{
		"AAA":      "AAA",
		"BBB":      "BBB",
		"CCC[155]": "CCC",
	}
	peptides = types.Peptides{
		"AAA": &types.PeptideStat{
			Genes: []string{"1", "2"},
			Modified: map[string]int{
				"AAA": 3,
			},
			Unique: false,
		},
		"BBB": &types.PeptideStat{
			Genes: []string{"1", "2"},
			Modified: map[string]int{
				"BBB": 2,
			},
			Unique: false,
		},
		"CCC": &types.PeptideStat{
			Genes: []string{"1", "2", "3"},
			Modified: map[string]int{
				"CCC[155]": 3,
			},
			Unique: false,
		},
	}

	actual = summarizePeptides(genes, peptideCount, peptides, peptideMap)
	expected = map[string]Peptide{
		"AAA": Peptide{
			AllottedSpectralCount: 1.5,
			TotalSpectralCount:    3,
			Unique:                false,
			UniqueShared:          true,
		},
		"BBB": Peptide{
			AllottedSpectralCount: 1,
			TotalSpectralCount:    2,
			Unique:                false,
			UniqueShared:          true,
		},
		"CCC[155]": Peptide{
			AllottedSpectralCount: 1,
			TotalSpectralCount:    3,
			Unique:                false,
			UniqueShared:          false,
		},
	}
	assert.Equal(t, expected, actual, "Should format peptides for output for two genes sharing peptides")
}

func TestFormat(t *testing.T) {
	geneIDtoName := map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
	}
	genes := types.Genes{
		"1": &types.Gene{
			Count: 5,
			PeptideCount: map[string]float64{
				"AAA": 4,
				"BBB": 1,
			},
			Shared:   []string{},
			Subsumed: []string{"4"},
			Unique:   1,
		},
		"2": &types.Gene{
			Count: 3,
			PeptideCount: map[string]float64{
				"BBB":      1,
				"CCC[115]": 2,
			},
			Shared:       []string{"3"},
			UniqueShared: 1,
		},
	}
	options := types.Parameters{
		Database:           "path/to/database.fasta",
		Enzyme:             "trypsin",
		FDR:                0.01,
		File:               "path/to/file.txt",
		MissedCleavages:    2,
		PeptideProbability: 0.85,
		Pipeline:           "tpp",
	}
	peptideMap := map[string]string{
		"AAA":      "AAA",
		"BBB":      "BBB",
		"CCC[115]": "CCC",
	}
	peptides := types.Peptides{
		"AAA": &types.PeptideStat{
			Genes: []string{"1"},
			Modified: map[string]int{
				"AAA": 4,
			},
			Unique: true,
		},
		"BBB": &types.PeptideStat{
			Genes: []string{"1", "2", "3"},
			Modified: map[string]int{
				"BBB": 3,
			},
			Unique: false,
		},
		"CCC": &types.PeptideStat{
			Genes: []string{"2", "3"},
			Modified: map[string]int{
				"CCC[115]": 4,
			},
			Unique: false,
		},
	}

	actual := Format(options, genes, geneIDtoName, peptides, peptideMap)
	expected := Data{
		Database: "database.fasta",
		Enzyme:   "trypsin",
		FDR:      0.01,
		File:     "file.txt",
		Genes: map[string]*Gene{
			"1": &Gene{
				Name: "one",
				Peptides: map[string]Peptide{
					"AAA": Peptide{
						AllottedSpectralCount: 4,
						TotalSpectralCount:    4,
						Unique:                true,
						UniqueShared:          false,
					},
					"BBB": Peptide{
						AllottedSpectralCount: 1,
						TotalSpectralCount:    3,
						Unique:                false,
						UniqueShared:          false,
					},
				},
				SharedIDs:     "",
				SharedNames:   "",
				SpectralCount: 5,
				Subsumed:      "4",
				Unique:        1,
				UniqueShared:  0,
			},
			"2": &Gene{
				Name: "two",
				Peptides: map[string]Peptide{
					"BBB": Peptide{
						AllottedSpectralCount: 1,
						TotalSpectralCount:    3,
						Unique:                false,
						UniqueShared:          false,
					},
					"CCC[115]": Peptide{
						AllottedSpectralCount: 2,
						TotalSpectralCount:    4,
						Unique:                false,
						UniqueShared:          true,
					},
				},
				SharedIDs:     "3",
				SharedNames:   "three",
				SpectralCount: 3,
				Subsumed:      "",
				Unique:        0,
				UniqueShared:  1,
			},
		},
		MissedCleavages:    2,
		PeptideProbability: 0.85,
		Pipeline:           "tpp",
	}
	assert.Equal(t, expected, actual, "Should format genes for output")
}
