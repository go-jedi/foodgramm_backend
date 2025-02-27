package recipeevent

import (
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
)

type serv struct {
	recipeEventRepository repository.RecipeEventRepository
	recipeTypesRepository repository.RecipeTypesRepository
	logger                *logger.Logger
	cache                 *redis.Redis
}

func NewService(
	recipeEventRepository repository.RecipeEventRepository,
	recipeTypesRepository repository.RecipeTypesRepository,
	logger *logger.Logger,
	cache *redis.Redis,
) service.RecipeEventService {
	return &serv{
		recipeEventRepository: recipeEventRepository,
		recipeTypesRepository: recipeTypesRepository,
		logger:                logger,
		cache:                 cache,
	}
}
