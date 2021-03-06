package handler

import (
	"backend_food_delivery/pkg/domain"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createList(c *gin.Context) {
	//userID, _ := c.Get(userCtx)
	////if err != nil{
	////	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	////	return
	////}
	//c.JSON(http.StatusOK, map[string]interface{}{
	//	"message": "ok",
	//	"name" : userID,
	//})
	//_, err := getUserID(c)
	var ctx context.Context
	//if err != nil{
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//	return
	//}
	var input domain.Todo
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.TodoList.Create(ctx, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) getAllLists(c *gin.Context) {
	var ctx context.Context
	lists, err := h.services.TodoList.GetAll(ctx)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, DataResponse{lists,
	})
}
