// Package read reads input files
package read

import (
	"errors"
	"log"

	"github.com/knightjdr/gene-peptide/fs"
	"github.com/knightjdr/gene-peptide/types"
)

// Peptides is an interface for opening a peptide file and passing it to the correct parser
func Peptides(filename string, pipeline string, fdr, peptideProbabilty float64) []types.Peptide {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	peptides := make([]types.Peptide, 0)
	if pipeline == "TPP" {
		peptides = tpp(file, peptideProbabilty)
	} else if pipeline == "MSPLIT_DDA" {
		peptides = msplitDDA(file, fdr)
	} else if pipeline == "MSPLIT_DIA" {
		peptides = msplitDIA(file)
	} else {
		log.Fatalln(errors.New("Unknown pipeline"))
	}

	return peptides
}
