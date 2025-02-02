package recipe

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
)

func (r *repo) AddFreeRecipesCountByTelegramID(ctx context.Context, telegramID string) (recipe.UserFreeRecipes, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var uufr recipe.UserFreeRecipes

	q := `
		UPDATE user_free_recipes SET
			free_recipes_used = free_recipes_used + 1,
			updated_at = NOW()
		WHERE telegram_id = $1
		RETURNING *;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q, telegramID,
	).Scan(
		&uufr.ID, &uufr.TelegramID,
		&uufr.FreeRecipesAllowed, &uufr.FreeRecipesUsed,
		&uufr.CreatedAt, &uufr.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while add free recipes by telegram id", "err", err)
			return recipe.UserFreeRecipes{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to add free recipes by telegram id", "err", err)
		return recipe.UserFreeRecipes{}, fmt.Errorf("could not add free recipes by telegram id: %w", err)
	}

	return uufr, nil
}
