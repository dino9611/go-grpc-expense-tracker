package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorMiddleware(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) == 0 {
		return
	}
	firstErr := ctx.Errors[0].Err
	e, ok := status.FromError(firstErr)
	if ok { // if error from grpc
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

	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error in gateway (client rpc) ", "err": firstErr.Error()})
}
