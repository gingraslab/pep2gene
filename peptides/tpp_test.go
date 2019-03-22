package peptides

import (
	"testing"

	"github.com/knightjdr/gene-peptide/fs"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var text = `<search_hit hit_rank="1" peptide="ABC">
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
		[]byte(text),
		0444,
	)

	file, _ := fs.Instance.Open("test/testfile.txt")
	peptides := tpp(file, 0.85)

	// TEST
	wanted := []Peptide{
		{Decoy: false, Modified: "ABC", Sequence: "ABC"},
		{Decoy: false, Modified: "JK[129]L", Sequence: "JKL"},
	}
	assert.Equal(t, wanted, peptides, "Should parse correct peptides from file")
}
