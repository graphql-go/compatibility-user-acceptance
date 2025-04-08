package extractor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractorRun(t *testing.T) {
	ex := &Extractor{}
	params := RunParams{
		Organization:   "graphql-go",
		RepositoryName: "graphql",
	}
	result, err := ex.Run(&params)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
