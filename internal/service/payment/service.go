package payment

import (
	"github.com/go-jedi/foodgramm_backend/internal/client"
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
)

type serv struct {
	userRepository repository.UserRepository
	client         *client.Client
	logger         *logger.Logger
}

func NewService(
	userRepository repository.UserRepository,
	client *client.Client,
	logger *logger.Logger,
) service.PaymentService {
	return &serv{
		userRepository: userRepository,
		client:         client,
		logger:         logger,
	}
}
