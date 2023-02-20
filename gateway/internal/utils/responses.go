package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, err error) {
	resCode := http.StatusInternalServerError
	var ew *CustomErrorWrapper

	if errors.As(err, &ew) {
		switch ew.Code {
		case CodeClientError:
			resCode = http.StatusBadRequest
		case CodeNotFoundError:
			resCode = http.StatusNotFound
		case CodeConflictError:
			resCode = http.StatusConflict
		}
	}
	c.JSON(resCode, gin.H{"message": err.Error()})
}
