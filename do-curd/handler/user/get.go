package user

import (
	. "do-curd/handler"
	"do-curd/model"
	"do-curd/pkg/errno"

	"github.com/gin-gonic/gin"
)

// 通过用户名获取用户信息
func Get(c *gin.Context) {
	username := c.Param("username")
	// 从数据库中根据 username 查询用户信息
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, user)
}
