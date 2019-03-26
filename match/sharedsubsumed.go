package match

import (
	"github.com/knightjdr/gene-peptide/helpers"
	"github.com/knightjdr/gene-peptide/types"
)

// SharedSubsumed calculates which genes share peptides and which are subsumed by others
func SharedSubsumed(genes types.Genes) types.Genes {
	for geneID := range genes {
		for comparedGene := range genes {
			if geneID != comparedGene {
				shared := helpers.SliceEqual(genes[geneID].Peptides, genes[comparedGene].Peptides)
				if shared {
					genes[geneID].Shared = append(genes[geneID].Shared, comparedGene)
				} else {
					subsumed := helpers.SliceContains(genes[comparedGene].Peptides, genes[geneID].Peptides)
					if subsumed {
						genes[geneID].IsSubsumed = subsumed
						genes[comparedGene].Subsumed = append(genes[comparedGene].Subsumed, genes[geneID].Subsumed...)
						genes[comparedGene].Subsumed = append(genes[comparedGene].Subsumed, geneID)
					}
				}
			}
		}
	}

	// Remove subsumed genes or remove duplicates from subsumed list
	for geneID := range genes {
		if genes[geneID].IsSubsumed {
			delete(genes, geneID)
		} else {
			genes[geneID].Subsumed = helpers.SliceUnique(genes[geneID].Subsumed)
		}
	}

	return genes
}
