package check

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyService_ProcessStatus(t *testing.T) {
	tests := []struct {
		name        string
		status      string
		expected    string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "Valid status",
			status:      "ok",
			expected:    "Success",
			expectError: false,
		},
		{
			name:        "Not found error",
			status:      "not_found",
			expected:    "",
			expectError: true,
			errorMsg:    "error: resource not found",
		},
		{
			name:        "Timeout error",
			status:      "timeout",
			expected:    "",
			expectError: true,
			errorMsg:    "error: request timeout",
		},
		{
			name:        "Unknown status",
			status:      "invalid",
			expected:    "",
			expectError: true,
			errorMsg:    fmt.Sprintf("error: unknown status %q", "invalid"),
		},
	}

	service := &MyService{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.ProcessStatus(tt.status)

			if tt.expectError {
				assert.Error(t, err)
				assert.Equal(t, tt.errorMsg, err.Error()) // Check exact error message
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
