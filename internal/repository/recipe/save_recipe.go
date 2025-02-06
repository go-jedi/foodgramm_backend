package recipe

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (r *repo) SaveRecipe(ctx context.Context, data recipe.InCache) (bool, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		INSERT INTO recipes(
			telegram_id,
		    title,
		    content
		) VALUES($1, $2, $3);
	`

	ct, err := r.db.Pool.Exec(
		ctxTimeout, q,
		data.TelegramID, data.Title, data.Content,
	)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while save recipe", "err", err)
			return false, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to save recipe", "err", err)
		return false, fmt.Errorf("could not save recipe: %w", err)
	}

	if ct.RowsAffected() != 1 {
		return false, apperrors.ErrNoRowsWereAffected
	}

	return true, nil
}
