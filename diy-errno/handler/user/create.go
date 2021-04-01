// 用户模块
package user

import (
	"diy-errno/pkg/errno"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 用户模块业务处理函数 1：创建用户
func Create(c *gin.Context) {
	var r struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error
	// 使用 Bind 函数对 gin 的 Context 对象传参
	if err = c.Bind(&r); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}

	logrus.Debugf("username is :[%s],password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		// 这里前后台日志不一样：后台日志中会输出敏感信息 username can not found in db: x.x.x.x，但是返回给用户的 message （{"code":20102,"message":"The user was not found. This is add message."}）不包含这些敏感信息，可以供前端直接对外展示
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in dh: x.x.x.x")).Add("This is add message.")
		logrus.Errorf("Get an error：%s", err)
	}

	if errno.IsErrUserNotFound(err) {
		logrus.Debug("err type is ErrUserNotFound")
	}

	if r.Password == "" {
		err = fmt.Errorf("password is empty")
	}

	code, message := errno.DecodeErr(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
}
