package user

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (r *repo) ExistsByID(ctx context.Context, userID int64) (bool, error) {
	r.logger.Debug("[ExistsByID] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	ie := false

	q := `
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE id = $1
		);
	`

	if err := r.db.Pool.QueryRow(ctxTimeout, q, userID).Scan(&ie); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while check exists user by id", "err", err)
			return false, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to check exists user by id", "err", err)
		return false, fmt.Errorf("could not check exists user by id: %w", err)
	}

	return ie, nil
}
