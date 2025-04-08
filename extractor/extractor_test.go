package extractor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractorRun(t *testing.T) {
	ex := &Extractor{}
	result, err := ex.Run(&RunParams{})
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
