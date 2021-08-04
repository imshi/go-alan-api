package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查header，如果X-Request-Id存在的话直接使用
		requestId := c.Request.Header.Get("X-Request-Id")

		// 使用 UUID4 创建 request id
		if requestId == "" {
			u4, _ := uuid.NewV4()
			requestId = u4.String()
		}

		// 在应用中使用生效
		c.Set("X-Request-Id", requestId)

		// 设置 header头信息
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()

	}
}
