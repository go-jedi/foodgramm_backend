package auth

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/auth"
)

func (s *serv) Refresh(ctx context.Context, dto auth.RefreshDTO) (auth.RefreshResponse, error) {
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
