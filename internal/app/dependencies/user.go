package dependencies

import (
	"github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/user"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	userRepository "github.com/go-jedi/foodgrammm-backend/internal/repository/user"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	userService "github.com/go-jedi/foodgrammm-backend/internal/service/user"
)

func (d *Dependencies) UserRepository() repository.UserRepository {
	if d.userRepository == nil {
		d.userRepository = userRepository.NewRepository(
			d.logger,
			d.db,
		)
	}

	return d.userRepository
}

func (d *Dependencies) UserService() service.UserService {
	if d.userService == nil {
		d.userService = userService.NewService(
			d.UserRepository(),
			d.logger,
			d.cache,
		)
	}

	return d.userService
}

func (d *Dependencies) UserHandler() *user.Handler {
	if d.userHandler == nil {
		d.userHandler = user.NewHandler(
			d.UserService(),
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.userHandler
}
