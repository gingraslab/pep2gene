// Package stats generates summary statistics
package stats

import "github.com/knightjdr/gene-peptide/typedef"

// QuantifyPeptides sums the spectral counts for peptides in both raw and modified forms
func QuantifyPeptides(peptides []typedef.Peptide) typedef.SpectralCounts {
	spectralCounts := make(typedef.SpectralCounts)
	for _, peptide := range peptides {
		if _, ok := spectralCounts[peptide.Sequence]; ok {
			spectralCounts[peptide.Sequence].Count = spectralCounts[peptide.Sequence].Count + 1
			if _, ok := spectralCounts[peptide.Sequence].Modified[peptide.Modified]; ok {
				spectralCounts[peptide.Sequence].Modified[peptide.Modified]++
			} else {
				spectralCounts[peptide.Sequence].Modified[peptide.Modified] = 1
			}
		} else {
			spectralCounts[peptide.Sequence] = &typedef.PeptideStat{
				Count:    1,
				Modified: map[string]int{peptide.Modified: 1},
			}
		}
	}
	return spectralCounts
}
