// 用户模块
package user

import (
	. "do-curd/handler"
	"do-curd/model"
	"do-curd/pkg/errno"
	"do-curd/util"
	"fmt"

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
	admin2 := c.Param("username")
	logrus.Infof("URL username: %s", admin2)
	if r.Username == "" {
		// 前后台日志不一样：后台日志中会输出敏感信息 username can not found in db: x.x.x.x，但是返回给用户的 message （{"code":20102,"message":"The user was not found. This is add message."}）不包含这些敏感信息，可以供前端直接对外展示
		// 封装了自己的返回函数，通过统一的返回函数 SendResponse 来格式化返回（这里控制的是后台日志输出）
		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: x.x.x.x")), nil)
		return
	}

	if r.Password == "" {
		SendResponse(c, fmt.Errorf("password is empty"), nil)
	}

	rsp := CreateResponse{
		Username: r.Username,
	}
	// 用户浏览器前端展示
	SendResponse(c, nil, rsp)
}
