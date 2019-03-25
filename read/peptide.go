// Package read reads input files
package read

import (
	"log"

	"github.com/knightjdr/gene-peptide/fs"
)

// Peptide contains the amino acid "Sequence" for a peptide, the "Modified" version of
// the peptide and whether it is "Decoy"
type Peptide struct {
	Decoy    bool
	Modified string
	Sequence string
}

// Peptides is an interface for opening a peptide file and passing it to the correct parser
func Peptides(filename string, pipeline string, fdr, peptideProbabilty float64) {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	if pipeline == "TPP" {
		tpp(file, peptideProbabilty)
	} else if pipeline == "MSPLIT_DDA" {
		msplitDDA(file, fdr)
	} else if pipeline == "MSPLIT_DIA" {
		msplitDIA(file)
	}
}
