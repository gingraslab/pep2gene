package read

import (
	"strings"
	"testing"

	"github.com/knightjdr/gene-peptide/fs"
	"github.com/knightjdr/gene-peptide/types"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var databaseText = `>gi|443497968|gn|NISCH:11188| nischarin isoform 2 [Homo sapiens]
MATARTFGPEREAEPAKEARVVGSELVDTY
LPPKKIIGKNSRSLVEKREKDLEVYLQKLL
>gi|443497964|gn|PRKAR1A:5573| cAMP-dependent protein kinase type I-alpha regulatory subunit isoform a [Homo sapiens]
MESGSTAASEEARSLRECELYVQKHNIQAL
AGTRTDSREDEISPPPPNPVVKGRRRRGAI
LDDNERSDIFDAMFSVSFIAGETVIQQGDE
>gi|443497952|gn|BBX:56987| HMG box transcription factor BBX isoform 3 [Homo sapiens]
MKGSNRNKDHSAEGEGVGKRPKRKCLQWHP
`

func TestAppendDatabase(t *testing.T) {
	currProtein := types.Protein{
		GeneID:   "123",
		GeneName: "abc",
		GI:       "456",
		Name:     "ABC",
		Sequence: "",
	}
	proteins := make([]types.Protein, 0)
	var sequence strings.Builder

	// TEST1: an empty string build
	result := appendDatabase(proteins, currProtein, &sequence)
	assert.Equal(t, proteins, result, "Should return input protein database")

	// TEST1: an empty string build
	sequence.WriteString("XYZ")
	result = appendDatabase(proteins, currProtein, &sequence)
	wanted := []types.Protein{
		{GeneID: "123", GeneName: "abc", GI: "456", Name: "ABC", Sequence: "XYZ"},
	}
	assert.Equal(t, wanted, result, "Should return updated protein database")
	assert.Equal(t, "", sequence.String(), "Should clear string builder")
}

func TestDatabase(t *testing.T) {
	// Mock fs.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/testfile.txt",
		[]byte(databaseText),
		0444,
	)

	wantedDB := []types.Protein{
		{
			GeneID:   "11188",
			GeneName: "NISCH",
			GI:       "443497968",
			Name:     "nischarin isoform 2",
			Sequence: "MATARTFGPEREAEPAKEARVVGSELVDTYLPPKKIIGKNSRSLVEKREKDLEVYLQKLL",
		},
		{
			GeneID:   "5573",
			GeneName: "PRKAR1A",
			GI:       "443497964",
			Name:     "cAMP-dependent protein kinase type I-alpha regulatory subunit isoform a",
			Sequence: "MESGSTAASEEARSLRECELYVQKHNIQALAGTRTDSREDEISPPPPNPVVKGRRRRGAILDDNERSDIFDAMFSVSFIAGETVIQQGDE",
		},
		{
			GeneID:   "56987",
			GeneName: "BBX",
			GI:       "443497952",
			Name:     "HMG box transcription factor BBX isoform 3",
			Sequence: "MKGSNRNKDHSAEGEGVGKRPKRKCLQWHP",
		},
	}
	wantedGeneMap := map[string]string{
		"11188": "NISCH",
		"5573":  "PRKAR1A",
		"56987": "BBX",
	}
	resultDB, resultGeneMap := Database("test/testfile.txt")
	assert.Equal(t, wantedDB, resultDB, "Should parse proteins from database")
	assert.Equal(t, wantedGeneMap, resultGeneMap, "Should generate gene ID to name map")
}
