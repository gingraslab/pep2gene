// Package match contains matching algorithms for peptides and proteins
package match

import (
	"strings"

	"github.com/knightjdr/gene-peptide/helpers"

	"github.com/knightjdr/gene-peptide/types"
)

func fullSequence(peptides types.Peptides, db []types.Protein) (types.Peptides, types.Genes) {
	genes := make(types.Genes)
	for peptide := range peptides {
		peptides[peptide].Genes = make([]string, 0)
		for _, entry := range db {
			if strings.Contains(entry.Sequence, peptide) {
				peptides[peptide].Genes = append(peptides[peptide].Genes, entry.GeneID)
				if _, ok := genes[entry.GeneID]; ok {
					genes[entry.GeneID].Peptides = append(genes[entry.GeneID].Peptides, peptide)
				} else {
					genes[entry.GeneID] = &types.Gene{
						Peptides: []string{peptide},
					}
				}
			}
		}
		peptides[peptide].Genes = helpers.SliceUnique(peptides[peptide].Genes)
	}
	return peptides, genes
}

func trypticSequence(peptides types.Peptides, db []types.Protein, enzyme string, missed int) (types.Peptides, types.Genes) {
	genes := make(types.Genes)
	for peptide := range peptides {
		peptides[peptide].Genes = make([]string, 0)
	}
	return peptides, genes
}

// Peptides finds proteins/genes that match to input peptides
func Peptides(peptides types.Peptides, db []types.Protein, enzyme string, missed int) (types.Peptides, types.Genes) {
	if enzyme == "" {
		return fullSequence(peptides, db)
	}
	return trypticSequence(peptides, db, enzyme, missed)
}
