package recipe

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-jedi/foodgrammm-backend/internal/domain/recipe"
	"github.com/go-jedi/foodgrammm-backend/pkg/apperrors"
)

// @Summary Save Recipe by Telegram ID
// @Description Save a recipe associated with a Telegram user
// @Tags Recipe
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token" default(Bearer <token>)
// @Param telegramID path string true "Telegram ID of the user"
// @Success 201 {boolean} bool "Recipe saved successfully"
// @Failure 400 {object} recipe.ErrorResponse "Invalid request"
// @Failure 404 {object} recipe.ErrorResponse "User does not exist"
// @Failure 500 {object} recipe.ErrorResponse "Internal server error"
// @Router /v1/recipe/save/telegram/{telegramID} [get]
func (h *Handler) saveRecipeByTelegramID(c *gin.Context) {
	telegramID := c.Param("telegramID")
	if telegramID == "" {
		h.logger.Error("failed to get param telegramID", "error", apperrors.ErrParamIsRequired)
		c.JSON(http.StatusBadRequest, recipe.ErrorResponse{
			Error:  "failed to get param telegramID",
			Detail: apperrors.ErrParamIsRequired.Error(),
		})
		return
	}

	res, err := h.recipeService.SaveRecipeByTelegramID(c.Request.Context(), telegramID)
	if err != nil {
		if errors.Is(err, apperrors.ErrUserDoesNotExist) {
			c.JSON(http.StatusNotFound, recipe.ErrorResponse{
				Error:  "user does not exists",
				Detail: err.Error(),
			})
			return
		}

		h.logger.Error("failed to save recipe by telegram id", "error", err)
		c.JSON(http.StatusInternalServerError, recipe.ErrorResponse{
			Error:  "failed to save recipe by telegram id",
			Detail: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}
