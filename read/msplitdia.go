package read

import (
	"encoding/csv"
	"io"
	"log"
	"regexp"

	"github.com/knightjdr/gene-peptide/types"
	"github.com/spf13/afero"
)

// msplitDIARawSequence removes any modifications from a peptide and returns the raw
// amino acid sequence
func msplitDIARawSequence(peptide string) string {
	modRegex, _ := regexp.Compile("\\[.*?\\]")
	sequence := modRegex.ReplaceAllString(peptide, "")
	return sequence
}

func msplitDIA(file afero.File) []types.Peptide {
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.LazyQuotes = true

	// Skip header.
	_, err := reader.Read()
	if err != nil {
		log.Fatalln(err)
	}

	peptides := make([]types.Peptide, 0)
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		sequence := msplitDIARawSequence(line[4])
		peptides = append(peptides, types.Peptide{Modified: line[4], Sequence: sequence})
	}

	return peptides
}
