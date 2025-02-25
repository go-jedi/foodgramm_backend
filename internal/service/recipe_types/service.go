package recipetypes

import (
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
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
