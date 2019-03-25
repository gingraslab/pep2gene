// Package typedef contains type definitions used throughout gene-peptide
package typedef

// Parameters for command line arguments.
type Parameters struct {
	Database           string
	FDR                float64
	File               string
	MapFile            string
	PeptideProbability float64
	Pipeline           string
}

// Peptide contains the amino acid "Sequence" for a peptide, the "Modified" version of
// the peptide and whether it is "Decoy"
type Peptide struct {
	Decoy    bool
	Modified string
	Sequence string
}

// PeptideStat contains the spectral count for a peptide and the individual counts
// for its modified forms
type PeptideStat struct {
	Count    int
	Modified map[string]int
}

// Protein contains the protein name, gene ID and sequence for a protein
type Protein struct {
	GeneID   string
	GeneName string
	GI       string
	Name     string
	Sequence string
}

// SpectralCounts is a map of peptides to their spectral counts and modified forms
type SpectralCounts map[string]*PeptideStat
