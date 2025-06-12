package recipetypes

import (
	"context"

	recipetypes "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_types"
)

func (s *serv) GetByID(ctx context.Context, recipeTypeID int64) (recipetypes.RecipeTypes, error) {
	s.logger.Debug("[get recipe type by id] execute service")

	return s.recipeTypesRepository.GetByID(ctx, recipeTypeID)
}
