package match

import "github.com/knightjdr/gene-peptide/types"

func assignModPeptides(modified map[string]int, counts map[string]float64, weight float64) map[string]float64 {
	updatedCounts := counts
	for modPeptide, modCount := range modified {
		updatedCounts[modPeptide] = weight * float64(modCount)
	}
	return updatedCounts
}

// Count spectra and unique peptides for each gene
func Count(peptides types.Peptides, genes types.Genes) types.Genes {
	// Allocate map for modified peptides
	for gene := range genes {
		genes[gene].PeptideCount = make(map[string]float64)
	}

	for peptide := range peptides {
		totalUnique := 0
		uniqueGenes := make([]string, 0)
		for _, gene := range peptides[peptide].Genes {
			if genes[gene].Unique > 0 {
				totalUnique += genes[gene].Unique
				uniqueGenes = append(uniqueGenes, gene)
			}
		}

		// For peptides shared between multiple genes, spectral counts are assigned
		// exclusively to those genes with unique peptides in proportion to the existing
		// evidence for those genes. If a peptide matches exclusively to genes that have
		// no unique peptides, then spectral counts are divided equally between the genes.
		if len(uniqueGenes) > 0 {
			for _, gene := range uniqueGenes {
				weight := float64(genes[gene].Unique) / float64(totalUnique)
				genes[gene].Count += weight * float64(peptides[peptide].Count)
				genes[gene].PeptideCount = assignModPeptides(peptides[peptide].Modified, genes[gene].PeptideCount, weight)
			}
		} else {
			weight := 1.0 / float64(len(peptides[peptide].Genes))
			for _, gene := range peptides[peptide].Genes {
				genes[gene].Count += weight * float64(peptides[peptide].Count)
				genes[gene].PeptideCount = assignModPeptides(peptides[peptide].Modified, genes[gene].PeptideCount, weight)
			}
		}
	}
	return genes
}
