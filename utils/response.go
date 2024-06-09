package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message    *string `json:"message,omitempty"`
	Error      *string `json:"error,omitempty"`
	StatusCode int     `json:"statusCode"`
	Data       any     `json:"data,omitempty"`
}

func WriteError(ctx *gin.Context, err error) {
	if cErr, ok := err.(*CustomError); ok {
		ctx.JSON(cErr.StatusCode, Response{Error: &cErr.Message, StatusCode: cErr.StatusCode})
		return
	}
	errstr := err.Error()
	ctx.JSON(http.StatusInternalServerError, Response{StatusCode: http.StatusInternalServerError, Error: &errstr})
}

func WriteResponse(ctx *gin.Context, res any) {
	if msg, ok := res.(string); ok {
		ctx.JSON(http.StatusOK, Response{Message: &msg, StatusCode: http.StatusOK})
		return
	}

	if data, ok := res.(interface{}); ok {
		ctx.JSON(http.StatusOK, Response{Data: data, StatusCode: http.StatusOK})
		return
	}

	// Default response for unknown types
	ctx.JSON(http.StatusOK, Response{Data: res, StatusCode: http.StatusOK})
}
