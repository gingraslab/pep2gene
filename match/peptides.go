// Package match contains matching algorithms for peptides and proteins
package match

import (
	"strings"

	"github.com/knightjdr/gene-peptide/digestion"
	"github.com/knightjdr/gene-peptide/helpers"
	"github.com/knightjdr/gene-peptide/types"
)

// addPeptide intializes a gene entry if it doesn't exist add adds
// a peptide to its matches. If it does exist, it just appends the peptide.
func addPeptide(geneID string, genes types.Genes, peptide string) {
	if _, ok := genes[geneID]; ok {
		genes[geneID].Peptides = append(genes[geneID].Peptides, peptide)
	} else {
		genes[geneID] = &types.Gene{
			Peptides: []string{peptide},
		}
	}
}

// filterPeptides Remove speptides with no matches to a gene and removes
// duplicate gene matches from each peptide.
func filterPeptides(peptides types.Peptides, peptide string) {
	if len(peptides[peptide].Genes) == 0 {
		delete(peptides, peptide)
	} else {
		peptides[peptide].Genes = helpers.SliceUnique(peptides[peptide].Genes)
		if len(peptides[peptide].Genes) == 1 {
			peptides[peptide].Unique = true
		}
	}
}

func fullSequence(peptides types.Peptides, db []types.Protein) (types.Peptides, types.Genes) {
	matchedPeptides := make(types.Peptides, len(peptides))

	genes := make(types.Genes)
	for peptide := range peptides {
		matchedPeptides[peptide] = peptides[peptide].Copy()
		matchedPeptides[peptide].Genes = make([]string, 0)
		for _, entry := range db {
			if strings.Contains(entry.Sequence, peptide) {
				matchedPeptides[peptide].Genes = append(matchedPeptides[peptide].Genes, entry.GeneID)
				addPeptide(entry.GeneID, genes, peptide)
			}
		}
		filterPeptides(matchedPeptides, peptide)
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
				addPeptide(entry.GeneID, genes, peptide)
			}
		}

		// Remove duplicate peptides.
		if _, ok := genes[entry.GeneID]; ok {
			genes[entry.GeneID].Peptides = helpers.SliceUnique(genes[entry.GeneID].Peptides)
		}
	}

	// Remove peptides with no matches to a gene and remove duplicate gene matches.
	for peptide := range peptides {
		filterPeptides(matchedPeptides, peptide)
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
