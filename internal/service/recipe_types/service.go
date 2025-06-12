package recipetypes

import (
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
	"github.com/go-jedi/foodgramm_backend/pkg/redis"
)

type serv struct {
	recipeTypesRepository repository.RecipeTypesRepository
	logger                *logger.Logger
	cache                 *redis.Redis
}

func NewService(
	recipeTypesRepository repository.RecipeTypesRepository,
	logger *logger.Logger,
	cache *redis.Redis,
) service.RecipeTypesService {
	return &serv{
		recipeTypesRepository: recipeTypesRepository,
		logger:                logger,
		cache:                 cache,
	}
}
