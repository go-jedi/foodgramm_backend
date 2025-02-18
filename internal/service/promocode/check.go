package promocode

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/promocode"
)

func (s *serv) Check(ctx context.Context, dto promocode.CheckDTO) error {
	s.logger.Debug("[check exists promo code] execute service")

	return s.promoCodeRepository.Check(ctx, dto)
}
