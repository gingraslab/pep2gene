package read

import (
	"bufio"
	"log"
	"regexp"
	"strconv"

	"github.com/knightjdr/pep2gene/types"
	"github.com/spf13/afero"
)

func tpp(file afero.File, peptideProbabilty float64, inferEnzyme bool) ([]types.Peptide, map[string]string, string) {
	enzymeRegex, _ := regexp.Compile("^<sample_enzyme name=\"([\\w-/]+)\"")
	modifiedRegex, _ := regexp.Compile("^<modification_info modified_peptide=\"([^\"]+)\"")
	peptideRegex, _ := regexp.Compile("^<search_hit hit_rank=\"\\d+\" peptide=\"([^\"]+)\"")
	propabilityRegex, _ := regexp.Compile("^<peptideprophet_result probability=\"([0-9\\.]+)")

	enzyme := ""
	peptideMap := make(map[string]string, 0)
	peptides := make([]types.Peptide, 0)
	var currPeptide types.Peptide

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if peptideRegex.MatchString(line) {
			peptidesMatches := peptideRegex.FindStringSubmatch(line)
			currPeptide.Modified = peptidesMatches[1]
			currPeptide.Sequence = peptidesMatches[1]
		} else if modifiedRegex.MatchString(line) {
			modifiedMatches := modifiedRegex.FindStringSubmatch(line)
			currPeptide.Modified = modifiedMatches[1]
		} else if propabilityRegex.MatchString(line) {
			probabilityMatches := propabilityRegex.FindStringSubmatch(line)
			prob, _ := strconv.ParseFloat(probabilityMatches[1], 64)
			if prob >= peptideProbabilty {
				peptides = append(peptides, currPeptide)
				peptideMap[currPeptide.Modified] = currPeptide.Sequence
			}
		} else if inferEnzyme && enzymeRegex.MatchString(line) {
			enzymeMatches := enzymeRegex.FindStringSubmatch(line)
			enzyme = enzymeMatches[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return peptides, peptideMap, enzyme
}
