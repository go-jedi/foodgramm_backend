package recipe

import (
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
)

type serv struct {
	recipeRepository repository.RecipeRepository
	logger           *logger.Logger
	cache            *redis.Redis
}

func NewService(
	recipeRepository repository.RecipeRepository,
	logger *logger.Logger,
	cache *redis.Redis) service.RecipeService {
	return &serv{
		recipeRepository: recipeRepository,
		logger:           logger,
		cache:            cache,
	}
}
