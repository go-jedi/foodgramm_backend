package auth

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/auth"
)

func (s *serv) Check(ctx context.Context, dto auth.CheckDTO) (auth.CheckResponse, error) {
	u, err := s.userRepository.GetByTelegramID(ctx, dto.TelegramID)
	if err != nil {
		return auth.CheckResponse{}, err
	}

	// check verify token
	vr, err := s.jwt.Verify(u.TelegramID, dto.Token)
	if err != nil {
		return auth.CheckResponse{}, err
	}

	return auth.CheckResponse{
		TelegramID: vr.TelegramID,
		Token:      dto.Token,
		ExpAt:      vr.ExpAt,
	}, nil
}
