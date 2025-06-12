package userblacklist

import (
	"context"

	userblacklist "github.com/go-jedi/foodgramm_backend/internal/domain/user_blacklist"
)

func (s *serv) AllBannedUsers(ctx context.Context) ([]userblacklist.UsersBlackList, error) {
	s.logger.Debug("[get all banned users] execute service")

	return s.blacklistUserService.AllBannedUsers(ctx)
}
