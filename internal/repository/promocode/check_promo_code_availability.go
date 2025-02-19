package promocode

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (r *repo) CheckPromoCodeAvailability(ctx context.Context, code string) (bool, error) {
	r.logger.Debug("[check promo code availability] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	pca := false

	q := `
		SELECT EXISTS(
			SELECT 1
			FROM promo_codes
			WHERE code = $1
			AND valid_from <= NOW() -- Промокод активен с указанной даты
			AND (valid_until IS NULL OR valid_until >= NOW()) -- Промокод еще не истек
			AND (max_uses_allowed = -1 OR amount_used < max_uses_allowed) -- Лимит использований не исчерпан
			AND (is_reusable OR amount_used = 0) -- Промокод многоразовый или еще не использован
		);
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		code,
	).Scan(&pca); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while check promo code availability", "err", err)
			return false, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to check promo code availability", "err", err)
		return false, fmt.Errorf("could not check promo code availability: %w", err)
	}

	return pca, nil
}
