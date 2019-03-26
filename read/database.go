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

// Database reads a fasta database
func Database(filename string) []types.Protein {
	file, err := fs.Instance.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	nameRegex, _ := regexp.Compile("^>gi\\|(\\d+)\\|gn\\|(\\w+):(\\d+)\\| (.+) \\[")

	var currProtein types.Protein
	proteins := make([]types.Protein, 0)
	var sequence strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if nameRegex.MatchString(line) {
			proteins = appendDatabase(proteins, currProtein, &sequence)
			nameMatches := nameRegex.FindStringSubmatch(line)
			currProtein.GeneID = nameMatches[3]
			currProtein.GeneName = nameMatches[2]
			currProtein.GI = nameMatches[1]
			currProtein.Name = nameMatches[4]
			currProtein.Sequence = ""
		} else {
			sequence.WriteString(line)
		}
	}
	proteins = appendDatabase(proteins, currProtein, &sequence)

	return proteins
}
