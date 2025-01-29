package user

import "context"

func (s *serv) ExistsByTelegramID(ctx context.Context, telegramID string) (bool, error) {
	return s.userRepository.ExistsByTelegramID(ctx, telegramID)
}
