package main

import (
	"log"

	"github.com/knightjdr/gene-peptide/peptides"
)

func main() {
	args, err := parseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	peptides.Read(args.File, args.Pipeline, args.FDR, args.PeptideProbability)
}
