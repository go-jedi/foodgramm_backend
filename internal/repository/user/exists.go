package user

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (r *repo) Exists(ctx context.Context, telegramID string, username string) (bool, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	ie := false

	q := `
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE telegram_id = $1
			OR username = $2
		);
	`

	if err := r.db.Pool.QueryRow(ctxTimeout, q, telegramID, username).Scan(&ie); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while check exists user", "err", err)
			return false, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to check exists user", "err", err)
		return false, fmt.Errorf("could not check exists user: %w", err)
	}

	return ie, nil
}
