package apperrors

import "errors"

var (
	ErrUserInBlackListAlreadyExists = errors.New("user in blacklist already exists")
	ErrUserInBlackListDoesNotExist  = errors.New("user in blacklist does not exist")
)
