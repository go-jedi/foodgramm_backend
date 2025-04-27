package userblacklist

import "context"

func (s *serv) Exists(ctx context.Context, telegramID string) (bool, error) {
	s.logger.Debug("[check user exists in blacklist by telegram id] execute service")

	return s.blacklistUserService.Exists(ctx, telegramID)
}
