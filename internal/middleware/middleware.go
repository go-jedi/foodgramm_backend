package middleware

import (
	"log"

	adminguard "github.com/go-jedi/foodgramm_backend/internal/middleware/admin_guard"
	"github.com/go-jedi/foodgramm_backend/internal/middleware/auth"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/pkg/jwt"
)

type Middleware struct {
	Auth       *auth.Middleware
	AdminGuard *adminguard.Middleware
}

func New(adminService service.AdminService, jwt *jwt.JWT) *Middleware {
	if jwt == nil {
		log.Fatal("JWT instance cannot be nil")
	}

	return &Middleware{
		Auth:       auth.New(jwt),
		AdminGuard: adminguard.New(adminService, jwt),
	}
}
