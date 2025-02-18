package recipeofdays

import (
	"net/http"

	"github.com/gin-gonic/gin"
	recipeofdays "github.com/go-jedi/foodgrammm-backend/internal/domain/recipe_of_days"
)

// @Summary Get Recipe of the Day
// @Description Retrieve the current recipe of the day.
// @Tags Recipe of days
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Success 200 {object} Recipe
// @Failure 500 {object} ErrorResponse
// @Router /v1/recipe_of_days [get]
func (h *Handler) get(c *gin.Context) {
	h.logger.Debug("[get recipe of the day] execute handler")

	result, err := h.recipeOfDaysService.Get(c.Request.Context())
	if err != nil {
		h.logger.Error("failed to get recipe of days", "error", err)
		c.JSON(http.StatusInternalServerError, recipeofdays.ErrorResponse{
			Error:  "failed to get recipe of days",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
