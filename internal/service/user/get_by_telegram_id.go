package user

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

func (s *serv) GetByTelegramID(ctx context.Context, telegramID string) (user.User, error) {
	return s.userRepository.GetByTelegramID(ctx, telegramID)
}
