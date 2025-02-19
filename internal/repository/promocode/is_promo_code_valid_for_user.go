package promocode

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (r *repo) IsPromoCodeValidForUser(ctx context.Context, code string, telegramID string) (bool, error) {
	r.logger.Debug("[is promo code valid for user] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	iv := false

	q := `
		SELECT EXISTS(
			SELECT 1
			FROM promo_codes pc
			WHERE pc.code = $1
			AND pc.valid_from <= NOW() -- Промокод активен с указанной даты.
			AND (pc.valid_until IS NULL OR pc.valid_until >= NOW()) -- Промокод еще не истек.
			AND (pc.max_uses_allowed = -1 OR pc.amount_used < pc.max_uses_allowed) -- Лимит использований не исчерпан.
			AND (pc.is_reusable OR pc.amount_used = 0) -- Промокод многоразовый или еще не использован.
			AND NOT EXISTS ( -- Пользователь еще не использовал этот промокод.
				SELECT 1
				FROM promo_code_uses pcu
				WHERE pcu.promo_code_id = pc.id
				AND pcu.telegram_id = $2
			)
		);
	`

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		code, telegramID,
	).Scan(&iv); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while check promo code valid for user", "err", err)
			return false, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to check promo code valid for user", "err", err)
		return false, fmt.Errorf("could not check promo code valid for user: %w", err)
	}

	return iv, nil
}
