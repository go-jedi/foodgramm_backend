package promocode

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (r *repo) Exists(ctx context.Context, code string) (bool, error) {
	r.logger.Debug("[check a promo code exists] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	ie := false

	q := `
		SELECT EXISTS(
			SELECT 1
			FROM promo_codes
			WHERE code = $1
		);
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		code,
	).Scan(&ie); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while check exists promo code", "err", err)
			return false, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to check exists promo code", "err", err)
		return false, fmt.Errorf("could not check exists promo code: %w", err)
	}

	return ie, nil
}
