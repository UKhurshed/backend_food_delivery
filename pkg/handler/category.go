package handler

import (
	"backend_food_delivery/pkg/domain"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (h *Handler) Delete(c *gin.Context) {

	var ctx context.Context
	id := c.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	err := h.services.Category.DeleteCategory(ctx, objectId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
	})
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	var ctx context.Context
	var categoryInput domain.Category
	if err := c.BindJSON(&categoryInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Category.UpdateCategory(ctx, categoryInput)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message":"category updated",
	})


}
