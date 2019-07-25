package output

import (
	"math"
	"path/filepath"
	"sort"
	"strings"

	"github.com/knightjdr/gene-peptide/helpers"
	"github.com/knightjdr/gene-peptide/types"
)

// Data is a map of genes to their output summary.
type Data struct {
	Database           string           `json:"database"`
	Enzyme             string           `json:"enzyme,omitempty"`
	MissedCleavages    int              `json:"missedCleavages,omitempty"`
	FDR                float64          `fdr:"fdr,omitempty"`
	File               string           `json:"file"`
	Genes              map[string]*Gene `json:"genes"`
	PeptideProbability float64          `json:"peptideProbability,omitempty"`
	Pipeline           string           `json:"pipeline"`
}

// Gene is summary of a gene for output.
type Gene struct {
	Name          string             `json:"name"`
	Peptides      map[string]Peptide `json:"peptides"`
	SharedIDs     string             `json:"sharedIDs"`
	SharedNames   string             `json:"sharedNames"`
	SpectralCount float64            `json:"spectralCount"`
	Subsumed      string             `json:"subsumed"`
	Unique        int                `json:"unique"`
}

// Peptide is summary of a gene's peptides for output.
type Peptide struct {
	AllottedSpectralCount float64 `json:"allottedSpectralCount"`
	TotalSpectralCount    int     `json:"totalSpectralCount"`
	Unique                bool    `json:"unique"`
}

func summarizePeptides(
	genes []string,
	peptideCount map[string]float64,
	peptides types.Peptides,
	peptideMap map[string]string,
) map[string]Peptide {
	summary := make(map[string]Peptide, 0)
	for peptide, spectralCount := range peptideCount {
		peptideStats := peptides[peptideMap[peptide]]
		unique := false
		if len(genes) == 1 {
			unique = peptideStats.Unique
		} else if len(genes) > 1 && helpers.SliceEqual(genes, peptideStats.Genes) {
			unique = true
		}
		newPeptide := Peptide{
			AllottedSpectralCount: spectralCount,
			TotalSpectralCount:    peptideStats.Modified[peptide],
			Unique:                unique,
		}
		summary[peptide] = newPeptide
	}

	return summary
}

// Format data for output.
func Format(
	options types.Parameters,
	genes types.Genes,
	geneIDtoName map[string]string,
	peptides types.Peptides,
	peptideMap map[string]string,
) Data {
	summary := Data{
		Database:           filepath.Base(options.Database),
		Enzyme:             options.Enzyme,
		MissedCleavages:    options.MissedCleavages,
		FDR:                options.FDR,
		File:               filepath.Base(options.File),
		Genes:              make(map[string]*Gene, 0),
		PeptideProbability: options.PeptideProbability,
		Pipeline:           options.Pipeline,
	}

	for geneID, details := range genes {
		// For each gene get a list of IDs and names for any genes with shared peptides.
		sharedIDs := make([]string, 0)
		sharedNames := make([]string, len(details.Shared))

		var unique int
		if len(details.Shared) > 0 {
			// Get gene names for shared gene IDs, sort and add to gene name list.
			nameMap := make(map[string]string, 0)
			for i, gene := range details.Shared {
				sharedNames[i] = geneIDtoName[gene]
				nameMap[geneIDtoName[gene]] = gene
			}
			sort.Strings(sharedNames)

			// Get gene IDs for sorted shared genes
			for _, sharedName := range sharedNames {
				sharedIDs = append(sharedIDs, nameMap[sharedName])
			}
			unique = details.UniqueShared
		} else {
			unique = int(math.Round(details.Unique))
		}

		peptideDetails := summarizePeptides(append(details.Shared, geneID), details.PeptideCount, peptides, peptideMap)
		sort.Strings(details.Subsumed)
		subsumedString := strings.Join(details.Subsumed, ", ")

		summary.Genes[geneID] = &Gene{
			Name:          geneIDtoName[geneID],
			Peptides:      peptideDetails,
			SharedIDs:     strings.Join(sharedIDs, ", "),
			SharedNames:   strings.Join(sharedNames, ", "),
			SpectralCount: details.Count,
			Subsumed:      subsumedString,
			Unique:        unique,
		}
	}

	return summary
}
