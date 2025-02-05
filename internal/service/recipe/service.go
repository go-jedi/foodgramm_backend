package recipe

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
	recipeRepository       repository.RecipeRepository
	userRepository         repository.UserRepository
	productRepository      repository.ProductRepository
	subscriptionRepository repository.SubscriptionRepository
	client                 *client.Client
	templates              *templates.Templates
	parser                 *parser.Parser
	logger                 *logger.Logger
	cache                  *redis.Redis
}

func NewService(
	recipeRepository repository.RecipeRepository,
	userRepository repository.UserRepository,
	productRepository repository.ProductRepository,
	subscriptionRepository repository.SubscriptionRepository,
	client *client.Client,
	templates *templates.Templates,
	parser *parser.Parser,
	logger *logger.Logger,
	cache *redis.Redis,
) service.RecipeService {
	return &serv{
		recipeRepository:       recipeRepository,
		userRepository:         userRepository,
		productRepository:      productRepository,
		subscriptionRepository: subscriptionRepository,
		client:                 client,
		templates:              templates,
		parser:                 parser,
		logger:                 logger,
		cache:                  cache,
	}
}
