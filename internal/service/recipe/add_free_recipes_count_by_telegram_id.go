package recipe

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) AddFreeRecipesCountByTelegramID(ctx context.Context, telegramID string) (recipe.UserFreeRecipes, error) {
	ie, err := s.userRepository.ExistsByTelegramID(ctx, telegramID)
	if err != nil {
		return recipe.UserFreeRecipes{}, err
	}

	if !ie {
		return recipe.UserFreeRecipes{}, apperrors.ErrUserDoesNotExist
	}

	return s.recipeRepository.AddFreeRecipesCountByTelegramID(ctx, telegramID)
}
