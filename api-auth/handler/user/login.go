package user

import (
	. "api-auth/handler"
	"api-auth/model"
	"api-auth/pkg/auth"
	"api-auth/pkg/errno"
	"api-auth/pkg/token"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	// 绑定数据模型
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// 通过登录用户名获取用户信息
	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// 登录密码验证
	if err = auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})

}
