package user

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) Update(ctx context.Context, dto user.UpdateDTO) (user.User, error) {
	ie, err := s.userRepository.ExistsExceptCurrent(ctx, dto.ID, dto.TelegramID, dto.Username)
	if err != nil {
		return user.User{}, err
	}

	if ie {
		return user.User{}, apperrors.ErrUserAlreadyExists
	}

	return s.userRepository.Update(ctx, dto)
}
