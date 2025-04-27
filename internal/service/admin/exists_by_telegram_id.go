package admin

import "context"

func (s *serv) ExistsByTelegramID(ctx context.Context, telegramID string) (bool, error) {
	s.logger.Debug("[check admin exists by telegram id] execute service")

	return s.adminRepository.ExistsByTelegramID(ctx, telegramID)
}
