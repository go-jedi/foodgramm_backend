package subscription

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) Create(ctx context.Context, telegramID string) error {
	s.logger.Debug("[Create] execute service")

	ie, err := s.userRepository.ExistsByTelegramID(ctx, telegramID)
	if err != nil {
		return err
	}

	if !ie {
		return apperrors.ErrUserDoesNotExist
	}

	return s.subscriptionRepository.Create(ctx, telegramID)
}
