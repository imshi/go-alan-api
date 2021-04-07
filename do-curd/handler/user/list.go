package user

import (
	. "do-curd/handler"
	"do-curd/pkg/errno"
	"do-curd/service"

	"github.com/gin-gonic/gin"
)

// 从数据库中获取用户列表
func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	// 分页查询，收到的请求中传入的 offset 和 limit 参数，分别对应于 MySQL 的 offset 和 limit
	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
