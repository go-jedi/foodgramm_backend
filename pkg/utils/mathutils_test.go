package utils

import (
	"math/big"
	"testing"
)

func TestStringToBigInt(t *testing.T) {
	tests := []struct {
		name      string
		inputStr  string
		base      int
		expected  *big.Int
		expectErr bool
	}{
		{
			name:      "valid decimal string",
			inputStr:  "12345",
			base:      10,
			expected:  big.NewInt(12345),
			expectErr: false,
		},
		{
			name:      "valid hexadecimal string",
			inputStr:  "1A3F",
			base:      16,
			expected:  big.NewInt(0x1A3F),
			expectErr: false,
		},
		{
			name:      "valid binary string",
			inputStr:  "10101",
			base:      2,
			expected:  big.NewInt(21),
			expectErr: false,
		},
		{
			name:      "invalid string",
			inputStr:  "notanumber",
			base:      10,
			expected:  nil,
			expectErr: true,
		},
		{
			name:      "empty string",
			inputStr:  "",
			base:      10,
			expected:  nil,
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := StringToBigInt(test.inputStr, test.base)

			switch {
			case test.expectErr && err == nil:
				t.Errorf("expected an error but got nil")
			case !test.expectErr && err != nil:
				t.Errorf("did not expect an error but got: %v", err)
			case !test.expectErr && result.Cmp(test.expected) != 0:
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func BenchmarkStringToBigInt(b *testing.B) {
	benchmarks := []struct {
		name     string
		inputStr string
		base     int
	}{
		{
			name:     "Large decimal number",
			inputStr: "123456789012345678901234567890",
			base:     10,
		},
		{
			name:     "Large hexadecimal number",
			inputStr: "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
			base:     16,
		},
		{
			name:     "Large binary number",
			inputStr: "1111111111111111111111111111111111111111111111111111111111111111",
			base:     2,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = StringToBigInt(bm.inputStr, bm.base)
			}
		})
	}
}
