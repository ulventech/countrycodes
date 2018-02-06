package countrycodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetName(t *testing.T) {
	assert.Equal(t, "Norway", Name("no"))
	assert.Equal(t, "Norway", Name("NO"))
	assert.Equal(t, "", Name("non-existant"))
	assert.Equal(t, "", Name(""))
}
