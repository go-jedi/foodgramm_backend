package product

import (
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
)

type serv struct {
	productRepository repository.ProductRepository
	userRepository    repository.UserRepository
	logger            *logger.Logger
	cache             *redis.Redis
}

func NewService(
	productRepository repository.ProductRepository,
	userRepository repository.UserRepository,
	logger *logger.Logger,
	cache *redis.Redis,
) service.ProductService {
	return &serv{
		productRepository: productRepository,
		userRepository:    userRepository,
		logger:            logger,
		cache:             cache,
	}
}
