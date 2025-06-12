package recipetypes

import (
	"context"

	recipetypes "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_types"
)

func (s *serv) All(ctx context.Context) ([]recipetypes.RecipeTypes, error) {
	s.logger.Debug("[get all recipe types] execute service")

	return s.recipeTypesRepository.All(ctx)
}
