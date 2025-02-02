package middleware

import (
	"log"

	"github.com/go-jedi/foodgrammm-backend/pkg/jwt"
)

type Middleware struct {
	Auth *AuthMiddleware
}

func NewMiddleware(jwt *jwt.JWT) *Middleware {
	if jwt == nil {
		log.Fatal("JWT instance cannot be nil")
	}

	return &Middleware{
		Auth: NewAuthMiddleware(jwt),
	}
}
