package user

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

func (s *serv) List(ctx context.Context, dto user.ListDTO) (user.ListResponse, error) {
	s.logger.Debug("[List] execute service")

	return s.userRepository.List(ctx, dto)
}
