package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createList(c *gin.Context) {
	userID, _ := c.Get(userCtx)
	//if err != nil{
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//	return
	//}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
		"name" : userID,
	})
	//userID, err := getUserID(c)
	//if err != nil{
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//var input
}
