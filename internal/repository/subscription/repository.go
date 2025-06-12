package subscription

import (
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
	"github.com/go-jedi/foodgramm_backend/pkg/postgres"
)

type repo struct {
	logger logger.ILogger
	db     *postgres.Postgres
}

func NewRepository(
	l logger.ILogger,
	p *postgres.Postgres,
) repository.SubscriptionRepository {
	return &repo{
		logger: l,
		db:     p,
	}
}
