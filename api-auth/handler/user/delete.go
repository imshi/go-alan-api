package user

import (
	. "gin-middleware/handler"
	"gin-middleware/model"
	"gin-middleware/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 通过用户 id 删除用户
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}