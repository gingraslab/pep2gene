package read

import (
	"bufio"
	"log"
	"regexp"
	"strings"

	"github.com/knightjdr/gene-peptide/fs"
	"github.com/knightjdr/gene-peptide/types"
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

func sequenceNames(line string) map[string]string {
	fullEntry, _ := regexp.Compile("^>..\\|([\\d\\w_-]+)\\|gn\\|(\\w+):(\\d+)\\| (.+) \\[")
	geneIDEntry, _ := regexp.Compile("^>..\\|([\\d\\w_-]+)\\|gn\\|(\\w+):(\\d+)")
	ascessionEntry, _ := regexp.Compile("^>..\\|([\\d\\w_-]+)\\|")
	if fullEntry.MatchString(line) {
		matches := fullEntry.FindStringSubmatch(line)
		return map[string]string{
			"geneid":   matches[3],
			"genename": matches[2],
			"gi":       matches[1],
			"name":     matches[4],
		}
	} else if geneIDEntry.MatchString(line) {
		matches := geneIDEntry.FindStringSubmatch(line)
		return map[string]string{
			"geneid":   matches[3],
			"genename": matches[2],
			"gi":       matches[1],
			"name":     matches[2],
		}
	} else if ascessionEntry.MatchString(line) {
		matches := ascessionEntry.FindStringSubmatch(line)
		return map[string]string{
			"geneid":   matches[1],
			"genename": matches[1],
			"gi":       matches[1],
			"name":     matches[1],
		}
	}
	id := strings.Split(line, " ")[0]
	id = id[1:]
	return map[string]string{
		"geneid":   id,
		"genename": id,
		"gi":       id,
		"name":     id,
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
			currProtein.GI = names["gi"]
			currProtein.Name = names["name"]
			currProtein.Sequence = ""
			geneMap[names["geneid"]] = names["genename"]
		} else {
			sequence.WriteString(line)
		}
	}
	proteins = appendDatabase(proteins, currProtein, &sequence)

	return proteins, geneMap
}
