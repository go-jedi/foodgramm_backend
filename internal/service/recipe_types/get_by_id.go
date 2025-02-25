package recipetypes

import (
	"context"

	recipetypes "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_types"
)

func (s *serv) GetByID(ctx context.Context, recipeTypeID int64) (recipetypes.RecipeTypes, error) {
	return s.recipeTypesRepository.GetByID(ctx, recipeTypeID)
}
