// Package types contains type definitions used throughout gene-peptide
package types

// Gene contains peptides matched to genes, genes with shared peptides and subsumed genes
type Gene struct {
	IsSubsumed bool
	Peptides   []string
	Shared     []string
	Subsumed   []string
}

// Genes is a map of gene IDS to their peptide and gene info
type Genes map[string]*Gene

// Parameters for command line arguments.
type Parameters struct {
	Database           string
	Enzyme             string
	FDR                float64
	File               string
	MapFile            string
	MissedCleavages    int
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
	Genes    []string
	Modified map[string]int
}

// Protein contains the protein name, gene ID and sequence for a protein
type Protein struct {
	GeneID          string
	GeneName        string
	GI              string
	MatchedPeptides []string
	Name            string
	Sequence        string
}

// Peptides is a map of peptides to their spectral counts and modified forms
type Peptides map[string]*PeptideStat
