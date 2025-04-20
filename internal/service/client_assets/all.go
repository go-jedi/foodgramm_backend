package clientassets

import (
	"context"

	clientassets "github.com/go-jedi/foodgrammm-backend/internal/domain/client_assets"
)

func (s *serv) All(ctx context.Context) ([]clientassets.ClientAssets, error) {
	s.logger.Debug("[get all client assets] execute service")

	return s.clientAssetsRepository.All(ctx)
}
