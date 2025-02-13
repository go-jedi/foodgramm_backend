package utils

import "testing"

func TestIsValidURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "valid HTTPS URL",
			input:    "https://www.example.com",
			expected: true,
		},
		{
			name:     "valid HTTP URL",
			input:    "http://example.com",
			expected: true,
		},
		{
			name:     "valid FTP URL",
			input:    "ftp://example.com/resource.txt",
			expected: true,
		},
		{
			name:     "invalid URL without scheme",
			input:    "example.com",
			expected: false,
		},
		{
			name:     "invalid URL with missing host",
			input:    "https://",
			expected: false,
		},
		{
			name:     "invalid URL with query and fragment",
			input:    "https://example.com/path?query=param#fragment",
			expected: true,
		},
		{
			name:     "invalid URL with only scheme",
			input:    "http://",
			expected: false,
		},
		{
			name:     "invalid URL with relative path",
			input:    "/path/to/resource",
			expected: false,
		},
		{
			name:     "invalid URL with local file",
			input:    "file:///path/to/file",
			expected: false,
		},
		{
			name:     "invalid URL with mailto",
			input:    "mailto:user@example.com",
			expected: false,
		},
		{
			name:     "invalid URL with malformed host",
			input:    "http://example..com",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsValidURL(test.input)

			if result != test.expected {
				t.Errorf("For input '%s', expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}
