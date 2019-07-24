package match

import "github.com/knightjdr/gene-peptide/types"

// Filter removes all genes from a peptide's list of matched genes that
// have no evidence that they exist in the sample (subsumed genes).
func Filter(peptides types.Peptides, genes types.Genes) types.Peptides {
	filteredPeptides := make(types.Peptides, len(peptides))
	for peptide := range peptides {
		filteredPeptides[peptide] = peptides[peptide].Copy()
		keepGenes := make([]string, 0)
		for _, gene := range filteredPeptides[peptide].Genes {
			if _, ok := genes[gene]; ok {
				keepGenes = append(keepGenes, gene)
			}
		}
		filteredPeptides[peptide].Genes = keepGenes
	}
	return filteredPeptides
}
