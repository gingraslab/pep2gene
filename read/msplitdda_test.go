package read

import (
	"testing"

	"github.com/knightjdr/gene-peptide/fs"
	"github.com/knightjdr/gene-peptide/types"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var msplitDDAText = `#SpecFile	SpecIndex	Scan#	FragMethod	Precursor	PMError(ppm)	Charge	Peptide	Protein	DeNovoScore	MSGFScore	SpecProb	P-value	FDR	PepFDR
file.mzXML	24063	24063	CID	753.6468	-2.1895566	4	R.ABC.L	gi|4503571|gn|ENO1:2023|	151	150	1.9758341E-35	4.6460027E-28	0.0	0.0
file.mzXML	24029	24029	CID	753.6468	-2.1895566	4	R.DEF.L	gi|4503571|gn|ENO1:2023|	171	167	8.718151E-34	2.0499977E-26	0.0	0.01
file.mzXML	24029	24029	CID	753.6468	-2.1895566	4	R.GHI.L	gi|4503571|gn|ENO1:2023|	171	167	8.718151E-34	2.0499977E-26	0.0	0.05
file.mzXML	24029	24029	CID	753.6468	-2.1895566	4	R.JK+15.995L.L	gi|4503571|gn|ENO1:2023|	171	167	8.718151E-34	2.0499977E-26	0.0	0.0
`

func TestMsplitDDASequence(t *testing.T) {
	peptide := "R.ABC.L"
	expected := "ABC"
	assert.Equal(t, expected, msplitDDASequence(peptide), "Should strip cleavage sites from peptide")
}

func TestMsplitDDARawSequence(t *testing.T) {
	peptide := "JK+15.995L"
	expected := "JKL"
	assert.Equal(t, expected, msplitDDARawSequence(peptide), "Should strip modifications from peptide")
}

func TestMsplitDDA(t *testing.T) {
	// Mock fs.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/testfile.txt",
		[]byte(msplitDDAText),
		0444,
	)

	file, _ := fs.Instance.Open("test/testfile.txt")
	actualPeptides, actualPeptideMap := msplitDDA(file, 0.01)

	// TEST.
	expectedPeptideMap := map[string]string{
		"ABC":        "ABC",
		"DEF":        "DEF",
		"JK+15.995L": "JKL",
	}
	expectedPeptides := []types.Peptide{
		{Modified: "ABC", Sequence: "ABC"},
		{Modified: "DEF", Sequence: "DEF"},
		{Modified: "JK+15.995L", Sequence: "JKL"},
	}
	assert.Equal(t, expectedPeptides, actualPeptides, "Should parse correct peptides from file")
	assert.Equal(t, expectedPeptideMap, actualPeptideMap, "Should create a map of modified peptides to raw sequence")
}
