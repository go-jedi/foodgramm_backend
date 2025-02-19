package promocode

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/promocode"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) Apply(ctx context.Context, dto promocode.ApplyDTO) (promocode.ApplyResponse, error) {
	s.logger.Debug("[apply promo code] execute service")

	// check user exists by telegram id.
	ieu, err := s.userRepository.ExistsByTelegramID(ctx, dto.TelegramID)
	if err != nil {
		return promocode.ApplyResponse{}, err
	}

	if !ieu {
		return promocode.ApplyResponse{}, apperrors.ErrUserDoesNotExist
	}

	// check promo code valid for user.
	pcv, err := s.promoCodeRepository.IsPromoCodeValidForUser(ctx, dto.Code, dto.TelegramID)
	if err != nil {
		return promocode.ApplyResponse{}, err
	}

	if !pcv {
		return promocode.ApplyResponse{}, apperrors.ErrPromoCodeIsNotValidForUser
	}

	return s.promoCodeRepository.Apply(ctx, dto)
}
