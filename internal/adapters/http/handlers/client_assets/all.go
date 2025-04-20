package clientassets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	clientassets "github.com/go-jedi/foodgrammm-backend/internal/domain/client_assets"
)

// @Summary Get all client assets
// @Description Retrieve a list of all client assets from the system
// @Tags Client Assets
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Success 200 {array} clientassets.ClientAssets
// @Failure 500 {object} clientassets.ErrorResponse
// @Router /v1/client_assets/all [get]
func (h *Handler) all(c *gin.Context) {
	h.logger.Debug("[get all client assets] execute handler")

	result, err := h.clientAssetsService.All(c.Request.Context())
	if err != nil {
		h.logger.Error("failed to get all client assets", "error", err)
		c.JSON(http.StatusInternalServerError, clientassets.ErrorResponse{
			Error:  "failed to get all client assets",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
