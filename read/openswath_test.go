package read

import (
	"testing"

	"github.com/knightjdr/pep2gene/fs"
	"github.com/knightjdr/pep2gene/types"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

var openswathText = "transition_group_id\tdecoy\t\t\t\t\t\t\t\t\t\tSequence\tFullPeptideName\t\t\tIntensity\t\t\t\t\tpeak_group_rank\t\tm_score\t\t\t\t\tm_score_peptide_experiment_wide\t\n" +
	"1\t1\t\t\t\t\t\t\t\t\t\tAAAA\tAAAA\t\t\t11\t\t\t\t\t1\t\t0.05\t\t\t\t\t0.01\t\n" +
	"2\t0\t\t\t\t\t\t\t\t\t\tBBBB\tBBBB\t\t\t12\t\t\t\t\t1\t\t0.05\t\t\t\t\t0.01\t\n" +
	"3\t0\t\t\t\t\t\t\t\t\t\tCCCC\tCCCC\t\t\t13\t\t\t\t\t3\t\t0.05\t\t\t\t\t0.01\t\n" +
	"4\t0\t\t\t\t\t\t\t\t\t\tDDDD\tDDDD\t\t\t14\t\t\t\t\t1\t\t0.06\t\t\t\t\t0.01\t\n" +
	"5\t0\t\t\t\t\t\t\t\t\t\tEEEE\tEEEE\t\t\t15\t\t\t\t\t1\t\t0.05\t\t\t\t\t0.02\t\n" +
	"6\t0\t\t\t\t\t\t\t\t\t\tFFFF\tFFF[(UniMod:7)]F\t\t\t16\t\t\t\t\t1\t\t0.05\t\t\t\t\t0.01\t\n" +
	"7\t0\t\t\t\t\t\t\t\t\t\tGGGG\tGGGG\t\t\t17\t\t\t\t\t2\t\t7.3E-04\t\t\t\t\t1.0E-06\t\n" +
	"8\t0\t\t\t\t\t\t\t\t\t\tGGGG\tGGGG\t\t\t18\t\t\t\t\t2\t\t0.05\t\t\t\t\t0\t\n"

func TestPassesOpenSwathFilters(t *testing.T) {
	// TEST 1: is decoy and decoys are ignored
	line := []string{"1", "1", "", "", "", "", "", "", "", "", "", "AAAA", "AAAA", "", "", "11", "", "", "", "", "1", "", "0.05", "", "", "", "", "0.01", ""}
	options := types.Parameters{
		IgnoreDecoys:                true,
		Mscore:                      0.05,
		MscorePeptideExperimentWide: 0.01,
		PeakGroupRank:               1,
	}
	assert.False(t, passesOpenSwathFilters(line, options), "Should return false for decoy when ignoring decoys")

	// TEST 2: is decoy and decoys are not ignored
	options.IgnoreDecoys = false
	assert.True(t, passesOpenSwathFilters(line, options), "Should return true for decoy when not ignoring decoys")

	// TEST 3: does not pass Mscore
	options.IgnoreDecoys = true
	line = []string{"1", "0", "", "", "", "", "", "", "", "", "", "AAAA", "AAAA", "", "", "11", "", "", "", "", "1", "", "0.06", "", "", "", "", "0.01", ""}
	assert.False(t, passesOpenSwathFilters(line, options), "Should return false when not passing Mscore")

	// TEST 4: does not pass MscorePeptideExperimentWide
	line = []string{"1", "0", "", "", "", "", "", "", "", "", "", "AAAA", "AAAA", "", "", "11", "", "", "", "", "1", "", "0.05", "", "", "", "", "0.02", ""}
	assert.False(t, passesOpenSwathFilters(line, options), "Should return false when not passing MscorePeptideExperimentWide")

	// TEST 5: does not pass PeakGroupRank
	line = []string{"1", "0", "", "", "", "", "", "", "", "", "", "AAAA", "AAAA", "", "", "11", "", "", "", "", "2", "", "0.05", "", "", "", "", "0.01", ""}
	assert.False(t, passesOpenSwathFilters(line, options), "Should return false when not passing PeakGroupRank")

	// TEST 6: passes all filters
	line = []string{"1", "0", "", "", "", "", "", "", "", "", "", "AAAA", "AAAA", "", "", "11", "", "", "", "", "1", "", "0.05", "", "", "", "", "0.01", ""}
	assert.True(t, passesOpenSwathFilters(line, options), "Should return true when passing all filters")
}

func TestOpenSwath(t *testing.T) {
	// Mock fs.
	oldFs := fs.Instance
	defer func() { fs.Instance = oldFs }()
	fs.Instance = afero.NewMemMapFs()

	// Create test directory and files.
	fs.Instance.MkdirAll("test", 0755)
	afero.WriteFile(
		fs.Instance,
		"test/testfile.txt",
		[]byte(openswathText),
		0444,
	)

	file, _ := fs.Instance.Open("test/testfile.txt")

	options := types.Parameters{
		IgnoreDecoys:                true,
		Mscore:                      0.05,
		MscorePeptideExperimentWide: 0.01,
		PeakGroupRank:               2,
	}
	actualPeptides, actualPeptideMap := openswath(file, options)

	expectedPeptideMap := map[string]string{
		"BBBB":             "BBBB",
		"FFF[(UniMod:7)]F": "FFFF",
		"GGGG":             "GGGG",
	}
	expectedPeptides := []types.Peptide{
		{Intensity: 12, Modified: "BBBB", Sequence: "BBBB"},
		{Intensity: 16, Modified: "FFF[(UniMod:7)]F", Sequence: "FFFF"},
		{Intensity: 17, Modified: "GGGG", Sequence: "GGGG"},
		{Intensity: 18, Modified: "GGGG", Sequence: "GGGG"},
	}
	assert.Equal(t, expectedPeptides, actualPeptides, "Should parse correct peptides from file")
	assert.Equal(t, expectedPeptideMap, actualPeptideMap, "Should create a map of modified peptides to raw sequence")
}
