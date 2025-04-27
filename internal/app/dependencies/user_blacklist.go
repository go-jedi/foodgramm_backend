package dependencies

import (
	userBlackListHandler "github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/user_blacklist"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	userBlackListRepository "github.com/go-jedi/foodgrammm-backend/internal/repository/user_blacklist"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	userBlackListService "github.com/go-jedi/foodgrammm-backend/internal/service/user_blacklist"
)

func (d *Dependencies) UserBlackListRepository() repository.UserBlackListRepository {
	if d.userBlackListRepository == nil {
		d.userBlackListRepository = userBlackListRepository.NewRepository(
			d.logger,
			d.db,
		)
	}

	return d.userBlackListRepository
}

func (d *Dependencies) UserBlackListService() service.UserBlackListService {
	if d.userBlackListService == nil {
		d.userBlackListService = userBlackListService.NewService(
			d.UserBlackListRepository(),
			d.logger,
			d.cache,
		)
	}

	return d.userBlackListService
}

func (d *Dependencies) UserBlackListHandler() *userBlackListHandler.Handler {
	if d.userBlackListHandler == nil {
		d.userBlackListHandler = userBlackListHandler.NewHandler(
			d.UserBlackListService(),
			d.cookie,
			d.middleware,
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.userBlackListHandler
}
