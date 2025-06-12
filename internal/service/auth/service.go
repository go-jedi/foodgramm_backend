package auth

import (
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/pkg/jwt"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
	"github.com/go-jedi/foodgramm_backend/pkg/redis"
)

type serv struct {
	userRepository repository.UserRepository
	logger         *logger.Logger
	jwt            *jwt.JWT
	cache          *redis.Redis
}

func NewService(
	userRepository repository.UserRepository,
	logger *logger.Logger,
	jwt *jwt.JWT,
	cache *redis.Redis,
) service.AuthService {
	return &serv{
		userRepository: userRepository,
		logger:         logger,
		jwt:            jwt,
		cache:          cache,
	}
}
