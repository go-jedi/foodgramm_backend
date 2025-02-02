package subscription

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (r *repo) ExistsByTelegramID(ctx context.Context, telegramID string) (bool, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	ie := false

	q := `
		SELECT is_active
		FROM subscriptions
		WHERE telegram_id = $1;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		telegramID,
	).Scan(&ie); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while check exists subscription by telegram id", "err", err)
			return false, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to check exists subscription by telegram id", "err", err)
		return false, fmt.Errorf("could not check exists subscription by telegram id: %w", err)
	}

	return ie, nil
}
