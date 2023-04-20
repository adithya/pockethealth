package storageutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateBlobUUID(t *testing.T) {
	assert.NotEmpty(t, GenerateBlobUUID())
}
