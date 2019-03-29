package output

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/knightjdr/gene-peptide/types"
	"github.com/spf13/afero"
)

func writePeptide(file afero.File, sep rune, peptideCount map[string]float64) {
	sequences := make([]string, 0)
	for sequence := range peptideCount {
		sequences = append(sequences, sequence)
	}
	sort.Strings(sequences)

	for _, sequence := range sequences {
		spectralCount := int(math.Round(peptideCount[sequence]))
		file.WriteString(fmt.Sprintf("%[2]s%[1]s%[3]d%[1]s\n", string(sep), sequence, spectralCount))
	}
}

func writeGene(file afero.File, sep rune, index int, gene string, details *types.Gene, geneMap, geneIDtoName map[string]string) {
	geneNames := make([]string, 0)
	geneNames = append(geneNames, gene)

	geneID := geneMap[gene]
	geneIDs := make([]string, 0)
	geneIDs = append(geneIDs, geneID)

	var unique int
	if len(details.Shared) > 0 {
		// Get gene names for shared gene IDs, sort and add to gene name list
		sharedGenes := make([]string, len(details.Shared))
		for i, gene := range details.Shared {
			sharedGenes[i] = geneIDtoName[gene]
		}
		sort.Strings(sharedGenes)
		geneNames = append(geneNames, sharedGenes...)

		// Get gene IDs for sorted shared genes
		for _, sharedGene := range sharedGenes {
			geneIDs = append(geneIDs, geneMap[sharedGene])
		}
		unique = details.UniqueShared
	} else {
		unique = int(math.Round(details.Unique))
	}

	geneNameString := strings.Join(geneNames, ", ")
	geneIDString := strings.Join(geneIDs, ", ")
	spectralCount := int(math.Round(details.Count))
	subsumedString := strings.Join(details.Subsumed, ", ")

	file.WriteString(fmt.Sprintf("\nHit_%[2]d%[1]s%[3]s%[1]s%[4]s%[1]s%[5]d%[1]s%[6]d%[1]s%[7]s%[1]s\n", string(sep), index, geneNameString, geneIDString, spectralCount, unique, subsumedString))
	writePeptide(file, sep, details.PeptideCount)
}
