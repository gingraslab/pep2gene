package digestion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetEnzyme(t *testing.T) {
	// TEST1: inferred and valid enzyme.
	inferredEnzyme := "trypsin"
	suppliedEnzyme := ""
	assert.Equal(t, "trypsin", SetEnzyme(inferredEnzyme, suppliedEnzyme), "Should return inferred enzyme")

	// TEST2: inferred and invalid enzyme.
	inferredEnzyme = "abcde"
	suppliedEnzyme = ""
	assert.Equal(t, "", SetEnzyme(inferredEnzyme, suppliedEnzyme), "Should return nil string for invalid inferred enzyme")

	// TEST2: user supplied and valid enzyme.
	inferredEnzyme = ""
	suppliedEnzyme = "trypsin"
	assert.Equal(t, "trypsin", SetEnzyme(inferredEnzyme, suppliedEnzyme), "Should return supplied enzyme")

	// TEST2: user supplied and invalid enzyme.
	inferredEnzyme = ""
	suppliedEnzyme = "abcde"
	assert.Equal(t, "", SetEnzyme(inferredEnzyme, suppliedEnzyme), "Should return nil string for invalid supplied enzyme")
}
