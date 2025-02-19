package promocode

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/promocode"
)

func (r *repo) Apply(ctx context.Context, dto promocode.ApplyDTO) (promocode.ApplyResponse, error) {
	r.logger.Debug("[apply promo code] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `SELECT * FROM public.promo_code_apply($1, $2);`

	var ar promocode.ApplyResponse

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		dto.Code, dto.TelegramID,
	).Scan(&ar); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while apply the promo code", "err", err)
			return promocode.ApplyResponse{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to apply promo code", "err", err)
		return promocode.ApplyResponse{}, fmt.Errorf("could not apply promo code: %w", err)
	}

	return ar, nil
}
