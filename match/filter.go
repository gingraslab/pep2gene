package match

import "github.com/knightjdr/gene-peptide/types"

// Filter removes all genes from a peptide's list of matched genes that
// have no evidence of existence (subsumed genes).
func Filter(peptides types.Peptides, genes types.Genes) types.Peptides {
	for peptide := range peptides {
		keepGenes := make([]string, 0)
		for _, gene := range peptides[peptide].Genes {
			if _, ok := genes[gene]; ok {
				keepGenes = append(keepGenes, gene)
			}
		}
		peptides[peptide].Genes = keepGenes
	}
	return peptides
}
