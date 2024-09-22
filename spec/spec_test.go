package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGoFile(t *testing.T) {
	got := NewGofile("./test_data/Gofile.yaml")

	assert.Equal(t, "github.com/owner/repo", got.GitRepo)
	assert.Equal(t, "./cmd/server", got.Path)
	assert.Equal(t, "main", got.GitRef)
}

func TestIsValid(t *testing.T) {
	goFile := NewGofile("./test_data/Gofile.yaml")

	valid, msg := goFile.IsValid()

	assert.True(t, valid)
	assert.Equal(t, "", msg)
}
