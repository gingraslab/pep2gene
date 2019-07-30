package main

import (
	"errors"
	"flag"
	"os"
	"strings"

	"github.com/knightjdr/pep2gene/types"
)

// ParseFlags parses command line arguments.
func parseFlags() (params types.Parameters, err error) {
	args := flag.NewFlagSet("args", flag.ContinueOnError)
	database := args.String("db", "", "FASTA database")
	enzyme := args.String("enzyme", "", "Cleavage enzyme")
	fdr := args.Float64("fdr", 0.01, "FDR cutoff")
	file := args.String("file", "", "File to process")
	inferEnzyme := args.Bool("inferenzyme", false, "Infer digestive enzyme")
	ignoreInvalid := args.Bool("ignoreinvalid", true, "Ignore invalid sequences")
	missedCleavages := args.Int("missedcleavages", 0, "Max number of missed cleavages")
	outFormat := args.String("output", "tsv", "Output file format")
	pepprob := args.Float64("pepprob", 0.85, "TPP peptide probability cutoff")
	pipeline := args.String("pipeline", "TPP", "Search engine type, should be one of MSPLIT_DDA, MSPLIT_DIA, TPP")
	args.Parse(os.Args[1:])

	params = types.Parameters{
		Database:           *database,
		Enzyme:             strings.ToLower(*enzyme),
		FDR:                *fdr,
		File:               *file,
		InferEnzyme:        *inferEnzyme,
		IgnoreInvalid:      *ignoreInvalid,
		MissedCleavages:    *missedCleavages,
		OutFormat:          *outFormat,
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

	// Format error message
	errorString := strings.Join(messages, "; ")
	if errorString != "" {
		err = errors.New(errorString)
	}

	// Set TPP as the default search engine to parse if selected engine is not recognized.
	availablePipelines := map[string]bool{
		"MSPLIT_DDA": true,
		"MSPLIT_DIA": true,
		"TPP":        true,
	}
	if _, ok := availablePipelines[params.Pipeline]; !ok {
		params.Pipeline = "TPP"
	}

	// Set tsv as the default output format if selected format is not recognized.
	availableFormats := map[string]bool{
		"json": true,
		"txt":  true,
	}
	if _, ok := availableFormats[params.OutFormat]; !ok {
		params.OutFormat = "json"
	}

	return
}
