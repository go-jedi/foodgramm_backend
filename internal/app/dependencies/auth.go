package dependencies

import (
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/auth"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	authService "github.com/go-jedi/foodgrammm-backend/internal/service/auth"
)

func (d *Dependencies) AuthService() service.AuthService {
	if d.authService == nil {
		d.authService = authService.NewService(
			d.UserRepository(),
			d.logger,
			d.jwt,
			d.cache,
		)
	}

	return d.authService
}

func (d *Dependencies) AuthHandler() *auth.Handler {
	if d.authHandler == nil {
		d.authHandler = auth.NewHandler(
			d.AuthService(),
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.authHandler
}
