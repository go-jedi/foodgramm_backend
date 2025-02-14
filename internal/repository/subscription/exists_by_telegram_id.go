package subscription

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (r *repo) ExistsByTelegramID(ctx context.Context, telegramID string) (bool, error) {
	r.logger.Debug("[ExistsByTelegramID] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	ie := false

	q := `SELECT * FROM public.subscription_exists($1);`

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
