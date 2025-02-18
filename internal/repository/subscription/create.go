package subscription

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (r *repo) Create(ctx context.Context, telegramID string) error {
	r.logger.Debug("[create subscription] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `SELECT * FROM public.subscription_create($1);`

	commandTag, err := r.db.Pool.Exec(
		ctxTimeout, q,
		telegramID,
	)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while create subscription for user", "err", err)
			return fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to create subscription for user", "err", err)
		return fmt.Errorf("could not create subscription for user: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return apperrors.ErrNoRowsWereAffected
	}

	return nil
}
