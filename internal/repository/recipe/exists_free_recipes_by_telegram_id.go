package recipe

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (r *repo) ExistsFreeRecipesByTelegramID(ctx context.Context, telegramID string) (bool, error) {
	r.logger.Debug("[check exists free recipes by telegram id] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	ie := false

	q := `
		SELECT EXISTS(
			SELECT 1
			FROM user_free_recipes
			WHERE telegram_id = $1
			AND free_recipes_used > 0
		);
	`

	if err := r.db.Pool.QueryRow(ctxTimeout, q, telegramID).Scan(&ie); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while check exists free recipes by telegram id", "err", err)
			return false, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to check exists free recipes by telegram id", "err", err)
		return false, fmt.Errorf("could not check exists free recipes by telegram id: %w", err)
	}

	return ie, nil
}
