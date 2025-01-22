package user

import (
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
)

type serv struct {
	userRepository repository.UserRepository
	logger         *logger.Logger
}

func NewService(
	userRepository repository.UserRepository,
	logger *logger.Logger,
) service.UserService {
	return &serv{
		userRepository: userRepository,
		logger:         logger,
	}
}
