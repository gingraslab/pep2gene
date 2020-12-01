// Package read reads input files
package read

import (
	"errors"
	"log"

	"github.com/knightjdr/pep2gene/fs"
	"github.com/knightjdr/pep2gene/types"
)

// Peptides is an interface for opening a peptide file and passing it to the correct parser
// func Peptides(filename string, pipeline string, fdr, peptideProbabilty float64, inferEnzyme bool) ([]types.Peptide, map[string]string, string) {
func Peptides(options types.Parameters) ([]types.Peptide, map[string]string, string) {
	file, err := fs.Instance.Open(options.File)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	enzyme := ""
	peptideMap := make(map[string]string, 0)
	peptides := make([]types.Peptide, 0)
	if options.Pipeline == "tpp" {
		peptides, peptideMap, enzyme = tpp(file, options.PeptideProbability, options.InferEnzyme)
	} else if options.Pipeline == "msplit_dda" {
		peptides, peptideMap = msplitDDA(file, options.FDR)
	} else if options.Pipeline == "msplit_dia" {
		peptides, peptideMap = msplitDIA(file)
	} else if options.Pipeline == "openswath" {
		peptides, peptideMap = openswath(file, options)
	} else {
		log.Fatalln(errors.New("Unknown pipeline"))
	}

	return peptides, peptideMap, enzyme
}
