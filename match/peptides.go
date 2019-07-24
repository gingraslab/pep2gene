// Package match contains matching algorithms for peptides and proteins
package match

import (
	"strings"

	"github.com/knightjdr/gene-peptide/digestion"
	"github.com/knightjdr/gene-peptide/helpers"
	"github.com/knightjdr/gene-peptide/types"
)

func fullSequence(peptides types.Peptides, db []types.Protein) (types.Peptides, types.Genes) {
	matchedPeptides := make(types.Peptides, len(peptides))

	genes := make(types.Genes)
	for peptide := range peptides {
		matchedPeptides[peptide] = peptides[peptide].Copy()
		matchedPeptides[peptide].Genes = make([]string, 0)
		for _, entry := range db {
			if strings.Contains(entry.Sequence, peptide) {
				matchedPeptides[peptide].Genes = append(matchedPeptides[peptide].Genes, entry.GeneID)
				if _, ok := genes[entry.GeneID]; ok {
					genes[entry.GeneID].Peptides = append(genes[entry.GeneID].Peptides, peptide)
				} else {
					genes[entry.GeneID] = &types.Gene{
						Peptides: []string{peptide},
					}
				}
			}
		}
		if len(matchedPeptides[peptide].Genes) == 0 {
			delete(matchedPeptides, peptide)
		} else {
			matchedPeptides[peptide].Genes = helpers.SliceUnique(matchedPeptides[peptide].Genes)
			if len(matchedPeptides[peptide].Genes) == 1 {
				matchedPeptides[peptide].Unique = true
			}
		}
	}

	// Remove duplicates peptides in genes.
	for gene := range genes {
		genes[gene].Peptides = helpers.SliceUnique(genes[gene].Peptides)
	}
	return matchedPeptides, genes
}

func digestedSequence(peptides types.Peptides, db []types.Protein, enzyme string, missed int) (types.Peptides, types.Genes) {
	matchedPeptides := make(types.Peptides, len(peptides))

	// Allocate peptide map.
	for peptide := range peptides {
		matchedPeptides[peptide] = peptides[peptide].Copy()
		matchedPeptides[peptide].Genes = make([]string, 0)
	}

	genes := make(types.Genes)
	for _, entry := range db {
		digested := digestion.Digest(entry.Sequence, enzyme, missed)
		for peptide := range peptides {
			if _, ok := digested[peptide]; ok {
				matchedPeptides[peptide].Genes = append(matchedPeptides[peptide].Genes, entry.GeneID)
				if _, ok := genes[entry.GeneID]; ok {
					genes[entry.GeneID].Peptides = append(genes[entry.GeneID].Peptides, peptide)
				} else {
					genes[entry.GeneID] = &types.Gene{
						Peptides: []string{peptide},
					}
				}
			}
		}
		if _, ok := genes[entry.GeneID]; ok {
			genes[entry.GeneID].Peptides = helpers.SliceUnique(genes[entry.GeneID].Peptides)
		}
	}

	// Remove peptides with no matches and remove duplicate gene matches.
	for peptide := range peptides {
		if len(matchedPeptides[peptide].Genes) == 0 {
			delete(matchedPeptides, peptide)
		} else {
			matchedPeptides[peptide].Genes = helpers.SliceUnique(matchedPeptides[peptide].Genes)
			if len(matchedPeptides[peptide].Genes) == 1 {
				matchedPeptides[peptide].Unique = true
			}
		}
	}
	return matchedPeptides, genes
}

// Peptides finds proteins/genes that match to input peptides
func Peptides(peptides types.Peptides, db []types.Protein, enzyme string, missed int) (types.Peptides, types.Genes) {
	if enzyme == "" {
		return fullSequence(peptides, db)
	}
	return digestedSequence(peptides, db, enzyme, missed)
}
