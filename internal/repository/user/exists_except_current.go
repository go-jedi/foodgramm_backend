package user

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (r *repo) ExistsExceptCurrent(ctx context.Context, id int64, telegramID string, username string) (bool, error) {
	r.logger.Debug("[check user exists except current] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	ie := false

	q := `
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE id != $1
			AND (
				telegram_id = $2 OR username = $3
			)
		);
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		id, telegramID, username,
	).Scan(&ie); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while check exists user except current", "err", err)
			return false, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to check exists user except current", "err", err)
		return false, fmt.Errorf("could not check exists user except current: %w", err)
	}

	return ie, nil
}
