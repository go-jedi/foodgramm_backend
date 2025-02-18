package promocode

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/promocode"
)

func (s *serv) Apply(ctx context.Context, dto promocode.ApplyDTO) error {
	s.logger.Debug("[apply promo code] execute service")

	return s.promoCodeRepository.Apply(ctx, dto)
}
