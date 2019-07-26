// Package read reads input files
package read

import (
	"errors"
	"log"

	"github.com/knightjdr/pep2gene/fs"
	"github.com/knightjdr/pep2gene/types"
)

// Peptides is an interface for opening a peptide file and passing it to the correct parser
func Peptides(filename string, pipeline string, fdr, peptideProbabilty float64) ([]types.Peptide, map[string]string) {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	peptideMap := make(map[string]string, 0)
	peptides := make([]types.Peptide, 0)
	if pipeline == "TPP" {
		peptides, peptideMap = tpp(file, peptideProbabilty)
	} else if pipeline == "MSPLIT_DDA" {
		peptides, peptideMap = msplitDDA(file, fdr)
	} else if pipeline == "MSPLIT_DIA" {
		peptides, peptideMap = msplitDIA(file)
	} else {
		log.Fatalln(errors.New("Unknown pipeline"))
	}

	return peptides, peptideMap
}
