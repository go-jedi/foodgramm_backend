package recipetypes

import (
	"context"

	recipetypes "github.com/go-jedi/foodgramm_backend/internal/domain/recipe_types"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

func (s *serv) Update(ctx context.Context, dto recipetypes.UpdateDTO) (recipetypes.RecipeTypes, error) {
	s.logger.Debug("[update recipe type] execute service")

	ie, err := s.recipeTypesRepository.ExistsExceptCurrent(ctx, dto.ID, dto.Title)
	if err != nil {
		return recipetypes.RecipeTypes{}, err
	}

	if ie {
		return recipetypes.RecipeTypes{}, apperrors.ErrRecipeTypeAlreadyExists
	}

	return s.recipeTypesRepository.Update(ctx, dto)
}
