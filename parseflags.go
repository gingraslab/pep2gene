package main

import (
	"errors"
	"flag"
	"os"
	"strings"

	"github.com/knightjdr/gene-peptide/typedef"
)

// ParseFlags parses command line arguments.
func parseFlags() (params typedef.Parameters, err error) {
	args := flag.NewFlagSet("args", flag.ContinueOnError)
	database := args.String("db", "", "FASTA database")
	fdr := args.Float64("fdr", 0.01, "FDR cutoff")
	file := args.String("file", "", "File to process")
	mapFile := args.String("map", "", "Map of GI to gene ID to gene name, in CSV format")
	pepprob := args.Float64("pepprob", 0.85, "TPP peptide probability cutoff")
	pipeline := args.String("pipeline", "TPP", "Search engine type, should be one of MSPLIT_DDA, MSPLIT_DIA, TPP")
	args.Parse(os.Args[1:])

	params = typedef.Parameters{
		Database:           *database,
		FDR:                *fdr,
		File:               *file,
		MapFile:            *mapFile,
		PeptideProbability: *pepprob,
		Pipeline:           *pipeline,
	}

	// Validate arguments.
	messages := make([]string, 0)
	if params.Database == "" {
		messages = append(messages, "missing FASTA database")
	}
	if params.File == "" {
		messages = append(messages, "missing search result peptide file")
	}
	if params.MapFile == "" {
		messages = append(messages, "missing gene ID mapping file")
	}

	// Format error message
	errorString := strings.Join(messages, "; ")
	if errorString != "" {
		err = errors.New(errorString)
	}

	// Set TPP as the default search engine to parse if selected engine is not recognized.
	availableEngines := map[string]int{
		"MSPLIT_DDA": 1,
		"MSPLIT_DIA": 1,
		"TPP":        1,
	}
	if _, ok := availableEngines[params.Pipeline]; !ok {
		params.Pipeline = "TPP"
	}
	return
}
