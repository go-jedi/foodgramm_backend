package recipe

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) SaveRecipeByTelegramID(ctx context.Context, telegramID string) (bool, error) {
	// check user exists by telegram id.
	ie, err := s.userRepository.ExistsByTelegramID(ctx, telegramID)
	if err != nil {
		return false, err
	}

	if !ie {
		return false, apperrors.ErrUserDoesNotExist
	}

	// get recipe from cache.
	ric, err := s.cache.Recipe.Get(ctx, telegramID)
	if err != nil {
		return false, err
	}

	// delete recipe in cache.
	if err := s.cache.Recipe.Del(ctx, telegramID); err != nil {
		return false, err
	}

	// save recipe in database.
	return s.recipeRepository.SaveRecipe(ctx, ric)
}
