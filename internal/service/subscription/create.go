package subscription

import (
	"context"

	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

func (s *serv) Create(ctx context.Context, telegramID string) error {
	s.logger.Debug("[create subscription] execute service")

	ie, err := s.userRepository.ExistsByTelegramID(ctx, telegramID)
	if err != nil {
		return err
	}

	if !ie {
		return apperrors.ErrUserDoesNotExist
	}

	return s.subscriptionRepository.Create(ctx, telegramID)
}
