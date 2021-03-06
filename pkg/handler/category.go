package handler

import (
	"backend_food_delivery/pkg/domain"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createCategory(c *gin.Context) {

	var ctx context.Context
	var categoryInput domain.Category
	if err := c.BindJSON(&categoryInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Category.CreateCategory(ctx, categoryInput)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
	})

}

func (h *Handler) GetAllCategory(c *gin.Context) {
	var ctx context.Context

	categoryLists, err := h.services.Category.GetAllCategory(ctx)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, DataResponse{categoryLists})
}
