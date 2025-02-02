package subscription

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/subscription"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) GetByTelegramID(ctx context.Context, telegramID string) (subscription.Subscription, error) {
	uie, err := s.userRepository.ExistsByTelegramID(ctx, telegramID)
	if err != nil {
		return subscription.Subscription{}, err
	}

	if !uie {
		return subscription.Subscription{}, apperrors.ErrUserDoesNotExist
	}

	return s.subscriptionRepository.GetByTelegramID(ctx, telegramID)
}
