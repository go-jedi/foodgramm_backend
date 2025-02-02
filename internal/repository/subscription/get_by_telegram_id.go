package subscription

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/subscription"
)

func (r *repo) GetByTelegramID(ctx context.Context, telegramID string) (subscription.Subscription, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	var s subscription.Subscription

	q := `
		SELECT *
		FROM subscriptions
		WHERE telegram_id = $1;
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		telegramID,
	).Scan(
		&s.ID, &s.TelegramID, &s.SubscribedAt, &s.ExpiresAt,
		&s.IsActive, &s.CreatedAt, &s.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while get subscription by telegram id", "err", err)
			return subscription.Subscription{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to get subscription by telegram id", "err", err)
		return subscription.Subscription{}, fmt.Errorf("could not get subscription by telegram id: %w", err)
	}

	return s, nil
}
