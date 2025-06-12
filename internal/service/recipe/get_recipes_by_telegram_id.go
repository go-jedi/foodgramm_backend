package recipe

import (
	"context"

	"github.com/go-jedi/foodgramm_backend/internal/domain/recipe"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

func (s *serv) GetRecipesByTelegramID(ctx context.Context, telegramID string) ([]recipe.Recipes, error) {
	s.logger.Debug("[get recipes by telegram id] execute service")

	ie, err := s.userRepository.ExistsByTelegramID(ctx, telegramID)
	if err != nil {
		return nil, err
	}

	if !ie {
		return nil, apperrors.ErrUserDoesNotExist
	}

	return s.recipeRepository.GetRecipesByTelegramID(ctx, telegramID)
}
