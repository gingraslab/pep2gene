package main

import (
	"errors"
	"os"
	"testing"

	"github.com/knightjdr/gene-peptide/types"
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
		"-missedcleavages", "2",
		"-pepprob", "0.9",
		"-pipeline", "MSPLIT_DDA",
	}
	wantArgs := types.Parameters{
		Database:           "database.fasta",
		Enzyme:             "trypsin",
		FDR:                0.05,
		File:               "peptide.txt",
		MissedCleavages:    2,
		PeptideProbability: 0.9,
		Pipeline:           "MSPLIT_DDA",
	}
	args, err := parseFlags()
	assert.Nil(t, err, "All required command line arguments should not return an error")
	assert.Equal(t, wantArgs, args, "Arguments are not returned correctly")

	// TEST2: returns default parameters when missing.
	os.Args = []string{
		"cmd",
	}
	wantArgs = types.Parameters{
		Database:           "",
		Enzyme:             "",
		FDR:                0.01,
		File:               "",
		MissedCleavages:    1,
		PeptideProbability: 0.85,
		Pipeline:           "TPP",
	}
	wantErr := errors.New("missing FASTA database; missing search result peptide file")
	args, err = parseFlags()
	assert.NotNil(t, err, "Missing arguments should return error")
	assert.Equal(t, wantArgs, args, "Default arguments were not returned")
	assert.Equal(t, wantErr, err, "Error message is not correct")
}
