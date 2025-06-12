package product

import (
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
	"github.com/go-jedi/foodgramm_backend/pkg/redis"
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
