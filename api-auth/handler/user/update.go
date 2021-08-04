package user

import (
	. "api-auth/handler"
	"api-auth/model"
	"api-auth/pkg/errno"
	"api-auth/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 更新一个已存在的用户信息
func Update(c *gin.Context) {
	logrus.Info("Update function called.", util.GetReqID(c))
	// 从请求 url 中获取用户 id
	userId, _ := strconv.Atoi(c.Param("id"))

	// 绑定用户数据
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 通过用户id更新用户信息
	u.Id = uint64(userId)

	// 验证数据
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// 密码加密
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// 变动保存
	if err := u.Update(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
