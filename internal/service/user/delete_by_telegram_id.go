package user

import "context"

func (s *serv) DeleteByTelegramID(ctx context.Context, telegramID string) (string, error) {
	return s.userRepository.DeleteByTelegramID(ctx, telegramID)
}
