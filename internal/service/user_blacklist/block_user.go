package userblacklist

import (
	"context"

	userblacklist "github.com/go-jedi/foodgramm_backend/internal/domain/user_blacklist"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

func (s *serv) BlockUser(ctx context.Context, dto userblacklist.BlockUserDTO, bannedByTelegramID string) (userblacklist.UsersBlackList, error) {
	s.logger.Debug("[block user] execute service")

	ie, err := s.Exists(ctx, dto.TelegramID)
	if err != nil {
		return userblacklist.UsersBlackList{}, err
	}

	if ie {
		return userblacklist.UsersBlackList{}, apperrors.ErrUserInBlackListAlreadyExists
	}

	return s.blacklistUserService.BlockUser(ctx, dto, bannedByTelegramID)
}
