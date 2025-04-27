package admin

import (
	"context"

	"github.com/go-jedi/foodgrammm-backend/internal/domain/admin"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

func (s *serv) AddAdminUser(ctx context.Context, telegramID string) (admin.Admin, error) {
	s.logger.Debug("[add a new admin user] execute service")

	ie, err := s.ExistsByTelegramID(ctx, telegramID)
	if err != nil {
		return admin.Admin{}, err
	}

	if ie {
		return admin.Admin{}, apperrors.ErrAdminAlreadyExists
	}

	return s.adminRepository.AddAdminUser(ctx, telegramID)
}
