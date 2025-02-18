package promocode

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/promocode"
)

func (s *serv) Create(ctx context.Context, dto promocode.CreateDTO) (promocode.PromoCode, error) {
	s.logger.Debug("[create promo code] execute service")

	return s.promoCodeRepository.Create(ctx, dto)
}
