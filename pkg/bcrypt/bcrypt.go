package bcrypt

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/go-jedi/foodgrammm-backend/config"
	"golang.org/x/crypto/bcrypt"
)

const (
	defaultCost = bcrypt.DefaultCost
	minCost     = bcrypt.MinCost
	maxCost     = bcrypt.MaxCost
)

var (
	reFormatBcrypt = regexp.MustCompile(`^\$2[aby]\$[0-9]{2}\$[./A-Za-z0-9]{53}$`)
	errInvalidCost = errors.New("invalid cost value")
)

// IBcrypt defines the interface for the bcrypt.
//
//go:generate mockery --name=IBcrypt --output=mocks --case=underscore
type IBcrypt interface {
	GenerateHash(password string) (string, error)
	CompareHashAndPassword(hashedPassword string, password string) error
	IsBcryptHash(in string) bool
}

type Bcrypt struct {
	cost int
}

// New creates a new Bcrypt instance with the default cost.
func New() (*Bcrypt, error) {
	b := &Bcrypt{
		cost: defaultCost,
	}

	if err := b.init(); err != nil {
		return nil, err
	}

	return b, nil
}

// NewBcryptWithCost creates a new Bcrypt instance with a specified cost.
func NewBcryptWithCost(cfg config.BcryptConfig) (*Bcrypt, error) {
	b := &Bcrypt{
		cost: cfg.Cost,
	}

	if err := b.init(); err != nil {
		return nil, err
	}

	return b, nil
}

// init initializes the Bcrypt instance, ensuring the cost is within valid bounds.
func (b *Bcrypt) init() error {
	if b.cost < minCost || b.cost > maxCost {
		return fmt.Errorf("%w: %d. Cost must be between %d and %d", errInvalidCost, b.cost, minCost, maxCost)
	}

	return nil
}

// GenerateHash generates a bcrypt hash for the given password.
func (b *Bcrypt) GenerateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// CompareHashAndPassword compares the provided hashed password with the plain password.
func (b *Bcrypt) CompareHashAndPassword(hashedPassword string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}

	return nil
}

// IsBcryptHash checks if the input string is a valid bcrypt hash.
func (b *Bcrypt) IsBcryptHash(in string) bool {
	return reFormatBcrypt.MatchString(in)
}
