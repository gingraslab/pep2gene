package main

import (
	"errors"
	"os"
	"testing"

	"github.com/knightjdr/pep2gene/types"
	"github.com/stretchr/testify/assert"
)

func TestParseFlags(t *testing.T) {
	// Argument unmocking.
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	// TEST1: returns parameters with correct flags specied.
	os.Args = []string{
		"cmd",
		"-db", "database.fasta",
		"-enzyme", "Trypsin",
		"-fdr", "0.05",
		"-file", "peptide.txt",
		"-ignoredecoys=false",
		"-ignoreinvalid=false",
		"-inferenzyme",
		"-missedcleavages", "2",
		"-mscore", "0.1",
		"-mscorepeptideexperimentwide", "0.1",
		"-output", "json",
		"-peakgrouprank", "2",
		"-pepprob", "0.9",
		"-pipeline", "MSPLIT_DDA",
	}
	wantArgs := types.Parameters{
		Database:                    "database.fasta",
		Enzyme:                      "trypsin",
		FDR:                         0.05,
		File:                        "peptide.txt",
		IgnoreDecoys:                false,
		IgnoreInvalid:               false,
		InferEnzyme:                 true,
		MissedCleavages:             2,
		Mscore:                      0.1,
		MscorePeptideExperimentWide: 0.1,
		OutFormat:                   "json",
		PeakGroupRank:               2,
		PeptideProbability:          0.9,
		Pipeline:                    "msplit_dda",
	}
	args, err := parseFlags()
	assert.Nil(t, err, "Should not return an error when all required command line arguments are present")
	assert.Equal(t, wantArgs, args, "Should return arguments as parameters")

	// TEST2: returns default parameters when missing.
	os.Args = []string{
		"cmd",
	}
	wantArgs = types.Parameters{
		Database:                    "",
		FDR:                         0.01,
		File:                        "",
		IgnoreDecoys:                true,
		IgnoreInvalid:               true,
		InferEnzyme:                 false,
		MissedCleavages:             0,
		Mscore:                      0.05,
		MscorePeptideExperimentWide: 0.01,
		OutFormat:                   "json",
		PeakGroupRank:               1,
		PeptideProbability:          0.85,
		Pipeline:                    "tpp",
	}
	wantErr := errors.New("missing FASTA database; missing search result peptide file")
	args, err = parseFlags()
	assert.NotNil(t, err, "Should return error when missing arguments")
	assert.Equal(t, wantArgs, args, "Should return default parameters")
	assert.Equal(t, wantErr, err, "Should return correct error message")

	// TEST3: returns default parameters when specifed args are not recognized.
	os.Args = []string{
		"cmd",
		"-db", "database.fasta",
		"-enzyme", "trypsin",
		"-fdr", "0.05",
		"-file", "peptide.txt",
		"-missedcleavages", "2",
		"-output", "Unknown",
		"-pepprob", "0.9",
		"-pipeline", "unknown",
	}
	args, _ = parseFlags()
	assert.Equal(t, "json", args.OutFormat, "Should set output format to tsv when arg not recognized")
	assert.Equal(t, "tpp", args.Pipeline, "Should set pipeline to TPP when arg not recognized")
}
