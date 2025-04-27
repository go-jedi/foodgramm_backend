package admin

import (
	"github.com/go-jedi/foodgrammm-backend/internal/repository"
	"github.com/go-jedi/foodgrammm-backend/internal/service"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/redis"
)

type serv struct {
	adminRepository repository.AdminRepository
	logger          *logger.Logger
	cache           *redis.Redis
}

func NewService(
	adminRepository repository.AdminRepository,
	logger *logger.Logger,
	cache *redis.Redis,
) service.AdminService {
	return &serv{
		adminRepository: adminRepository,
		logger:          logger,
		cache:           cache,
	}
}
