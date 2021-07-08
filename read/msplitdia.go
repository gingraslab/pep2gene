package read

import (
	"encoding/csv"
	"io"
	"log"
	"regexp"

	"github.com/gingraslab/pep2gene/types"
	"github.com/spf13/afero"
)

// msplitDIARawSequence removes any modifications from a peptide and returns the raw
// amino acid sequence
func msplitDIARawSequence(peptide string) string {
	modRegex, _ := regexp.Compile("\\[.*?\\]")
	sequence := modRegex.ReplaceAllString(peptide, "")
	return sequence
}

func msplitDIA(file afero.File) ([]types.Peptide, map[string]string) {
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.LazyQuotes = true

	// Skip header.
	_, err := reader.Read()
	if err != nil {
		log.Fatalln(err)
	}

	peptideMap := make(map[string]string, 0)
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
		peptideMap[line[4]] = sequence
	}

	return peptides, peptideMap
}
