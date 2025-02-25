package recipetypes

import (
	"context"

	recipetypes "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_types"
)

func (s *serv) Update(ctx context.Context, dto recipetypes.UpdateDTO) (recipetypes.RecipeTypes, error) {
	return s.recipeTypesRepository.Update(ctx, dto)
}
