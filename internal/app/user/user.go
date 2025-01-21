package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/pkg/logger"
	"github.com/go-jedi/foodgrammm-backend/pkg/postgres"
	"github.com/go-jedi/foodgrammm-backend/pkg/validator"
)

type User struct {
	engine    *gin.Engine
	logger    *logger.Logger
	validator *validator.Validator
	db        *postgres.Postgres
}
