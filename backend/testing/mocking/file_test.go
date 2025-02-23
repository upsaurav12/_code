package main

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadFile(t *testing.T) {
	tests := []struct {
		name        string
		fs          FileSystem
		expected    string
		expectedErr string
	}{
		{
			name:        "valid filename",
			fs:          mockFileSystem{openResult: io.NopCloser(strings.NewReader("test content"))},
			expected:    "test content",
			expectedErr: "",
		},
		{
			name:        "file not found",
			fs:          mockFileSystem{openError: errors.New("file not found")},
			expected:    "",
			expectedErr: "file not found",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			content, err := ReadFile(tc.fs, "test.txt")
			if tc.expectedErr != "" {
				require.EqualError(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, content)
			}
		})
	}
}
