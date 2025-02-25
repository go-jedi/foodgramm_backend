package recipetypes

import (
	"context"

	recipetypes "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_types"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) Create(ctx context.Context, dto recipetypes.CreateDTO) (recipetypes.RecipeTypes, error) {
	s.logger.Debug("[create a new recipe type] execute service")

	ie, err := s.recipeTypesRepository.Exists(ctx, dto.Title)
	if err != nil {
		return recipetypes.RecipeTypes{}, err
	}

	if ie {
		return recipetypes.RecipeTypes{}, apperrors.ErrRecipeTypeAlreadyExists
	}

	return s.recipeTypesRepository.Create(ctx, dto)
}
