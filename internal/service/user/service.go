package user

import (
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
	"github.com/go-jedi/foodgramm_backend/pkg/redis"
)

type serv struct {
	userRepository repository.UserRepository
	logger         *logger.Logger
	cache          *redis.Redis
}

func NewService(
	userRepository repository.UserRepository,
	logger *logger.Logger,
	cache *redis.Redis,
) service.UserService {
	return &serv{
		userRepository: userRepository,
		logger:         logger,
		cache:          cache,
	}
}
