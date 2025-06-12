package recipeofdays

import (
	"github.com/go-jedi/foodgramm_backend/internal/client"
	"github.com/go-jedi/foodgramm_backend/internal/parser"
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/internal/templates"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
	"github.com/go-jedi/foodgramm_backend/pkg/redis"
)

type serv struct {
	recipeOfDaysRepository repository.RecipeOfDaysRepository
	client                 *client.Client
	templates              *templates.Templates
	parser                 *parser.Parser
	logger                 *logger.Logger
	cache                  *redis.Redis
}

func NewService(
	recipeOfDaysRepository repository.RecipeOfDaysRepository,
	client *client.Client,
	templates *templates.Templates,
	parser *parser.Parser,
	logger *logger.Logger,
	cache *redis.Redis,
) service.RecipeOfDaysService {
	return &serv{
		recipeOfDaysRepository: recipeOfDaysRepository,
		client:                 client,
		templates:              templates,
		parser:                 parser,
		logger:                 logger,
		cache:                  cache,
	}
}
