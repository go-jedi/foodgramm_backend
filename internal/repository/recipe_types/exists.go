package recipetypes

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (r *repo) Exists(ctx context.Context, title string) (bool, error) {
	r.logger.Debug("[check a recipe type exists] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	ie := false

	q := `
		SELECT EXISTS(
			SELECT 1
			FROM recipe_types
			WHERE title = $1
		);
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		title,
	).Scan(&ie); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while check exists recipe type", "err", err)
			return false, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to check exists recipe type", "err", err)
		return false, fmt.Errorf("could not check exists recipe type: %w", err)
	}

	return ie, nil
}
