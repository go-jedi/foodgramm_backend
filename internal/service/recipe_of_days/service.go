package recipeofdays

import (
	"github.com/go-jedi/foodgrammm-backend/internal/client"
	"github.com/go-jedi/foodgrammm-backend/internal/parser"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/internal/templates"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
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
