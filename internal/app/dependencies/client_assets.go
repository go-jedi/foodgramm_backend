package dependencies

import (
	clientassets "github.com/go-jedi/foodgrammm-backend/internal/adapters/http/handlers/client_assets"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	clientAssetsRepository "github.com/go-jedi/foodgrammm-backend/internal/repository/client_assets"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	clientAssetsService "github.com/go-jedi/foodgrammm-backend/internal/service/client_assets"
)

func (d *Dependencies) ClientAssetsRepository() repository.ClientAssetsRepository {
	if d.clientAssetsRepository == nil {
		d.clientAssetsRepository = clientAssetsRepository.NewRepository(
			d.logger,
			d.db,
		)
	}

	return d.clientAssetsRepository
}

func (d *Dependencies) ClientAssetsService() service.ClientAssetsService {
	if d.clientAssetsService == nil {
		d.clientAssetsService = clientAssetsService.NewService(
			d.ClientAssetsRepository(),
			d.fileServer,
			d.logger,
			d.cache,
		)
	}

	return d.clientAssetsService
}

func (d *Dependencies) ClientAssetsHandler() *clientassets.Handler {
	if d.clientAssetsHandler == nil {
		d.clientAssetsHandler = clientassets.NewHandler(
			d.ClientAssetsService(),
			d.cookie,
			d.middleware,
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.clientAssetsHandler
}
