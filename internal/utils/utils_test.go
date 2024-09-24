package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{
			name:     "dot",
			path:     ".",
			expected: "./bin/server",
		},
		{
			name:     "dot and forward slash",
			path:     "./",
			expected: "./bin/server",
		},
		{
			name:     "dot and forward slash",
			path:     "./example/demo",
			expected: "../../bin/server",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ConstructPath(tc.path)
			assert.Equal(t, tc.expected, got)
		})
	}
}
