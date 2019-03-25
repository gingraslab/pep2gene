package main

import (
	"log"

	"github.com/knightjdr/gene-peptide/read"
)

func main() {
	args, err := parseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	read.Peptides(args.File, args.Pipeline, args.FDR, args.PeptideProbability)
}
