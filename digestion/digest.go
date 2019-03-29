// Package digestion contains functions for digesting a sequence to peptides
package digestion

import (
	"fmt"
	"regexp"
)

// Digest processes an amino acid sequence into peptides
func Digest(sequence, enzyme string, missedClevages int) map[string]bool {
	var match string
	var afterExclude string
	var terminus string
	switch enzyme {
	case "arg-c":
		match = "R"
		afterExclude = "P"
		terminus = "c"
	case "asp-n":
		match = "BD"
		terminus = "n"
	case "asp-n_ambic":
		match = "DE"
		terminus = "n"
	case "chymotrypsin":
		match = "FYWL"
		afterExclude = "P"
		terminus = "c"
	case "cnbr":
		match = "M"
		terminus = "c"
	case "lys-c":
		match = "K"
		afterExclude = "P"
		terminus = "c"
	case "lys-c/p":
		match = "K"
		terminus = "c"
	case "lys-n":
		match = "K"
		terminus = "n"
	case "pepsina":
		match = "FL]"
		terminus = "c"
	case "trypchymo":
		match = "FYWLKR"
		afterExclude = "P"
		terminus = "c"
	case "trypsin":
		match = "KR"
		afterExclude = "P"
		terminus = "c"
	case "trypsin/p":
		match = "KR"
		terminus = "c"
	case "v8-de":
		match = "BDEZ"
		afterExclude = "P"
		terminus = "c"
	case "v8-e":
		match = "EZ"
		afterExclude = "P"
		terminus = "c"
	default:
		match = "KR"
		afterExclude = "P"
		terminus = "c"
	}

	afterRegex, _ := regexp.Compile(fmt.Sprintf("^%s", afterExclude))
	cutRegex, _ := regexp.Compile(fmt.Sprintf("([%s])", match))

	return cut(sequence, terminus, cutRegex, afterRegex, missedClevages)
}
