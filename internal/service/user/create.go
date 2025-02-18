package user

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) Create(ctx context.Context, dto user.CreateDTO) (user.User, error) {
	s.logger.Debug("[create a new user] execute service")

	ie, err := s.userRepository.Exists(ctx, dto.TelegramID, dto.Username)
	if err != nil {
		return user.User{}, err
	}

	if ie {
		return user.User{}, apperrors.ErrUserAlreadyExists
	}

	return s.userRepository.Create(ctx, dto)
}
