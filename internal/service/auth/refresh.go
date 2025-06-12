package auth

import (
	"context"

	"github.com/go-jedi/foodgramm_backend/internal/domain/auth"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

func (s *serv) Refresh(ctx context.Context, dto auth.RefreshDTO) (auth.RefreshResponse, error) {
	s.logger.Debug("[refresh user token] execute service")

	ie, err := s.userRepository.ExistsByTelegramID(ctx, dto.TelegramID)
	if err != nil {
		return auth.RefreshResponse{}, err
	}

	if !ie {
		return auth.RefreshResponse{}, apperrors.ErrUserDoesNotExist
	}

	u, err := s.userRepository.GetByTelegramID(ctx, dto.TelegramID)
	if err != nil {
		return auth.RefreshResponse{}, err
	}

	// check verify token
	vr, err := s.jwt.Verify(u.TelegramID, dto.RefreshToken)
	if err != nil {
		return auth.RefreshResponse{}, err
	}

	// generate access, refresh tokens
	tokens, err := s.jwt.Generate(vr.TelegramID)
	if err != nil {
		return auth.RefreshResponse{}, err
	}

	return auth.RefreshResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		AccessExpAt:  tokens.AccessExpAt,
		RefreshExpAt: tokens.RefreshExpAt,
	}, nil
}
