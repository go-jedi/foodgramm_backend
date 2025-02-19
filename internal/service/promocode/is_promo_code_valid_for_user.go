package promocode

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/promocode"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) IsPromoCodeValidForUser(ctx context.Context, dto promocode.IsPromoCodeValidForUserDTO) (bool, error) {
	s.logger.Debug("[is promo code valid for user] execute service")

	// check user exists by telegram id.
	ie, err := s.userRepository.ExistsByTelegramID(ctx, dto.TelegramID)
	if err != nil {
		return false, err
	}

	if !ie {
		return false, apperrors.ErrUserDoesNotExist
	}

	return s.promoCodeRepository.IsPromoCodeValidForUser(ctx, dto.Code, dto.TelegramID)
}
