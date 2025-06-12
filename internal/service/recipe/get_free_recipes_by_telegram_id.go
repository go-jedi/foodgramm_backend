package recipe

import (
	"context"

	"github.com/go-jedi/foodgramm_backend/internal/domain/recipe"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

func (s *serv) GetFreeRecipesByTelegramID(ctx context.Context, telegramID string) (recipe.UserFreeRecipes, error) {
	s.logger.Debug("[get free recipes by telegram id] execute service")

	ie, err := s.userRepository.ExistsByTelegramID(ctx, telegramID)
	if err != nil {
		return recipe.UserFreeRecipes{}, err
	}

	if !ie {
		return recipe.UserFreeRecipes{}, apperrors.ErrUserDoesNotExist
	}

	return s.recipeRepository.GetFreeRecipesByTelegramID(ctx, telegramID)
}
