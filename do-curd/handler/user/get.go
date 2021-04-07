package user

import (
	. "do-curd/handler"
	"do-curd/model"
	"do-curd/pkg/errno"

	"github.com/gin-gonic/gin"
)

// 通过用户名获取用户信息
func Get(c *gin.Context) {
	// 根据 URL 路径解析出 username 的值
	username := c.Param("username")
	// 调用 model.GetUser() 函数查询该用户的数据库记录并返回
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, user)
}
