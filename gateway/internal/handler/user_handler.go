package handler

import (
	"grpc-finance-app/gateway/internal/dto"
	authpb "grpc-finance-app/proto"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		// fmt.Println("tesaja : ", err)
		// ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error", "err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": user})

}

func (h *Handler) Login(ctx *gin.Context) {
	var authLogindto dto.AuthLoginReq
	if err := ctx.ShouldBindJSON(&authLogindto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"messsage": "validate error", "error": err.Error()})
		return
	}

	user, err := h.authServiceClient.Login(ctx, &authpb.UserLoginReq{
		Username: authLogindto.Username,
		Password: authLogindto.Password,
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			switch e.Code() {
			case codes.Internal:
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "service internal error", "err": e.Message()})
				return
			case codes.NotFound:
				ctx.JSON(http.StatusBadRequest, gin.H{"message": "service data not found", "err": e.Message()})
				return
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "service unknown error", "err": e.Message()})
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in gateway (client rpc) ", "err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "login success", "data": user})

}
