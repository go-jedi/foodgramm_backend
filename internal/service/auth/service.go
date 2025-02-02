package auth

import (
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/jwt"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
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
