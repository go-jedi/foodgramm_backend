package userblacklist

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (r *repo) Exists(ctx context.Context, telegramID string) (bool, error) {
	r.logger.Debug("[check user exists in blacklist by telegram id] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	ie := false

	q := `
		SELECT EXISTS(
			SELECT 1
			FROM users_blacklist
			WHERE telegram_id = $1
		);
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		telegramID,
	).Scan(&ie); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while check exists in blacklist by telegram id", "err", err)
			return false, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to check exists in blacklist by telegram id", "err", err)
		return false, fmt.Errorf("could not check exists in blacklist by telegram id: %w", err)
	}

	return ie, nil
}
