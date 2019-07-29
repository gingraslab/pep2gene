package digestion

// SetEnzyme defines the enzyme, if any, to use for digestion.
func SetEnzyme(inferredEnzyme, suppliedEnzyme string) string {
	enzyme := ""
	if inferredEnzyme != "" {
		enzyme = inferredEnzyme
	} else if suppliedEnzyme != "" {
		enzyme = suppliedEnzyme
	}

	availableEnzymes := Enzymes()
	if _, ok := availableEnzymes[enzyme]; ok {
		return enzyme
	}
	return ""
}
