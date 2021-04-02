// 用户模块
package user

import (
	. "do-request/handler"
	"do-request/pkg/errno"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 用户模块业务处理函数 1：创建用户
func Create(c *gin.Context) {
	var r CreateRequest
	// 将消息体 - body 作为指定的格式（这里是JSON）解析到 Go struct 变量（这里是 SendResponse）中
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 展示如何通过 Bind()、Param()、Query() 和 GetHeader() 来获取相应的参数
	// c.Param()：获取 URL 的参数值
	admin2 := c.Param("username")
	logrus.Infof("URL username: %s", admin2)

	// c.Query()：获取 URL中的地址参数
	desc := c.Query("desc")
	logrus.Infof("URL key param desc :%s", desc)

	// c.GetHeader()：获取 HTTP 头信息
	contentType := c.GetHeader("Content-Type")
	logrus.Infof("Header Content-Type:%s", contentType)

	logrus.Debugf("username is :[%s],password is [%s]", r.Username, r.Password)
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
