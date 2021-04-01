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
