package clientassets

import (
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	fileserver "github.com/go-jedi/foodgrammm-backend/pkg/file_server"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
)

type serv struct {
	clientAssetsRepository repository.ClientAssetsRepository
	fileServer             *fileserver.FileServer
	logger                 *logger.Logger
	cache                  *redis.Redis
}

func NewService(
	clientAssetsRepository repository.ClientAssetsRepository,
	fileServer *fileserver.FileServer,
	logger *logger.Logger,
	cache *redis.Redis,
) service.ClientAssetsService {
	return &serv{
		clientAssetsRepository: clientAssetsRepository,
		fileServer:             fileServer,
		logger:                 logger,
		cache:                  cache,
	}
}
