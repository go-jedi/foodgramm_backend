package apperrors

import "errors"

var (
	ErrCacheKeyNotExists = errors.New("cache key does not exists")
)
