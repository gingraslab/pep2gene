package digestion

import (
	"regexp"
	"strings"
)

func cut(sequence, re, terminus string, missedCleavages int) map[string]bool {
	cutRegex, _ := regexp.Compile(re)
	matches := cutRegex.FindAllStringIndex(sequence, -1)

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
