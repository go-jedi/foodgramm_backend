package bcrypt

import (
	"testing"

	"github.com/go-jedi/foodgramm_backend/config"
)

// TestNewBcrypt tests the NewBcrypt function.
func TestNewBcrypt(t *testing.T) {
	bcryptInstance, err := New()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if bcryptInstance.cost != defaultCost {
		t.Errorf("expected cost to be %d, got %d", defaultCost, bcryptInstance.cost)
	}
}

// TestNewBcryptWithCostValid tests the NewBcryptWithCost function with a valid cost.
func TestNewBcryptWithCostValid(t *testing.T) {
	cfg := config.BcryptConfig{Cost: 10}

	bcryptInstance, err := NewBcryptWithCost(cfg)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if bcryptInstance.cost != cfg.Cost {
		t.Errorf("expected cost to be %d, got %d", cfg.Cost, bcryptInstance.cost)
	}
}

// TestNewBcryptWithCostInvalid tests the NewBcryptWithCost function with an invalid cost.
func TestNewBcryptWithCostInvalid(t *testing.T) {
	cfg := config.BcryptConfig{Cost: 100} // Invalid cost

	_, err := NewBcryptWithCost(cfg)
	if err == nil {
		t.Errorf("expected error for invalid cost, got none")
	}

	if err.Error() != "invalid cost value: 100. Cost must be between 4 and 31" {
		t.Errorf("expected specific error message, got: %v", err)
	}
}

// TestGenerateHash tests the GenerateHash function.
func TestGenerateHash(t *testing.T) {
	bcryptInstance, err := New()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	password := "test_password"
	hashedPassword, err := bcryptInstance.GenerateHash(password)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if !bcryptInstance.IsBcryptHash(hashedPassword) {
		t.Errorf("generated hash is not a valid bcrypt hash: %s", hashedPassword)
	}
}

// TestCompareHashAndPasswordCorrect tests the CompareHashAndPassword function with a correct password.
func TestCompareHashAndPasswordCorrect(t *testing.T) {
	bcryptInstance, err := New()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	password := "test_password"
	hashedPassword, err := bcryptInstance.GenerateHash(password)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if err = bcryptInstance.CompareHashAndPassword(hashedPassword, password); err != nil {
		t.Errorf("expected no error for correct password, got: %v", err)
	}
}

// TestCompareHashAndPasswordIncorrect tests the CompareHashAndPassword function with an incorrect password.
func TestCompareHashAndPasswordIncorrect(t *testing.T) {
	bcryptInstance, err := New()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	password := "test_password"
	hashedPassword, err := bcryptInstance.GenerateHash(password)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	incorrectPassword := "wrong_password"
	if err = bcryptInstance.CompareHashAndPassword(hashedPassword, incorrectPassword); err == nil {
		t.Errorf("expected error for incorrect password, got none")
	}
}

// TestIsBcryptHashValid tests the IsBcryptHash function with a valid bcrypt hash.
func TestIsBcryptHashValid(t *testing.T) {
	bcryptInstance, err := New()
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	password := "test_password"
	hashedPassword, err := bcryptInstance.GenerateHash(password)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if !bcryptInstance.IsBcryptHash(hashedPassword) {
		t.Errorf("expected true for valid bcrypt hash, got false")
	}
}

// TestIsBcryptHashInvalid tests the IsBcryptHash function with an invalid bcrypt hash.
func TestIsBcryptHashInvalid(t *testing.T) {
	bcryptInstance := &Bcrypt{}

	invalidHash := "invalid_hash"
	if bcryptInstance.IsBcryptHash(invalidHash) {
		t.Errorf("expected false for invalid bcrypt hash, got true")
	}
}

// BenchmarkGenerateHash benchmarks the GenerateHash function.
func BenchmarkGenerateHash(b *testing.B) {
	bcryptInstance, err := New()
	if err != nil {
		b.Fatalf("expected no error, got: %v", err)
	}

	password := "test_password"
	for i := 0; i < b.N; i++ {
		_, err := bcryptInstance.GenerateHash(password)
		if err != nil {
			b.Fatalf("expected no error, got: %v", err)
		}
	}
}

// BenchmarkCompareHashAndPasswordCorrect benchmarks the CompareHashAndPassword function with a correct password.
func BenchmarkCompareHashAndPasswordCorrect(b *testing.B) {
	bcryptInstance, err := New()
	if err != nil {
		b.Fatalf("expected no error, got: %v", err)
	}

	password := "test_password"
	hashedPassword, err := bcryptInstance.GenerateHash(password)
	if err != nil {
		b.Fatalf("expected no error, got: %v", err)
	}

	for i := 0; i < b.N; i++ {
		err := bcryptInstance.CompareHashAndPassword(hashedPassword, password)
		if err != nil {
			b.Fatalf("expected no error, got: %v", err)
		}
	}
}

// BenchmarkCompareHashAndPasswordIncorrect benchmarks the CompareHashAndPassword function with an incorrect password.
func BenchmarkCompareHashAndPasswordIncorrect(b *testing.B) {
	bcryptInstance, err := New()
	if err != nil {
		b.Fatalf("expected no error, got: %v", err)
	}

	password := "test_password"
	hashedPassword, err := bcryptInstance.GenerateHash(password)
	if err != nil {
		b.Fatalf("expected no error, got: %v", err)
	}

	incorrectPassword := "wrong_password"
	for i := 0; i < b.N; i++ {
		_ = bcryptInstance.CompareHashAndPassword(hashedPassword, incorrectPassword)
	}
}
