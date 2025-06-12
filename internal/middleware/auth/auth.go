package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgramm_backend/pkg/jwt"
)

const (
	authorizationHeader = "Authorization"
	authorizationType   = "Bearer"
	telegramIDCtx       = "telegramID"
)

var (
	ErrEmptyAuthorizationHeader              = errors.New("empty authorization header")
	ErrInvalidAuthorizationHeader            = errors.New("invalid authorization header")
	ErrTokenIsEmpty                          = errors.New("token is empty")
	ErrTelegramIDMakingRequestNotFound       = errors.New("telegram id making request not found")
	ErrTelegramIDMakingRequestHasInvalidType = errors.New("telegram id making request has invalid type")
)

type Middleware struct {
	jwt *jwt.JWT
}

func New(jwt *jwt.JWT) *Middleware {
	return &Middleware{
		jwt: jwt,
	}
}

// AuthMiddleware check authenticate user.
func (m *Middleware) AuthMiddleware(c *gin.Context) {
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

	c.Set(telegramIDCtx, vr.TelegramID)
}

// GetTelegramIDFromContext get telegram id making request from context.
func (m *Middleware) GetTelegramIDFromContext(c *gin.Context) (string, error) {
	val, ok := c.Get(telegramIDCtx)
	if !ok {
		return "", ErrTelegramIDMakingRequestNotFound
	}

	telegramID, ok := val.(string)
	if !ok {
		return "", ErrTelegramIDMakingRequestHasInvalidType
	}

	return telegramID, nil
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
