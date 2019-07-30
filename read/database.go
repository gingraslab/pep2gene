package read

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/knightjdr/pep2gene/fs"
	"github.com/knightjdr/pep2gene/types"
)

func appendDatabase(
	proteins []types.Protein,
	currProtein types.Protein,
	sequence *strings.Builder,
	geneMap map[string]string,
	ignoreInvalid bool,
) ([]types.Protein, map[string]string) {
	str := sequence.String()
	if str != "" && (!ignoreInvalid || (ignoreInvalid && currProtein.Valid)) {
		currProtein.Sequence = str
		sequence.Reset()
		updatedDB := append(proteins, currProtein)
		geneMap[currProtein.GeneID] = currProtein.GeneName
		return updatedDB, geneMap
	}
	return proteins, geneMap
}

func proteinEntry(id string) string {
	return fmt.Sprintf("p-%s", id)
}

func sequenceNames(line string) (map[string]string, bool) {
	geneIDEntry, _ := regexp.Compile("^>..\\|[\\d\\w_-]+\\|gn\\|([\\w-]+):(\\d+)")
	if geneIDEntry.MatchString(line) {
		matches := geneIDEntry.FindStringSubmatch(line)
		description := map[string]string{
			"geneid":   matches[2],
			"genename": matches[1],
		}
		return description, true
	}
	id := strings.Split(line, " ")[0]
	id = proteinEntry(id[1:])
	description := map[string]string{
		"geneid":   id,
		"genename": id,
	}
	return description, false
}

// Database reads a fasta database and generates a gene ID to gene name map
func Database(filename string, ignoreInvalid bool) ([]types.Protein, map[string]string) {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	geneMap := make(map[string]string, 0)
	proteins := make([]types.Protein, 0)
	var sequence strings.Builder
	scanner := bufio.NewScanner(file)
	var currProtein types.Protein
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			proteins, geneMap = appendDatabase(proteins, currProtein, &sequence, geneMap, ignoreInvalid)
			names, valid := sequenceNames(line)
			currProtein.GeneID = names["geneid"]
			currProtein.GeneName = names["genename"]
			currProtein.Sequence = ""
			currProtein.Valid = valid
		} else {
			sequence.WriteString(line)
		}
	}
	proteins, geneMap = appendDatabase(proteins, currProtein, &sequence, geneMap, ignoreInvalid)

	return proteins, geneMap
}
