package handler

import (
	"do-curd/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义返回格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 统一的返回函数 SendResponse 来格式化返回
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)
	// 固定返回格式，http返回状态码固定为200，响应信息中返回自定义错误码和错误信息
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
