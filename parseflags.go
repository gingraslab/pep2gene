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
	ignoreDecoys := args.Bool("ignoredecoys", true, "Ignore decoy peptides")
	ignoreInvalid := args.Bool("ignoreinvalid", true, "Ignore invalid sequences")
	inferEnzyme := args.Bool("inferenzyme", false, "Infer digestive enzyme")
	missedCleavages := args.Int("missedcleavages", 0, "Max number of missed cleavages")
	mScore := args.Float64("mscore", 0.05, "m_score to filter by")
	mScorePeptideExperimentWide := args.Float64("mscorepeptideexperimentwide", 0.01, "m_score_peptide_experiment_wide to filter by")
	outFormat := args.String("output", "tsv", "Output file format")
	peakGroupRank := args.Int("peakgrouprank", 1, "peak_group_rank to filter by")
	pepprob := args.Float64("pepprob", 0.85, "TPP peptide probability cutoff")
	pipeline := args.String("pipeline", "TPP", "Search engine type, should be one of MSPLIT_DDA, MSPLIT_DIA, OPENSWATH, TPP")
	args.Parse(os.Args[1:])

	params = types.Parameters{
		Database:                    *database,
		Enzyme:                      strings.ToLower(*enzyme),
		FDR:                         *fdr,
		File:                        *file,
		IgnoreDecoys:                *ignoreDecoys,
		IgnoreInvalid:               *ignoreInvalid,
		InferEnzyme:                 *inferEnzyme,
		MissedCleavages:             *missedCleavages,
		Mscore:                      *mScore,
		MscorePeptideExperimentWide: *mScorePeptideExperimentWide,
		OutFormat:                   *outFormat,
		PeakGroupRank:               *peakGroupRank,
		PeptideProbability:          *pepprob,
		Pipeline:                    strings.ToLower(*pipeline),
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
		"msplit_dda": true,
		"msplit_dia": true,
		"openswath":  true,
		"tpp":        true,
	}
	if _, ok := availablePipelines[params.Pipeline]; !ok {
		params.Pipeline = "tpp"
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
