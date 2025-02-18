package promocode

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/promocode"
)

func (r *repo) Apply(_ context.Context, _ promocode.ApplyDTO) error {
	r.logger.Debug("[apply promo code] execute repository")

	return nil
}
