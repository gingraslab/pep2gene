// Package peptides opens a file with peptide information and reads it.
package peptides

import (
	"log"

	"github.com/knightjdr/gene-peptide/fs"
)

// Read is an interface for opening a file and passing it to the correct parser
func Read(filename string, pipeline string, fdr, peptidProbabilty float64) {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	if pipeline == "TPP" {
		tpp(file, peptidProbabilty)
	}
}
