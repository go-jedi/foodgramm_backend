package userblacklist

import (
	"context"

	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

func (s *serv) UnblockUser(ctx context.Context, telegramID string) (string, error) {
	s.logger.Debug("[unblock user] execute service")

	ie, err := s.Exists(ctx, telegramID)
	if err != nil {
		return "", err
	}

	if !ie {
		return "", apperrors.ErrUserInBlackListDoesNotExist
	}

	return s.blacklistUserService.UnblockUser(ctx, telegramID)
}
