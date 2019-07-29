package digestion

// Enzymes is a dictionary of available digestive enzymes.
func Enzymes() map[string]bool {
	return map[string]bool{
		"arg-c":        true,
		"asp-n":        true,
		"asp-n_ambic":  true,
		"chymotrypsin": true,
		"cnbr":         true,
		"lys-c":        true,
		"lys-c/p":      true,
		"lys-n":        true,
		"pepsina":      true,
		"trypchymo":    true,
		"trypsin":      true,
		"trypsin/p":    true,
		"v8-de":        true,
		"v8-e":         true,
	}
}
