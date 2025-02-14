package user

import "context"

func (s *serv) ExistsByTelegramID(ctx context.Context, telegramID string) (bool, error) {
	s.logger.Debug("[ExistsByTelegramID] execute service")

	return s.userRepository.ExistsByTelegramID(ctx, telegramID)
}
