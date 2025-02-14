package auth

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/auth"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) Check(ctx context.Context, dto auth.CheckDTO) (auth.CheckResponse, error) {
	s.logger.Debug("[Check] execute service")

	ie, err := s.userRepository.ExistsByTelegramID(ctx, dto.TelegramID)
	if err != nil {
		return auth.CheckResponse{}, err
	}

	if !ie {
		return auth.CheckResponse{}, apperrors.ErrUserDoesNotExist
	}

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
