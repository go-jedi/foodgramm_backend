package user

import "context"

func (s *serv) ExistsByTelegramID(ctx context.Context, telegramID string) (bool, error) {
	s.logger.Debug("[check user exists by telegram id] execute service")

	return s.userRepository.ExistsByTelegramID(ctx, telegramID)
}
