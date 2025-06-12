package user

import (
	"context"

	"github.com/go-jedi/foodgramm_backend/internal/domain/user"
	"github.com/go-jedi/foodgramm_backend/pkg/apperrors"
)

func (s *serv) Update(ctx context.Context, dto user.UpdateDTO) (user.User, error) {
	s.logger.Debug("[update user] execute service")

	ie, err := s.userRepository.ExistsExceptCurrent(ctx, dto.ID, dto.TelegramID, dto.Username)
	if err != nil {
		return user.User{}, err
	}

	if ie {
		return user.User{}, apperrors.ErrUserAlreadyExists
	}

	return s.userRepository.Update(ctx, dto)
}
