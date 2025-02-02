package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/pkg/jwt"
)

const (
	authorizationHeader = "Authorization"
	authorizationType   = "Bearer"
	telegramIDCtx       = "telegramID"
)

var (
	ErrEmptyAuthorizationHeader   = errors.New("empty authorization header")
	ErrInvalidAuthorizationHeader = errors.New("invalid authorization header")
	ErrTokenIsEmpty               = errors.New("token is empty")
)

type AuthMiddleware struct {
	jwt *jwt.JWT
}

func NewAuthMiddleware(jwt *jwt.JWT) *AuthMiddleware {
	return &AuthMiddleware{
		jwt: jwt,
	}
}

// AuthMiddleware check authenticate user.
func (am *AuthMiddleware) AuthMiddleware(c *gin.Context) {
	token, err := am.extractTokenFromHeader(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}

	vr, err := am.jwt.ParseToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": err,
		})
		return
	}

	c.Set(telegramIDCtx, vr.TelegramID)
}

// GetTelegramIDFromContext get telegram id from context.
func (am *AuthMiddleware) GetTelegramIDFromContext(c *gin.Context) (string, bool) {
	if value, exists := c.Get(telegramIDCtx); exists {
		if telegramID, ok := value.(string); ok {
			return telegramID, true
		}
	}
	return "", false
}

// extractTokenFromHeader extract token.
func (am *AuthMiddleware) extractTokenFromHeader(c *gin.Context) (string, error) {
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
