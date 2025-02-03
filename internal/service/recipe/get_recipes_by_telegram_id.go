package recipe

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) GetRecipesByTelegramID(ctx context.Context, telegramID string) ([]recipe.Recipes, error) {
	ie, err := s.userRepository.ExistsByTelegramID(ctx, telegramID)
	if err != nil {
		return nil, err
	}

	if !ie {
		return nil, apperrors.ErrUserDoesNotExist
	}

	return s.recipeRepository.GetRecipesByTelegramID(ctx, telegramID)
}
