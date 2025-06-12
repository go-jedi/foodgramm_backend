package dependencies

import (
	"github.com/go-jedi/foodgramm_backend/internal/adapters/http/handlers/auth"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	authService "github.com/go-jedi/foodgramm_backend/internal/service/auth"
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
			d.cookie,
			d.middleware,
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.authHandler
}
