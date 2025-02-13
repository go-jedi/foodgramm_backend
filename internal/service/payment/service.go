package payment

import (
	"github.com/go-jedi/foodgrammm-backend/internal/client"
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
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
