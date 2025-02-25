package recipetypes

import (
	"context"

	recipetypes "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_types"
)

func (r *repo) Update(_ context.Context, _ recipetypes.UpdateDTO) (recipetypes.RecipeTypes, error) {
	return recipetypes.RecipeTypes{}, nil
}
