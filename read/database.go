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

func appendDatabase(proteins []types.Protein, currProtein types.Protein, sequence *strings.Builder) []types.Protein {
	str := sequence.String()
	if str != "" {
		currProtein.Sequence = str
		sequence.Reset()
		return append(proteins, currProtein)
	}
	return proteins
}

func proteinEntry(id string) string {
	return fmt.Sprintf("p-%s", id)
}

func sequenceNames(line string) map[string]string {
	geneIDEntry, _ := regexp.Compile("^>..\\|[\\d\\w_-]+\\|gn\\|([\\w-]+):(\\d+)")
	ascessionEntry, _ := regexp.Compile("^>..\\|([\\d\\w_-]+)\\|")
	if geneIDEntry.MatchString(line) {
		matches := geneIDEntry.FindStringSubmatch(line)
		return map[string]string{
			"geneid":   matches[2],
			"genename": matches[1],
		}
	} else if ascessionEntry.MatchString(line) {
		matches := ascessionEntry.FindStringSubmatch(line)
		id := proteinEntry(matches[1])
		return map[string]string{
			"geneid":   id,
			"genename": id,
		}
	}
	id := strings.Split(line, " ")[0]
	id = proteinEntry(id[1:])
	return map[string]string{
		"geneid":   id,
		"genename": id,
	}
}

// Database reads a fasta database and generates a gene ID to gene name map
func Database(filename string) ([]types.Protein, map[string]string) {
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
			proteins = appendDatabase(proteins, currProtein, &sequence)
			names := sequenceNames(line)
			currProtein.GeneID = names["geneid"]
			currProtein.GeneName = names["genename"]
			currProtein.Sequence = ""
			geneMap[names["geneid"]] = names["genename"]
		} else {
			sequence.WriteString(line)
		}
	}
	proteins = appendDatabase(proteins, currProtein, &sequence)

	return proteins, geneMap
}
