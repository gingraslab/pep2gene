package output

import (
	"testing"

	"github.com/knightjdr/pep2gene/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestWriteJSON(t *testing.T) {
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
		},
		MissedCleavages:    2,
		PeptideProbability: 0.85,
		Pipeline:           "tpp",
	}
	outfile, _ := fs.Instance.Create("test/out.json")

	writeJSON(outfile, data)
	outfile.Close()

	expected := "{\n" +
		"\t\"database\": \"database.fasta\",\n" +
		"\t\"enzyme\": \"trypsin\",\n" +
		"\t\"fdr\": 0.01,\n" +
		"\t\"file\": \"file.txt\",\n" +
		"\t\"genes\": {\n" +
		"\t\t\"1\": {\n" +
		"\t\t\t\"name\": \"one\",\n" +
		"\t\t\t\"peptides\": {\n" +
		"\t\t\t\t\"AAA\": {\n" +
		"\t\t\t\t\t\"allottedSpectralCount\": 4,\n" +
		"\t\t\t\t\t\"totalSpectralCount\": 4,\n" +
		"\t\t\t\t\t\"unique\": true,\n" +
		"\t\t\t\t\t\"uniqueShared\": false\n" +
		"\t\t\t\t}\n" +
		"\t\t\t},\n" +
		"\t\t\t\"sharedIDs\": \"\",\n" +
		"\t\t\t\"sharedNames\": \"\",\n" +
		"\t\t\t\"spectralCount\": 5,\n" +
		"\t\t\t\"subsumed\": \"4\",\n" +
		"\t\t\t\"unique\": 1,\n" +
		"\t\t\t\"uniqueShared\": 0\n" +
		"\t\t}\n" +
		"\t},\n" +
		"\t\"missedCleavages\": 2,\n" +
		"\t\"peptideProbability\": 0.85,\n" +
		"\t\"pipeline\": \"tpp\"\n" +
		"}"
	bytes, _ := afero.ReadFile(fs.Instance, "test/out.json")
	assert.Equal(t, expected, string(bytes), "Should output data to json")
}
