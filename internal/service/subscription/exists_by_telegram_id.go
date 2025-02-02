package subscription

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) ExistsByTelegramID(ctx context.Context, telegramID string) (bool, error) {
	ie, err := s.userRepository.ExistsByTelegramID(ctx, telegramID)
	if err != nil {
		return false, err
	}

	if !ie {
		return false, apperrors.ErrUserDoesNotExist
	}

	return s.subscriptionRepository.ExistsByTelegramID(ctx, telegramID)
}
