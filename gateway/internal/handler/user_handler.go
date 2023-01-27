package handler

import (
	"grpc-finance-app/gateway/internal/dto"
	authpb "grpc-finance-app/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(ctx *gin.Context) {
	var authReq dto.AuthReq
	if err := ctx.ShouldBindJSON(&authReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"messsage": "validate error", "error": err.Error()})
		return
	}

	user, err := h.authServiceClient.Register(ctx, &authpb.UserRegisterReq{
		Username: authReq.Username,
		Password: authReq.Password,
		Email:    authReq.Email,
	})
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": user})

}