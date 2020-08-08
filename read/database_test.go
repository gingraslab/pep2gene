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
>gi|1111|gn|BBX-1:1111| Gene test 1 [Homo sapiens]
MKGSNRNKDHSAEGEGVGKRPKRKCLQWHP
>sp|ALBU_BOVIN|gn|ALB:280717|
MKWVTFISLLLLFSSAYSRGV
>sp|CAS1_BOVIN|
MKLLILTCLVAVALARPKHPIKHQGLPQEVLNENLLRFFVAPFPEVFGKE

>Q9BYR8 SWISS-PROT:Q9BYR8 Tax_Id=9606 Gene_Symbol=KRTAP3-1;LOC100132802 Keratin-associated protein 3-1
MYCCALRSCSVPTGPATTFCSFDKSCRCGVCLPSTCPHEISLLQPICCDTCPPPCCKPDT
>gi|2222|gn|CCC:2222| Gene test 2 [Homo sapiens]
AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
>sp|CAS2_BOVIN|
MKLLILTCLVAVALARPKHPIKHQGLPQEVLNENLLRFFVAPFPEVFGKE
`

func TestAppendDatabase(t *testing.T) {
	var sequence strings.Builder

	// TEST1: an empty sequence.
	currProtein := types.Protein{
		GeneID:   "123",
		GeneName: "abc",
		Sequence: "",
		Valid:    true,
	}
	database := make([]types.Protein, 0)
	geneMap := make(map[string]string, 0)
	actualDatabase, _ := appendDatabase(database, currProtein, &sequence, geneMap, false)
	expectedDatabase := []types.Protein{}
	assert.Equal(t, expectedDatabase, actualDatabase, "Should return input protein database")

	// TEST2: an non-empty sequence
	currProtein = types.Protein{
		GeneID:   "123",
		GeneName: "abc",
		Sequence: "",
		Valid:    true,
	}
	database = make([]types.Protein, 0)
	geneMap = make(map[string]string, 0)
	sequence.WriteString("XYZ")
	actualDatabase, _ = appendDatabase(database, currProtein, &sequence, geneMap, false)
	expectedDatabase = []types.Protein{
		{GeneID: "123", GeneName: "abc", Sequence: "XYZ", Valid: true},
	}
	assert.Equal(t, expectedDatabase, actualDatabase, "Should update protein database with sequence")
	assert.Equal(t, "", sequence.String(), "Should clear string builder")

	// TEST3: an invalid sequence
	currProtein = types.Protein{
		GeneID:   "123",
		GeneName: "123",
		Sequence: "",
		Valid:    false,
	}
	database = make([]types.Protein, 0)
	geneMap = make(map[string]string, 0)
	sequence.WriteString("XYZ")
	actualDatabase, _ = appendDatabase(database, currProtein, &sequence, geneMap, true)
	expectedDatabase = []types.Protein{}
	assert.Equal(t, expectedDatabase, actualDatabase, "Should not add sequence to protein database")
}

func TestSequenceNames(t *testing.T) {
	// TEST1: read valid sequence.
	line := ">gi|443497968|gn|NISCH:11188| nischarin isoform 2"
	actualDescription, actualValid := sequenceNames(line)
	expectedDescription := map[string]string{
		"geneid":   "11188",
		"genename": "NISCH",
	}
	assert.Equal(t, expectedDescription, actualDescription, "Should read valid sequence header")
	assert.True(t, actualValid, "Should report valid sequence header")

	// TEST2: read invalid sequence.
	line = ">sp|CAS1_BOVIN|"
	actualDescription, actualValid = sequenceNames(line)
	expectedDescription = map[string]string{
		"geneid":   "p-sp|CAS1_BOVIN|",
		"genename": "p-sp|CAS1_BOVIN|",
	}
	assert.Equal(t, expectedDescription, actualDescription, "Should read invalid sequence header")
	assert.False(t, actualValid, "Should report invalid sequence header")
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

	// TEST1: read invalid sequences.
	expectedDB := []types.Protein{
		{
			GeneID:   "11188",
			GeneName: "NISCH",
			Sequence: "MATARTFGPEREAEPAKEARVVGSELVDTYLPPKKIIGKNSRSLVEKREKDLEVYLQKLL",
			Valid:    true,
		},
		{
			GeneID:   "5573",
			GeneName: "PRKAR1A",
			Sequence: "MESGSTAASEEARSLRECELYVQKHNIQALAGTRTDSREDEISPPPPNPVVKGRRRRGAILDDNERSDIFDAMFSVSFIAGETVIQQGDE",
			Valid:    true,
		},
		{
			GeneID:   "56987",
			GeneName: "BBX",
			Sequence: "MKGSNRNKDHSAEGEGVGKRPKRKCLQWHP",
			Valid:    true,
		},
		{
			GeneID:   "1111",
			GeneName: "BBX-1",
			Sequence: "MKGSNRNKDHSAEGEGVGKRPKRKCLQWHP",
			Valid:    true,
		},
		{
			GeneID:   "280717",
			GeneName: "ALB",
			Sequence: "MKWVTFISLLLLFSSAYSRGV",
			Valid:    true,
		},
		{
			GeneID:   "p-sp|CAS1_BOVIN|",
			GeneName: "p-sp|CAS1_BOVIN|",
			Sequence: "MKLLILTCLVAVALARPKHPIKHQGLPQEVLNENLLRFFVAPFPEVFGKE",
			Valid:    false,
		},
		{
			GeneID:   "p-Q9BYR8",
			GeneName: "p-Q9BYR8",
			Sequence: "MYCCALRSCSVPTGPATTFCSFDKSCRCGVCLPSTCPHEISLLQPICCDTCPPPCCKPDT",
			Valid:    false,
		},
		{
			GeneID:   "2222",
			GeneName: "CCC",
			Sequence: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			Valid:    true,
		},
		{
			GeneID:   "p-sp|CAS2_BOVIN|",
			GeneName: "p-sp|CAS2_BOVIN|",
			Sequence: "MKLLILTCLVAVALARPKHPIKHQGLPQEVLNENLLRFFVAPFPEVFGKE",
			Valid:    false,
		},
	}
	expectedGeneMap := map[string]string{
		"11188":            "NISCH",
		"5573":             "PRKAR1A",
		"56987":            "BBX",
		"1111":             "BBX-1",
		"280717":           "ALB",
		"p-sp|CAS1_BOVIN|": "p-sp|CAS1_BOVIN|",
		"p-Q9BYR8":         "p-Q9BYR8",
		"2222":             "CCC",
		"p-sp|CAS2_BOVIN|": "p-sp|CAS2_BOVIN|",
	}
	actualDB, actualGeneMap := Database("test/testfile.txt", false)
	assert.Equal(t, expectedDB, actualDB, "Should parse all proteins from database")
	assert.Equal(t, expectedGeneMap, actualGeneMap, "Should generate gene ID to name map")

	// TEST2: read valid sequences.
	expectedDB = []types.Protein{
		{
			GeneID:   "11188",
			GeneName: "NISCH",
			Sequence: "MATARTFGPEREAEPAKEARVVGSELVDTYLPPKKIIGKNSRSLVEKREKDLEVYLQKLL",
			Valid:    true,
		},
		{
			GeneID:   "5573",
			GeneName: "PRKAR1A",
			Sequence: "MESGSTAASEEARSLRECELYVQKHNIQALAGTRTDSREDEISPPPPNPVVKGRRRRGAILDDNERSDIFDAMFSVSFIAGETVIQQGDE",
			Valid:    true,
		},
		{
			GeneID:   "56987",
			GeneName: "BBX",
			Sequence: "MKGSNRNKDHSAEGEGVGKRPKRKCLQWHP",
			Valid:    true,
		},
		{
			GeneID:   "1111",
			GeneName: "BBX-1",
			Sequence: "MKGSNRNKDHSAEGEGVGKRPKRKCLQWHP",
			Valid:    true,
		},
		{
			GeneID:   "280717",
			GeneName: "ALB",
			Sequence: "MKWVTFISLLLLFSSAYSRGV",
			Valid:    true,
		},
		{
			GeneID:   "2222",
			GeneName: "CCC",
			Sequence: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			Valid:    true,
		},
	}
	expectedGeneMap = map[string]string{
		"11188":  "NISCH",
		"5573":   "PRKAR1A",
		"56987":  "BBX",
		"1111":   "BBX-1",
		"280717": "ALB",
		"2222":   "CCC",
	}
	actualDB, actualGeneMap = Database("test/testfile.txt", true)
	assert.Equal(t, expectedDB, actualDB, "Should parse valid proteins from database")
	assert.Equal(t, expectedGeneMap, actualGeneMap, "Should generate gene ID to name map")
}
