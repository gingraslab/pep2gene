package match

import (
	"github.com/gingraslab/pep2gene/helpers"
	"github.com/gingraslab/pep2gene/types"
)

// SharedSubsumed calculates which genes share peptides and which are subsumed by others
func SharedSubsumed(genes types.Genes) types.Genes {
	// Copy genes
	updatedGenes := make(types.Genes, len(genes))
	for geneID := range genes {
		updatedGenes[geneID] = genes[geneID].Copy()
	}

	for geneID := range genes {
		for comparedGene := range genes {
			if geneID != comparedGene {
				shared := helpers.SliceEqual(updatedGenes[geneID].Peptides, updatedGenes[comparedGene].Peptides)
				if shared {
					updatedGenes[geneID].Shared = append(updatedGenes[geneID].Shared, comparedGene)
				} else {
					subsumed := helpers.SliceContains(updatedGenes[comparedGene].Peptides, updatedGenes[geneID].Peptides)
					if subsumed {
						updatedGenes[geneID].IsSubsumed = subsumed
						updatedGenes[comparedGene].Subsumed = append(updatedGenes[comparedGene].Subsumed, updatedGenes[geneID].Subsumed...)
						updatedGenes[comparedGene].Subsumed = append(updatedGenes[comparedGene].Subsumed, geneID)
					}
				}
			}
		}
	}

	// Remove subsumed genes or remove duplicates from subsumed list
	for geneID := range updatedGenes {
		if updatedGenes[geneID].IsSubsumed {
			delete(updatedGenes, geneID)
		} else {
			updatedGenes[geneID].Subsumed = helpers.SliceUnique(updatedGenes[geneID].Subsumed)
		}
	}

	return updatedGenes
}
