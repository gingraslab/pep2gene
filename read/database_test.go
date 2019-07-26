package read

import (
	"strings"
	"testing"

	"github.com/knightjdr/pep2gene/fs"
	"github.com/knightjdr/pep2gene/types"
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
>sp|ALBU_BOVIN|gn|ALB:280717|
MKWVTFISLLLLFSSAYSRGV
>sp|CAS1_BOVIN|
MKLLILTCLVAVALARPKHPIKHQGLPQEVLNENLLRFFVAPFPEVFGKE

>Q9BYR8 SWISS-PROT:Q9BYR8 Tax_Id=9606 Gene_Symbol=KRTAP3-1;LOC100132802 Keratin-associated protein 3-1
MYCCALRSCSVPTGPATTFCSFDKSCRCGVCLPSTCPHEISLLQPICCDTCPPPCCKPDT

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
	expected := []types.Protein{
		{GeneID: "123", GeneName: "abc", GI: "456", Name: "ABC", Sequence: "XYZ"},
	}
	assert.Equal(t, expected, result, "Should return updated protein database")
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

	expectedDB := []types.Protein{
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
		{
			GeneID:   "280717",
			GeneName: "ALB",
			GI:       "ALBU_BOVIN",
			Name:     "ALB",
			Sequence: "MKWVTFISLLLLFSSAYSRGV",
		},
		{
			GeneID:   "CAS1_BOVIN",
			GeneName: "CAS1_BOVIN",
			GI:       "CAS1_BOVIN",
			Name:     "CAS1_BOVIN",
			Sequence: "MKLLILTCLVAVALARPKHPIKHQGLPQEVLNENLLRFFVAPFPEVFGKE",
		},
		{
			GeneID:   "Q9BYR8",
			GeneName: "Q9BYR8",
			GI:       "Q9BYR8",
			Name:     "Q9BYR8",
			Sequence: "MYCCALRSCSVPTGPATTFCSFDKSCRCGVCLPSTCPHEISLLQPICCDTCPPPCCKPDT",
		},
	}
	expectedGeneMap := map[string]string{
		"11188":      "NISCH",
		"5573":       "PRKAR1A",
		"56987":      "BBX",
		"280717":     "ALB",
		"CAS1_BOVIN": "CAS1_BOVIN",
		"Q9BYR8":     "Q9BYR8",
	}
	resultDB, resultGeneMap := Database("test/testfile.txt")
	assert.Equal(t, expectedDB, resultDB, "Should parse proteins from database")
	assert.Equal(t, expectedGeneMap, resultGeneMap, "Should generate gene ID to name map")
}
