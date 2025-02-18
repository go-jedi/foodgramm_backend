package promocode

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/promocode"
)

func (r *repo) Check(_ context.Context, _ promocode.CheckDTO) error {
	r.logger.Debug("[check exists promo code] execute repository")

	return nil
}
