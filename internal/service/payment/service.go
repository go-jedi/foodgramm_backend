package payment

import (
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
)

type serv struct {
	logger *logger.Logger
}

func NewService(
	logger *logger.Logger,
) service.PaymentService {
	return &serv{
		logger: logger,
	}
}
