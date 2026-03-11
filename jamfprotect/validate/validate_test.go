package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOneOf(t *testing.T) {
	tests := []struct {
		name      string
		fieldName string
		value     string
		allowed   []string
		wantErr   bool
	}{
		{
			name:      "valid value",
			fieldName: "status",
			value:     "active",
			allowed:   []string{"active", "inactive", "pending"},
			wantErr:   false,
		},
		{
			name:      "empty value allowed",
			fieldName: "status",
			value:     "",
			allowed:   []string{"active", "inactive"},
			wantErr:   false,
		},
		{
			name:      "invalid value",
			fieldName: "status",
			value:     "invalid",
			allowed:   []string{"active", "inactive"},
			wantErr:   true,
		},
		{
			name:      "case sensitive",
			fieldName: "type",
			value:     "Active",
			allowed:   []string{"active", "inactive"},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := OneOf(tt.fieldName, tt.value, tt.allowed...)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.fieldName)
				assert.Contains(t, err.Error(), tt.value)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestIntBetween(t *testing.T) {
	tests := []struct {
		name      string
		fieldName string
		value     int
		min       int
		max       int
		wantErr   bool
	}{
		{
			name:      "value within range",
			fieldName: "level",
			value:     5,
			min:       0,
			max:       10,
			wantErr:   false,
		},
		{
			name:      "value at min",
			fieldName: "level",
			value:     0,
			min:       0,
			max:       10,
			wantErr:   false,
		},
		{
			name:      "value at max",
			fieldName: "level",
			value:     10,
			min:       0,
			max:       10,
			wantErr:   false,
		},
		{
			name:      "value below min",
			fieldName: "level",
			value:     -1,
			min:       0,
			max:       10,
			wantErr:   true,
		},
		{
			name:      "value above max",
			fieldName: "level",
			value:     11,
			min:       0,
			max:       10,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := IntBetween(tt.fieldName, tt.value, tt.min, tt.max)
			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.fieldName)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
