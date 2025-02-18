package recipe

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
)

func (r *repo) GetFreeRecipesByTelegramID(ctx context.Context, telegramID string) (recipe.UserFreeRecipes, error) {
	r.logger.Debug("[get free recipes by telegram id] execute repository")

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
			r.logger.Error("request timed out while get free recipes by telegram id", "err", err)
			return recipe.UserFreeRecipes{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get free recipes by telegram id", "err", err)
		return recipe.UserFreeRecipes{}, fmt.Errorf("could not get free recipes by telegram id: %w", err)
	}

	return ufr, nil
}
