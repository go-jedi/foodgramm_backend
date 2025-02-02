package recipe

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
)

func (s *serv) GetFreeRecipesByTelegramID(ctx context.Context, telegramID string) (recipe.UserFreeRecipes, error) {
	return s.recipeRepository.GetFreeRecipesByTelegramID(ctx, telegramID)
}
