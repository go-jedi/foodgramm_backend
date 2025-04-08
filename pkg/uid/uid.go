package uid

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

const (
	defaultCharSet = "0123456789ABCDEFHKLMNPQRSTUVWXYZabcdefghkmnpqrstuvwxyz"
	defaultLength  = 15
	minCharSetLen  = 10 // minimum length of symbol set for cryptographic strength
)

var (
	ErrInvalidOption                         = errors.New("invalid option")
	ErrCharacterSetTooShort                  = errors.New("character set must contain at least 10 unique characters for cryptographic security")
	ErrCharacterSetContainsNonPrintableASCII = errors.New("character set contains non-printable ASCII characters")
)

// IUID defines the interface for the uid.
//
//go:generate mockery --name=IUID --output=mocks --case=underscore
type IUID interface {
	Generate() (string, error)
	Validate(id string) bool
}

type UID struct {
	charSet []rune
	length  int
}

// NewUID creates a new uid instance with the given parameters or default values.
func NewUID(opt Option) (*UID, error) {
	if err := ValidateOption(opt); err != nil {
		return nil, err
	}

	u := &UID{
		charSet: []rune(opt.Chars),
		length:  opt.Count,
	}

	if err := u.init(); err != nil {
		return nil, err
	}

	return u, nil
}

// init initializes the uid, setting default values if none were provided.
func (u *UID) init() error {
	if len(u.charSet) == 0 {
		u.charSet = []rune(defaultCharSet)
	}

	if u.length <= 0 {
		u.length = defaultLength
	}

	return nil
}

// Generate generate unique uid.
func (u *UID) Generate() (string, error) {
	uid := make([]rune, u.length)
	maxIndex := big.NewInt(int64(len(u.charSet)))

	for i := range uid {
		num, err := rand.Int(rand.Reader, maxIndex)
		if err != nil {
			return "", err
		}

		uid[i] = u.charSet[num.Int64()]
	}

	return string(uid), nil
}

// Validate check correct uid.
func (u *UID) Validate(id string) bool {
	if len(id) != u.length {
		return false
	}

	for _, char := range id {
		found := false

		for _, validChar := range u.charSet {
			if char == validChar {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

// ValidateOption checks the validity of the passed options.
func ValidateOption(opt Option) error {
	if opt.Count < 0 {
		return ErrInvalidOption
	}

	if opt.Chars == "" {
		return nil
	}

	if len(opt.Chars) < minCharSetLen {
		return fmt.Errorf("character set too short: %w", ErrCharacterSetTooShort)
	}

	for _, char := range opt.Chars {
		if !isPrintableASCII(char) {
			return fmt.Errorf("character set contains non-printable ascii characters: %w", ErrCharacterSetContainsNonPrintableASCII)
		}
	}

	return nil
}

// isPrintableASCII checks if a character is printable ASCII.
func isPrintableASCII(r rune) bool {
	return r >= 32 && r <= 126
}
