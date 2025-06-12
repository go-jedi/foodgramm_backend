package admin

import (
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
	"github.com/go-jedi/foodgramm_backend/pkg/redis"
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
