package user

import (
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
)

type repo struct {
	logger logger.ILogger
	db     *postgres.Postgres
}

func NewRepository(
	l logger.ILogger,
	p *postgres.Postgres,
) repository.UserRepository {
	return &repo{
		logger: l,
		db:     p,
	}
}
