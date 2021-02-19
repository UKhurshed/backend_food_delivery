package handler

import (
	"backend_food_delivery/model"
	"github.com/gin-gonic/gin"
	"net/http"
)


func (h *Handler) signUp(ctx *gin.Context)   {
	var input model.User

	if err := ctx.BindJSON(&input); err != nil{
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	h
}

func (h *Handler) signIn(ctx *gin.Context)   {

}