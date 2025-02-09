package recipeofdays

import (
	"context"

	recipeofdays "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_of_days"
)

func (s *serv) Get(ctx context.Context) (recipeofdays.Recipe, error) {
	return s.recipeOfDaysRepository.Get(ctx)
}
