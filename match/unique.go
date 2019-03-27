package match

import "github.com/knightjdr/gene-peptide/types"

// Unique counts the number of unique peptides per fene
func Unique(peptides types.Peptides, genes types.Genes) types.Genes {
	for gene := range genes {
		genes[gene].Unique = 0
		for _, peptide := range genes[gene].Peptides {
			if len(peptides[peptide].Genes) == 1 {
				genes[gene].Unique++
			}
		}
	}
	return genes
}
