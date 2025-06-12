package dependencies

import (
	"github.com/go-jedi/foodgramm_backend/internal/adapters/http/handlers/admin"
	"github.com/go-jedi/foodgramm_backend/internal/repository"
	adminRepository "github.com/go-jedi/foodgramm_backend/internal/repository/admin"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	adminService "github.com/go-jedi/foodgramm_backend/internal/service/admin"
)

func (d *Dependencies) AdminRepository() repository.AdminRepository {
	if d.adminRepository == nil {
		d.adminRepository = adminRepository.NewRepository(
			d.logger,
			d.db,
		)
	}

	return d.adminRepository
}

func (d *Dependencies) AdminService() service.AdminService {
	if d.adminService == nil {
		d.adminService = adminService.NewService(
			d.AdminRepository(),
			d.logger,
			d.cache,
		)
	}

	return d.adminService
}

func (d *Dependencies) AdminHandler() *admin.Handler {
	if d.adminHandler == nil {
		d.adminHandler = admin.NewHandler(
			d.AdminService(),
			d.cookie,
			d.middleware,
			d.engine,
			d.logger,
			d.validator,
		)
	}

	return d.adminHandler
}
