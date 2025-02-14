package user

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/user"
)

func (s *serv) All(ctx context.Context) ([]user.User, error) {
	s.logger.Debug("[All] execute service")

	return s.userRepository.All(ctx)
}
