package recipe

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) GetFreeRecipesByTelegramID(ctx context.Context, telegramID string) (recipe.UserFreeRecipes, error) {
	s.logger.Debug("[GetFreeRecipesByTelegramID] execute service")

	ie, err := s.userRepository.ExistsByTelegramID(ctx, telegramID)
	if err != nil {
		return recipe.UserFreeRecipes{}, err
	}

	if !ie {
		return recipe.UserFreeRecipes{}, apperrors.ErrUserDoesNotExist
	}

	return s.recipeRepository.GetFreeRecipesByTelegramID(ctx, telegramID)
}
