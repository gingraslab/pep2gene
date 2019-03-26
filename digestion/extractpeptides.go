package digestion

func extractPeptides(sequence, terminus string, matches [][]int) []string {
	// C-terminal cleavages should include the cut site
	includeCutAA := 0
	if terminus == "c" {
		includeCutAA = 1
	}

	peptides := make([]string, 0)
	lastCut := 0
	for i := range matches {
		cutSite := matches[i][0] + includeCutAA
		cleavedPeptide := sequence[lastCut:cutSite]
		peptides = append(peptides, cleavedPeptide)
		lastCut = cutSite
	}

	// Add trailing peptide
	if lastCut != len(sequence) {
		cleavedPeptide := sequence[lastCut:]
		peptides = append(peptides, cleavedPeptide)
	}
	return peptides
}
