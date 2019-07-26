package output

import (
	"testing"

	"github.com/knightjdr/gene-peptide/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestWriteTXT(t *testing.T) {
	// Mock fs.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)

	data := Data{
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
	outfile, _ := fs.Instance.Create("test/out.txt")

	writeTXT(outfile, data.Genes)
	outfile.Close()

	expected := "HitNumber;;Gene;;GeneID;;SpectralCount;;Unique;;Subsumed\n" +
		"Peptide;;TotalSpectralCount;;IsUnique\n\n" +
		"Hit_1;;one;;1;;5.00;;1;;4\n" +
		"AAA;;4;;yes\n\n" +
		"Hit_2;;two, three;;2, 3;;3.00;;0;;\n" +
		"BBB;;3;;no\n" +
		"CCC[115];;4;;no\n"
	bytes, _ := afero.ReadFile(fs.Instance, "test/out.txt")
	assert.Equal(t, expected, string(bytes), "Should output data to txt")
}
