// Package output writes results to a file.
package output

import (
	"fmt"
	"log"

	"github.com/knightjdr/gene-peptide/fs"
	"github.com/knightjdr/gene-peptide/helpers"
)

func outFileName(filePath, format string) string {
	filename := helpers.Filename(filePath)
	return fmt.Sprintf("%s.%s", filename, format)
}

// Write gene results to a file.
func Write(
	filePath,
	format string,
	outputData Data,
) {
	// Open file for writing.
	outfile, err := fs.Instance.Create(outFileName(filePath, format))
	if err != nil {
		log.Fatalln(err)
	}
	defer outfile.Close()

	switch format {
	case "json":
		writeJSON(outfile, outputData)
	case "txt":
		writeTXT(outfile, outputData.Genes)
	}
}
