package recipetypes

import (
	"context"

	recipetypes "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_types"
)

func (s *serv) List(ctx context.Context) ([]recipetypes.RecipeTypes, error) {
	return s.recipeTypesRepository.List(ctx)
}
