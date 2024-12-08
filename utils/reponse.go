package utils

import "github.com/gin-gonic/gin"

type HTTPError struct {
	Success bool   `json:"success" default:"false" `
	Message string `json:"message,omitempty"`
	Error   string `json:"error"`
}

type APIResponse struct {
	Success bool        `json:"success" default:"true" `
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func RespondSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func RespondError(ctx *gin.Context, code int, message string, err error) {
	ctx.JSON(code, HTTPError{
		Success: false,
		Message: message,
		Error:   err.Error(),
	})
}
