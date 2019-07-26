package main

import (
	"log"

	"github.com/knightjdr/pep2gene/match"
	"github.com/knightjdr/pep2gene/output"
	"github.com/knightjdr/pep2gene/read"
	"github.com/knightjdr/pep2gene/stats"
)

func main() {
	args, err := parseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	// Read peptides from file.
	peptideList, peptideMap := read.Peptides(args.File, args.Pipeline, args.FDR, args.PeptideProbability)

	// Count spectra.
	peptideSummary := stats.QuantifyPeptides(peptideList)

	// Read database.
	db, geneIDtoName := read.Database(args.Database)

	// Match genes to peptides and peptides to genes.
	matchedPeptides, matchedGenes := match.Peptides(peptideSummary, db, args.Enzyme, args.MissedCleavages)

	// Find shared and subsumed genes.
	genes := match.SharedSubsumed(matchedGenes)

	// Filter out subsumed genes from peptides.
	peptides := match.Filter(matchedPeptides, genes)

	// Find and count unique peptides for each gene.
	genes = match.Unique(peptides, genes)

	// Sum spectra for each gene.
	genes = match.Count(peptides, genes)

	// Create output.
	outputData := output.Format(args, genes, geneIDtoName, peptides, peptideMap)

	// Output.
	output.Write(args.File, args.OutFormat, outputData)
}
