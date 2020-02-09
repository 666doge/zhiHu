package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code int `json:"code"`
	Message string 	`json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func RespError(c *gin.Context, code int) {
	respData := &ResponseData{
		Code : code,
		Message: GetMessage(code),
	}
	c.JSON(http.StatusOK, respData)
}
func RespSuccess(c *gin.Context, data interface{}) {
	respData := &ResponseData{
		Code : ErrCodeSuccess,
		Message: GetMessage(ErrCodeSuccess),
		Data: data,
	}
	c.JSON(http.StatusOK, respData)
}