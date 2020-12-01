package read

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/knightjdr/pep2gene/types"
	"github.com/spf13/afero"
)

// passesOpenSwathFilters checks if a peptide passes supplied filters.
func passesOpenSwathFilters(line []string, options types.Parameters) bool {
	isDecoy, _ := strconv.Atoi(line[1])
	if options.IgnoreDecoys && isDecoy == 1 {
		return false
	}

	mScore, _ := strconv.ParseFloat(line[22], 64)
	mScorePeptideExperimentWide, _ := strconv.ParseFloat(line[27], 64)
	peakGroupRank, _ := strconv.Atoi(line[20])

	if mScore <= options.Mscore &&
		mScorePeptideExperimentWide <= options.MscorePeptideExperimentWide &&
		peakGroupRank <= options.PeakGroupRank {
		return true
	}

	return false
}

func openswath(file afero.File, options types.Parameters) ([]types.Peptide, map[string]string) {
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

		if passesOpenSwathFilters(line, options) {
			intensity, _ := strconv.ParseFloat(line[15], 64)
			modPeptide := line[12]
			sequence := line[11]
			peptides = append(peptides, types.Peptide{
				Intensity: intensity,
				Modified:  modPeptide,
				Sequence:  sequence},
			)
			peptideMap[modPeptide] = sequence
		}
	}

	return peptides, peptideMap
}
