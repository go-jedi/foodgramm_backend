package mathutils

import (
	"errors"
	"fmt"
	"math/big"
)

var ErrConversionFailed = errors.New("failed to convert string to a number")

// StringToBigInt convert string to *big.Int.
func StringToBigInt(str string, base int) (*big.Int, error) {
	bigInt := new(big.Int)
	_, ok := bigInt.SetString(str, base)
	if !ok {
		return nil, fmt.Errorf("%w: input '%s', base %d", ErrConversionFailed, str, base)
	}
	return bigInt, nil
}
