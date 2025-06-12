package recipeevent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

func (r *repo) DeleteByID(ctx context.Context, recipeID int64) (int64, error) {
	r.logger.Debug("[delete recipe event by id] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `
		DELETE FROM event_recipes
		WHERE id = $1;
	`

	commandTag, err := r.db.Pool.Exec(
		ctxTimeout, q,
		recipeID,
	)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while delete recipe event by id", "err", err)
			return 0, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to delete recipe event by id", "err", err)
		return 0, fmt.Errorf("could not delete recipe event by id: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return 0, apperrors.ErrNoRowsWereAffected
	}

	return recipeID, nil
}
