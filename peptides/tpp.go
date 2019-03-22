package peptides

import (
	"bufio"
	"log"
	"regexp"
	"strconv"

	"github.com/spf13/afero"
)

// Peptide contains the amino acid "Sequence" for a peptide, the "Modified" version of
// the peptide and whether it is "Decoy"
type Peptide struct {
	Decoy    bool
	Modified string
	Sequence string
}

func tpp(file afero.File, peptidProbabilty float64) []Peptide {
	scanner := bufio.NewScanner(file)

	decoyRegex, _ := regexp.Compile("^<alternative_protein protein=\"DECOY")
	modifiedRegex, _ := regexp.Compile("^<modification_info modified_peptide=\"([^\"]+)\"")
	peptideRegex, _ := regexp.Compile("^<search_hit hit_rank=\"\\d+\" peptide=\"([^\"]+)\"")
	propabilityRegex, _ := regexp.Compile("^<peptideprophet_result probability=\"([0-9\\.]+)")

	peptides := make([]Peptide, 0)
	var currPeptide Peptide
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
			if prob >= peptidProbabilty && !currPeptide.Decoy {
				peptides = append(peptides, currPeptide)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	return peptides
}
