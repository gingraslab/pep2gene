package match

import (
	"github.com/knightjdr/gene-peptide/helpers"
	"github.com/knightjdr/gene-peptide/types"
)

// Unique counts the number of unique peptides per gene. If a gene perfectly shares its
// peptides with other genes, then for each peptide in the group a gene will get a
// unique portion equal to 1 / (genes in group). Otherwise, peptides need to be specific
// to one gene to be unique.
func Unique(peptides types.Peptides, genes types.Genes) types.Genes {
	updatedGenes := make(types.Genes, len(genes))

	for gene := range genes {
		updatedGenes[gene] = genes[gene].Copy()
		updatedGenes[gene].Unique = 0
		if len(updatedGenes[gene].Shared) > 0 {
			updatedGenes[gene].UniqueShared = 0
			sharedGenes := updatedGenes[gene].Shared
			for _, peptide := range updatedGenes[gene].Peptides {
				sharedGeneNum := len(sharedGenes) + 1
				group := make([]string, sharedGeneNum)
				group = append(sharedGenes, gene)
				if helpers.SliceEqual(group, peptides[peptide].Genes) {
					uniquePortion := float64(1) / float64(sharedGeneNum)
					updatedGenes[gene].Unique += uniquePortion
					updatedGenes[gene].UniqueShared++
				}
			}
		} else {
			for _, peptide := range updatedGenes[gene].Peptides {
				if peptides[peptide].Unique {
					updatedGenes[gene].Unique += float64(1)
				}
			}
		}
	}
	return updatedGenes
}
