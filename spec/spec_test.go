package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGoFile(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		expected *Gofile
	}{
		{
			name:     "local",
			filepath: "./test_data/local/Gofile.yaml",
			expected: &Gofile{
				Path:    "./cmd/server",
				Scratch: true,
			},
		},
		{
			name:     "remote",
			filepath: "./test_data/remote/Gofile.yaml",
			expected: &Gofile{
				GitRepo: "github.com/owner/repo",
				GitRef:  "main",
				Path:    "./cmd/server",
				Scratch: true,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NewGofile(tc.filepath)

			assert.Equal(t, tc.expected, got)
		})

	}
}
