package promocode

import (
	"context"

	"github.com/go-jedi/foodgramm_backend/internal/domain/promocode"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

func (s *serv) Create(ctx context.Context, dto promocode.CreateDTO) (promocode.PromoCode, error) {
	s.logger.Debug("[create promo code] execute service")

	ie, err := s.promoCodeRepository.Exists(ctx, dto.Code)
	if err != nil {
		return promocode.PromoCode{}, err
	}

	if ie {
		return promocode.PromoCode{}, apperrors.ErrPromoCodeAlreadyExists
	}

	return s.promoCodeRepository.Create(ctx, dto)
}
