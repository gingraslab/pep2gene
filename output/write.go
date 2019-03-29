// Package output writes results to a file.
package output

import (
	"fmt"
	"log"
	"sort"

	"github.com/knightjdr/gene-peptide/fs"
	"github.com/knightjdr/gene-peptide/helpers"
	"github.com/knightjdr/gene-peptide/types"
	"github.com/spf13/afero"
)

var acceptedTypes = map[string]rune{
	"csv": ',',
	"dsv": ';',
	"tsv": '\t',
}

func outFileName(filePath, format string) string {
	filename := helpers.Filename(filePath)
	return fmt.Sprintf("%s.%s", filename, format)
}

func fileHeader(file afero.File, sep rune) {
	file.WriteString(fmt.Sprintf("HitNumber%[1]sGene%[1]sGeneID%[1]sSpectralCount%[1]sUnique%[1]sSubsumed\n", string(sep)))
	file.WriteString(fmt.Sprintf("Peptide%[1]sSpectralCount%[1]sIsUnique\n", string(sep)))
}

// Write gene results to a file.
func Write(filePath, format string, genes types.Genes, geneIDtoName map[string]string, peptides types.Peptides) {
	// Get map of genes to output.
	geneMap := make(map[string]string, 0)
	geneNames := make([]string, 0)
	for geneID := range genes {
		geneName := geneIDtoName[geneID]
		geneMap[geneName] = geneID
		geneNames = append(geneNames, geneName)
	}
	sort.Strings(geneNames)

	// Open file for writing.
	outfile, err := fs.Instance.Create(outFileName(filePath, format))
	if err != nil {
		log.Fatalln(err)
	}
	defer outfile.Close()

	sep := acceptedTypes[format]

	// Write header.
	fileHeader(outfile, sep)

	for i, gene := range geneNames {
		geneID := geneMap[gene]
		writeGene(outfile, sep, i+1, gene, genes[geneID], geneMap, geneIDtoName)
	}
}
