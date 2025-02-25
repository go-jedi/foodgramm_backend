package recipetypes

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (r *repo) ExistsExceptCurrent(ctx context.Context, recipeTypeID int64, title string) (bool, error) {
	r.logger.Debug("[check recipe type exists except current] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	ie := false

	q := `
		SELECT EXISTS(
			SELECT 1
			FROM recipe_types
			WHERE id != $1
		    AND title = $2
		);
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		recipeTypeID, title,
	).Scan(&ie); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while check exists recipe type except current", "err", err)
			return false, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to check exists recipe type except current", "err", err)
		return false, fmt.Errorf("could not check exists recipe type except current: %w", err)
	}

	return ie, nil
}
