package userblacklist

import (
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/pkg/logger"
	"github.com/go-jedi/foodgramm_backend/pkg/redis"
)

type serv struct {
	blacklistUserService repository.UserBlackListRepository
	logger               *logger.Logger
	cache                *redis.Redis
}

func NewService(
	blacklistUserService repository.UserBlackListRepository,
	logger *logger.Logger,
	cache *redis.Redis,
) service.UserBlackListService {
	return &serv{
		blacklistUserService: blacklistUserService,
		logger:               logger,
		cache:                cache,
	}
}
