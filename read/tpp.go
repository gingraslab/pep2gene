package read

import (
	"bufio"
	"log"
	"regexp"
	"strconv"

	"github.com/knightjdr/gene-peptide/typedef"
	"github.com/spf13/afero"
)

func tpp(file afero.File, peptideProbabilty float64) []typedef.Peptide {
	decoyRegex, _ := regexp.Compile("^<alternative_protein protein=\"DECOY")
	modifiedRegex, _ := regexp.Compile("^<modification_info modified_peptide=\"([^\"]+)\"")
	peptideRegex, _ := regexp.Compile("^<search_hit hit_rank=\"\\d+\" peptide=\"([^\"]+)\"")
	propabilityRegex, _ := regexp.Compile("^<peptideprophet_result probability=\"([0-9\\.]+)")

	peptides := make([]typedef.Peptide, 0)
	var currPeptide typedef.Peptide

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if peptideRegex.MatchString(line) {
			peptidesMatches := peptideRegex.FindStringSubmatch(line)
			currPeptide.Decoy = false
			currPeptide.Modified = peptidesMatches[1]
			currPeptide.Sequence = peptidesMatches[1]
		} else if decoyRegex.MatchString(line) {
			currPeptide.Decoy = true
		} else if modifiedRegex.MatchString(line) {
			modifiedMatches := modifiedRegex.FindStringSubmatch(line)
			currPeptide.Modified = modifiedMatches[1]
		} else if propabilityRegex.MatchString(line) {
			probabilityMatches := propabilityRegex.FindStringSubmatch(line)
			prob, _ := strconv.ParseFloat(probabilityMatches[1], 64)
			if prob >= peptideProbabilty && !currPeptide.Decoy {
				peptides = append(peptides, currPeptide)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return peptides
}
