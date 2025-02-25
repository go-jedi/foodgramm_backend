package recipetypes

import (
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
)

type repo struct {
	logger *logger.Logger
	db     *postgres.Postgres
}

func NewRepository(
	l *logger.Logger,
	p *postgres.Postgres,
) repository.RecipeTypesRepository {
	return &repo{
		logger: l,
		db:     p,
	}
}
