// Package stats generates summary statistics
package stats

import (
	"github.com/knightjdr/gene-peptide/types"
)

// QuantifyPeptides sums the spectral counts for peptides in both raw and modified forms
func QuantifyPeptides(peptides []types.Peptide) types.Peptides {
	spectralCounts := make(types.Peptides)
	for _, peptide := range peptides {
		if _, ok := spectralCounts[peptide.Sequence]; ok {
			spectralCounts[peptide.Sequence].Count++
			if _, ok := spectralCounts[peptide.Sequence].Modified[peptide.Modified]; ok {
				spectralCounts[peptide.Sequence].Modified[peptide.Modified]++
			} else {
				spectralCounts[peptide.Sequence].Modified[peptide.Modified] = 1
			}
		} else {
			spectralCounts[peptide.Sequence] = &types.PeptideStat{
				Count:    1,
				Modified: map[string]int{peptide.Modified: 1},
			}
		}
	}
	return spectralCounts
}
