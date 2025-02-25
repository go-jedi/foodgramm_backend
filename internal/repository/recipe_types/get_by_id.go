package recipetypes

import (
	"context"

	recipetypes "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_types"
)

func (r *repo) GetByID(_ context.Context, _ int64) (recipetypes.RecipeTypes, error) {
	return recipetypes.RecipeTypes{}, nil
}
