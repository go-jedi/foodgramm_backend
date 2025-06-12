package promocode

import (
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
	"github.com/go-jedi/foodgramm_backend/pkg/postgres"
)

type repo struct {
	logger *logger.Logger
	db     *postgres.Postgres
}

func NewRepository(
	l *logger.Logger,
	p *postgres.Postgres,
) repository.PromoCodeRepository {
	return &repo{
		logger: l,
		db:     p,
	}
}
