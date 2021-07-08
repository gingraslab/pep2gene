// Package stats generates summary statistics
package stats

import (
	"github.com/gingraslab/pep2gene/types"
)

func getCount(peptide types.Peptide) float64 {
	if peptide.Intensity != 0 {
		return peptide.Intensity
	}
	return 1
}

// QuantifyPeptides sums the spectral counts or intensities for peptides in both raw and modified forms.
func QuantifyPeptides(peptides []types.Peptide) types.Peptides {
	counts := make(types.Peptides)
	for _, peptide := range peptides {
		count := getCount(peptide)
		if _, ok := counts[peptide.Sequence]; ok {
			counts[peptide.Sequence].Count += count
			if _, ok := counts[peptide.Sequence].Modified[peptide.Modified]; ok {
				counts[peptide.Sequence].Modified[peptide.Modified] += count
			} else {
				counts[peptide.Sequence].Modified[peptide.Modified] = count
			}
		} else {
			counts[peptide.Sequence] = &types.PeptideStat{
				Count:    count,
				Modified: map[string]float64{peptide.Modified: count},
			}
		}
	}
	return counts
}
