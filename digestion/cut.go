package digestion

import (
	"regexp"
	"strings"
)

func cut(sequence, terminus string, cutRegex, afterRegex *regexp.Regexp, missedCleavages int) map[string]bool {
	possibleMatches := cutRegex.FindAllStringIndex(sequence, -1)
	matches := make([]int, 0)
	for i := range possibleMatches {
		matchIndex := possibleMatches[i][0]
		trailingSequence := sequence[matchIndex+1:]
		if !afterRegex.MatchString(trailingSequence) {
			matches = append(matches, matchIndex)
		}
	}

	// Generate peptides
	singleCuts := extractPeptides(sequence, terminus, matches)

	numCuts := len(singleCuts)
	digest := make(map[string]bool)
	for i := 0; i <= missedCleavages; i++ {
		lastPeptideIndex := numCuts - i
		for j := 0; j < lastPeptideIndex; j++ {
			missCleavedPeptide := strings.Join(singleCuts[j:j+i+1], "")
			digest[missCleavedPeptide] = true
		}
	}

	return digest
}
