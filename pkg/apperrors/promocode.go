package apperrors

import "errors"

var (
	ErrPromoCodeAlreadyExists     = errors.New("promo code already exists")
	ErrPromoCodeIsNotValidForUser = errors.New("promo code is not valid for user")
)
