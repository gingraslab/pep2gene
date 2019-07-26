package output

import (
	"fmt"
	"sort"

	"github.com/spf13/afero"
)

func fileHeader(file afero.File) {
	file.WriteString("HitNumber;;Gene;;GeneID;;SpectralCount;;Unique;;Subsumed\n")
	file.WriteString("Peptide;;TotalSpectralCount;;IsUnique\n")
}

func writeTXT(file afero.File, genes map[string]*Gene) {
	fileHeader(file)

	// Determine output order for genes (alphabetical by name).
	geneOrder := make([]string, 0)
	nameToID := make(map[string]string, 0)
	for geneID, details := range genes {
		geneOrder = append(geneOrder, details.Name)
		nameToID[details.Name] = geneID
	}
	sort.Strings(geneOrder)

	for i, geneName := range geneOrder {
		// Output gene summary.
		geneID := nameToID[geneName]
		details := genes[geneID]
		sharedIDs := geneID
		if details.SharedIDs != "" {
			sharedIDs = fmt.Sprintf("%s, %s", sharedIDs, details.SharedIDs)
		}
		sharedNames := details.Name
		if details.SharedNames != "" {
			sharedNames = fmt.Sprintf("%s, %s", sharedNames, details.SharedNames)
		}
		file.WriteString(
			fmt.Sprintf("\nHit_%d;;%s;;%s;;%.2f;;%d;;%s\n",
				i+1,
				sharedNames,
				sharedIDs,
				details.SpectralCount,
				details.Unique,
				details.Subsumed,
			),
		)

		// Output peptide details.
		sequences := make([]string, 0)
		for sequence := range details.Peptides {
			sequences = append(sequences, sequence)
		}
		sort.Strings(sequences)

		for _, sequence := range sequences {
			peptideDetails := details.Peptides[sequence]
			unique := "no"
			if peptideDetails.Unique {
				unique = "yes"
			}
			file.WriteString(
				fmt.Sprintf(
					"%[1]s;;%[2]d;;%[3]s\n",
					sequence,
					peptideDetails.TotalSpectralCount,
					unique,
				),
			)
		}
	}
}
