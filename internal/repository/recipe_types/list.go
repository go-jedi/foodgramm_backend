package recipetypes

import (
	"context"

	recipetypes "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_types"
)

func (r *repo) List(_ context.Context) ([]recipetypes.RecipeTypes, error) {
	return nil, nil
}
