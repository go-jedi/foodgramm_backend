package subscription

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/subscription"
)

func (r *repo) Create(ctx context.Context, telegramID string) (subscription.Subscription, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `SELECT * FROM public.subscription_create($1);`

	var ns subscription.Subscription

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		telegramID,
	).Scan(
		&ns.ID, &ns.TelegramID, &ns.SubscribedAt, &ns.ExpiresAt,
		&ns.IsActive, &ns.CreatedAt, &ns.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while create subscription for user", "err", err)
			return subscription.Subscription{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to create subscription for user", "err", err)
		return subscription.Subscription{}, fmt.Errorf("could not create subscription for user: %w", err)
	}

	return ns, nil
}
