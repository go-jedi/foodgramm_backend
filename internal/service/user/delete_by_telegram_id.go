package user

import "context"

func (s *serv) DeleteByTelegramID(ctx context.Context, telegramID string) (string, error) {
	s.logger.Debug("[delete by telegram id] execute service")

	return s.userRepository.DeleteByTelegramID(ctx, telegramID)
}
