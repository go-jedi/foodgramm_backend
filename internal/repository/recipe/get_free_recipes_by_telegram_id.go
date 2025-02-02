package recipe

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
)

func (r *repo) GetFreeRecipesByTelegramID(ctx context.Context, telegramID string) (recipe.UserFreeRecipes, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var ufr recipe.UserFreeRecipes

	q := `
		SELECT *
		FROM user_free_recipes
		WHERE telegram_id = $1;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q, telegramID,
	).Scan(
		&ufr.ID, &ufr.TelegramID,
		&ufr.FreeRecipesAllowed, &ufr.FreeRecipesUsed,
		&ufr.CreatedAt, &ufr.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get recipes by telegram id", "err", err)
			return recipe.UserFreeRecipes{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get recipes by telegram id", "err", err)
		return recipe.UserFreeRecipes{}, fmt.Errorf("could not get recipes by telegram id: %w", err)
	}

	return ufr, nil
}
