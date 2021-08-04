// 用户模块
package user

import (
	. "gin-middleware/handler"
	"gin-middleware/model"
	"gin-middleware/pkg/errno"
	"gin-middleware/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 用户模块业务处理函数 1：创建用户
func Create(c *gin.Context) {
	logrus.Info("User Create function called", util.GetReqID(c))
	var r CreateRequest
	// 将消息体 - body 作为指定的格式（这里是JSON）解析到 Go struct 变量（这里是 SendResponse）中
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}
	// 数据验证
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	//密码加密
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// 插入用户数据到数据库
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// 用户浏览器前端展示
	SendResponse(c, nil, rsp)
}
