package recipe

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary      Add free recipes count by Telegram ID
// @Description  This endpoint increments the count of free recipes available for a user identified by their Telegram ID.
// @Tags         Recipe
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Authorization token" default(Bearer <token>)
// @Param        telegramID path string true "Telegram ID of the user"
// @Success      200 {object} UserFreeRecipes "Successfully incremented free recipes count"
// @Failure      400 {object} recipe.ErrorResponse "Bad Request"
// @Failure      500 {object} recipe.ErrorResponse "Internal Server Error"
// @Router       /v1/recipe/free/telegram/{telegramID} [post]
func (h *Handler) addFreeRecipesCountByTelegramID(c *gin.Context) {
	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, recipe.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	result, err := h.recipeService.AddFreeRecipesCountByTelegramID(c.Request.Context(), telegramID)
	if err != nil {
		h.logger.Error("failed to add free recipes by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, recipe.ErrorResponse{
			Error:  "failed to add free recipes by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
