package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuoteTableName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No fullstop in table name",
			input:    "schema_table",
			expected: "schema_table",
		},
		{
			name:     "Fullstop in table name",
			input:    "schema.table",
			expected: "\"schema.table\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := QuoteTableName(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
