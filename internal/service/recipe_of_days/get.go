package recipeofdays

import (
	"context"

	recipeofdays "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_of_days"
)

func (s *serv) Get(ctx context.Context) (recipeofdays.Recipe, error) {
	s.logger.Debug("[get recipe of the day] execute service")

	return s.recipeOfDaysRepository.Get(ctx)
}
