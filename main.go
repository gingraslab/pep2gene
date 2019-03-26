package main

import (
	"log"

	"github.com/knightjdr/gene-peptide/match"
	"github.com/knightjdr/gene-peptide/read"
	"github.com/knightjdr/gene-peptide/stats"
)

func main() {
	args, err := parseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	// Read peptides from file.
	peptides := read.Peptides(args.File, args.Pipeline, args.FDR, args.PeptideProbability)

	// Count spectra.
	peptideSummary := stats.QuantifyPeptides(peptides)

	// Read database.
	db := read.Database(args.Database)

	// Match peptides to genes
	_, matchedGenes := match.Peptides(peptideSummary, db, args.Enzyme, args.MissedCleavages)

	// Find shared and subsumed genes.
	match.SharedSubsumed(matchedGenes)
}
