package apperrors

import "errors"

var (
	ErrRecipeTypeAlreadyExists = errors.New("recipe type already exists")
	ErrRecipeTypeDoesNotExists = errors.New("recipe type does not exist")
)
