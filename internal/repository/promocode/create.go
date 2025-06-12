package promocode

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-jedi/foodgramm_backend/internal/domain/promocode"
	jsoniter "github.com/json-iterator/go"
)

func (r *repo) Create(ctx context.Context, dto promocode.CreateDTO) (promocode.PromoCode, error) {
	r.logger.Debug("[create promo code] execute repository")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Duration(r.db.QueryTimeout)*time.Second)
	defer cancel()

	q := `SELECT * FROM public.promo_code_create($1);`

	rawData, err := jsoniter.Marshal(dto)
	if err != nil {
		return promocode.PromoCode{}, err
	}

	var npc promocode.PromoCode

	if err := r.db.Pool.QueryRow(
		ctxTimeout, q,
		rawData,
	).Scan(
		&npc.ID, &npc.Code, &npc.DiscountPercent,
		&npc.MaxUsesAllowed, &npc.AmountUsed, &npc.IsReusable,
		&npc.ValidFrom, &npc.ValidUntil, &npc.CreatedAt, &npc.UpdatedAt,
	); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			r.logger.Error("request timed out while creating the promo code", "err", err)
			return promocode.PromoCode{}, fmt.Errorf("the request timed out: %w", err)
		}
		r.logger.Error("failed to create promo code", "err", err)
		return promocode.PromoCode{}, fmt.Errorf("could not create promo code: %w", err)
	}

	return npc, nil
}
