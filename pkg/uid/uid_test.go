package uid

import (
	"fmt"
	"testing"
)

// TestNewUID_DefaultValues tests that NewUID sets default values when no options are provided.
func TestNewUID_DefaultValues(t *testing.T) {
	u, err := NewUID(Option{})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if string(u.charSet) != defaultCharSet {
		t.Errorf("expected charSet to be %s, got %s", defaultCharSet, string(u.charSet))
	}

	if u.length != defaultLength {
		t.Errorf("expected length to be %d, got %d", defaultLength, u.length)
	}
}

// TestNewUID_CustomValues tests that NewUID sets custom values when provided.
func TestNewUID_CustomValues(t *testing.T) {
	customCharSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	customLength := 10

	u, err := NewUID(Option{Chars: customCharSet, Count: customLength})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if string(u.charSet) != customCharSet {
		t.Errorf("expected charSet to be %s, got %s", customCharSet, string(u.charSet))
	}

	if u.length != customLength {
		t.Errorf("expected length to be %d, got %d", customLength, u.length)
	}
}

// TestNewUID_InvalidOption tests that NewUID returns an error for invalid options.
func TestNewUID_InvalidOption(t *testing.T) {
	tests := []struct {
		option  Option
		wantErr error
	}{
		{
			option:  Option{Count: -1},
			wantErr: ErrInvalidOption,
		},
		{
			option:  Option{Chars: "short"},
			wantErr: fmt.Errorf("character set too short: %w", ErrCharacterSetTooShort),
		},
	}

	for _, test := range tests {
		_, err := NewUID(test.option)
		if err == nil {
			t.Errorf("expected error %v, got nil", test.wantErr)
		} else if err.Error() != test.wantErr.Error() {
			t.Errorf("expected error %v, got %v", test.wantErr, err)
		}
	}
}

// TestGenerate tests that Generate produces a valid UID.
func TestGenerate(t *testing.T) {
	u, err := NewUID(Option{})
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	uid, err := u.Generate()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if !u.Validate(uid) {
		t.Errorf("eenerated UID %s is not valid", uid)
	}
}

// TestValidate tests that Validate correctly checks UIDs.
func TestValidate(t *testing.T) {
	tests := []struct {
		id      string
		valid   bool
		options Option
	}{
		{
			id:      "1234567890ABCDE",
			valid:   true,
			options: Option{},
		},
		{
			id:      "1234567890ABCDEF",
			valid:   false,
			options: Option{},
		},
		{
			id:    "1234567890ABCDEFG",
			valid: false,
			options: Option{
				Count: 15,
			},
		},
		{
			id:      "1234567890",
			valid:   false,
			options: Option{},
		},
		{
			id:    "1234567890abcdefg",
			valid: false,
			options: Option{
				Chars: "0123456789abcdefg",
			},
		},
	}

	for _, test := range tests {
		u, err := NewUID(test.options)
		if err != nil {
			t.Log(test.options)
			t.Fatalf("expected no error, got: %v", err)
		}

		if valid := u.Validate(test.id); valid != test.valid {
			t.Errorf("expected UID %s to be valid=%v, got valid=%v", test.id, test.valid, valid)
		}
	}
}

// BenchmarkGenerateDefault tests the performance of Generate method with default settings.
func BenchmarkGenerateDefault(b *testing.B) {
	uid, err := NewUID(Option{})
	if err != nil {
		b.Fatalf("failed to create UID instance: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := uid.Generate()
		if err != nil {
			b.Fatalf("failed to generate UID: %v", err)
		}
	}
}

// BenchmarkGenerateCustom tests the performance of Generate method with custom settings.
func BenchmarkGenerateCustom(b *testing.B) {
	customCharset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	customLength := 20

	uid, err := NewUID(Option{Chars: customCharset, Count: customLength})
	if err != nil {
		b.Fatalf("failed to create UID instance: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := uid.Generate()
		if err != nil {
			b.Fatalf("failed to generate UID: %v", err)
		}
	}
}

// BenchmarkValidateDefault tests the performance of Validate method with default settings.
func BenchmarkValidateDefault(b *testing.B) {
	uid, err := NewUID(Option{})
	if err != nil {
		b.Fatalf("failed to create UID instance: %v", err)
	}

	testUID, err := uid.Generate()
	if err != nil {
		b.Fatalf("failed to generate UID for validation test: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uid.Validate(testUID)
	}
}

// BenchmarkValidateCustom tests the performance of Validate method with custom settings.
func BenchmarkValidateCustom(b *testing.B) {
	customCharset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	customLength := 20

	uid, err := NewUID(Option{Chars: customCharset, Count: customLength})
	if err != nil {
		b.Fatalf("failed to create UID instance: %v", err)
	}

	testUID, err := uid.Generate()
	if err != nil {
		b.Fatalf("failed to generate UID for validation test: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uid.Validate(testUID)
	}
}
