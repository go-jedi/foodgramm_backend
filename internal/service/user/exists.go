package user

import "context"

func (s *serv) Exists(ctx context.Context, telegramID string, username string) (bool, error) {
	s.logger.Debug("[Exists] execute service")

	return s.userRepository.Exists(ctx, telegramID, username)
}
