package recipeevent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	recipeevent "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_event"
)

// @Summary Get all recipes event
// @Description Retrieve a list of all recipe's event.
// @Tags Recipe event
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Success 200 {array} recipeevent.Recipe
// @Failure 500 {object} recipeevent.ErrorResponse
// @Router /v1/recipe_event/all [get]
func (h *Handler) all(c *gin.Context) {
	h.logger.Debug("[get all recipe event] execute handler")

	result, err := h.recipeEventService.All(c.Request.Context())
	if err != nil {
		h.logger.Error("failed to get all recipes even", "error", err)
		c.JSON(http.StatusInternalServerError, recipeevent.ErrorResponse{
			Error:  "failed to get all recipes event",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
