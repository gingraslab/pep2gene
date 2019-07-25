package read

import (
	"encoding/csv"
	"io"
	"log"
	"regexp"
	"strconv"

	"github.com/knightjdr/gene-peptide/types"
	"github.com/spf13/afero"
)

// msplitDDASequence converts peptides of the form K.NQVAM+15.995NPTNTVFDAK.R to a
// a sequence stripped of leading and trailing cleavage sites
func msplitDDASequence(peptide string) string {
	pepLength := len(peptide)
	sequence := peptide[:pepLength-2]
	sequence = peptide[2 : pepLength-2]
	return sequence
}

// msplitDDARawSequence removes any modifications from a peptide and returns the raw
// amino acid sequence
func msplitDDARawSequence(peptide string) string {
	modRegex, _ := regexp.Compile("[+-.0-9]")
	sequence := modRegex.ReplaceAllString(peptide, "")
	return sequence
}

func msplitDDA(file afero.File, fdr float64) ([]types.Peptide, map[string]string) {
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

		peptideFDR, _ := strconv.ParseFloat(line[14], 64)
		if peptideFDR <= fdr {
			modPeptide := msplitDDASequence(line[7])
			sequence := msplitDDARawSequence(modPeptide)
			peptides = append(peptides, types.Peptide{Modified: modPeptide, Sequence: sequence})
			peptideMap[modPeptide] = sequence
		}
	}

	return peptides, peptideMap
}
