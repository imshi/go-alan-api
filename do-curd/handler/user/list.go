package user

import (
	. "do-curd/handler"
	"do-curd/pkg/errno"
	"do-curd/service"

	"github.com/gin-gonic/gin"
)

// 从数据库中获取用户列表
func List(c *gin.Context) {
	var r ListResult
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	infos, count, err := service.ListUsers(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	SendResponse(c, nil, ListResponse{
		TocalCount: count,
		UserList:   infos,
	})
}
