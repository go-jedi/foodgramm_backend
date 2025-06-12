package adminguard

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgramm_backend/internal/service"
	"github.com/go-jedi/foodgramm_backend/pkg/jwt"
)

const (
	authorizationHeader = "Authorization"
	authorizationType   = "Bearer"
)

var (
	ErrEmptyAuthorizationHeader   = errors.New("empty authorization header")
	ErrInvalidAuthorizationHeader = errors.New("invalid authorization header")
	ErrTokenIsEmpty               = errors.New("token is empty")
	ErrAccessDenied               = errors.New("access denied: you do not have permission to perform this action")
)

type Middleware struct {
	adminService service.AdminService
	jwt          *jwt.JWT
}

func New(adminService service.AdminService, jwt *jwt.JWT) *Middleware {
	return &Middleware{
		adminService: adminService,
		jwt:          jwt,
	}
}

func (m *Middleware) AdminGuardMiddleware(c *gin.Context) {
	token, err := m.extractTokenFromHeader(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}

	vr, err := m.jwt.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}

	ie, err := m.adminService.ExistsByTelegramID(c.Request.Context(), vr.TelegramID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "internal server error",
		})
		return
	}

	if !ie {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status":  http.StatusForbidden,
			"message": ErrAccessDenied.Error(),
		})
		return
	}

	c.Next()
}

// extractTokenFromHeader extract token.
func (m *Middleware) extractTokenFromHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", ErrEmptyAuthorizationHeader
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != authorizationType {
		return "", ErrInvalidAuthorizationHeader
	}

	if len(headerParts[1]) == 0 {
		return "", ErrTokenIsEmpty
	}

	return headerParts[1], nil
}
