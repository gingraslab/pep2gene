// Package digestion contains functions for digesting a sequence to peptides
package digestion

// Digest processes an amino acid sequence into peptides
func Digest(sequence, enzyme string, missedClevages int) map[string]bool {
	var re string
	var terminus string
	switch enzyme {
	case "arg-c":
		re = "([R])[^P]"
		terminus = "c"
	case "asp-n":
		re = "([BD])"
		terminus = "n"
	case "asp-n_ambic":
		re = "([DE])"
		terminus = "n"
	case "chymotrypsin":
		re = "([FYWL])[^P]"
		terminus = "c"
	case "cnbr":
		re = "([M])"
		terminus = "c"
	case "lys-c":
		re = "([K])[^P]"
		terminus = "c"
	case "lys-c/p":
		re = "([K])"
		terminus = "c"
	case "lys-n":
		re = "([K])"
		terminus = "n"
	case "pepsina":
		re = "([FL]"
		terminus = "c"
	case "trypchymo":
		re = "([FYWLKR])[^P]"
		terminus = "c"
	case "trypsin":
		re = "([KR])[^P]"
		terminus = "c"
	case "trypsin/p":
		re = "([KR])"
		terminus = "c"
	case "v8-de":
		re = "([BDEZ])[^P]"
		terminus = "c"
	case "v8-e":
		re = "([EZ])[^P]"
		terminus = "c"
	default:
		re = "([KR])[^P]"
		terminus = "c"
	}
	return cut(sequence, re, terminus, missedClevages)
}
