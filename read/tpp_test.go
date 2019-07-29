package read

import (
	"testing"

	"github.com/knightjdr/pep2gene/fs"
	"github.com/knightjdr/pep2gene/types"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var tppText = `<sample_enzyme name="trypsin">
<search_hit hit_rank="1" peptide="ABC">
<peptideprophet_result probability="0.95">
</peptideprophet_result>
</search_hit>
<search_hit hit_rank="2" peptide="DEF">
<peptideprophet_result probability="0.8">
</peptideprophet_result>
</search_hit>
<search_hit hit_rank="1" peptide="GHI">
<alternative_protein protein="DECOY12345"/>
<peptideprophet_result probability="0.95">
</peptideprophet_result>
</search_hit>
<search_hit hit_rank="1" peptide="JKL">
<modification_info modified_peptide="JK[129]L">
<peptideprophet_result probability="0.9">
</peptideprophet_result>
</search_hit>`

func TestTPP(t *testing.T) {
	// Mock fs.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/testfile.txt",
		[]byte(tppText),
		0444,
	)

	
	// TEST1: infer enzyme.
	file, _ := fs.Instance.Open("test/testfile.txt")
	actualPeptides, actualPeptideMap, actualEnzyme := tpp(file, 0.85, true)
	expectedPeptideMap := map[string]string{
		"ABC":      "ABC",
		"GHI":      "GHI",
		"JK[129]L": "JKL",
	}
	expectedPeptides := []types.Peptide{
		{Modified: "ABC", Sequence: "ABC"},
		{Modified: "GHI", Sequence: "GHI"},
		{Modified: "JK[129]L", Sequence: "JKL"},
	}
	assert.Equal(t, expectedPeptides, actualPeptides, "Should parse correct peptides from file")
	assert.Equal(t, expectedPeptideMap, actualPeptideMap, "Should create a map of modified peptides to raw sequence")
	assert.Equal(t, "trypsin", actualEnzyme, "Should infer enzyme")

	// TEST2: do not infer enzyme.
	file, _ = fs.Instance.Open("test/testfile.txt")
	_, _, actualEnzyme = tpp(file, 0.85, false)
	assert.Equal(t, "", actualEnzyme, "Should not infer enzyme")
}
