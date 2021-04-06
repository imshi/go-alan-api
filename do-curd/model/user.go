// 数据库curd操作及参数校验
package model

import (
	"do-curd/pkg/auth"
	"do-curd/pkg/constvar"
	"fmt"

	validator "gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

func (c *UserModel) TableName() string {
	return "tb_users"
}

// 创建一个新用户
func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

// 通过用户标识(id)删除用户
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.BaseModel.Id = id
	return DB.Self.Delete(&user).Error
}

// 更新用户信息
func (u *UserModel) Update() error {
	return DB.Self.Save(u).Error
}

// 通过用户标识(username)获取用户账户信息
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Self.Where("username = ?", username).First(&u)
	return u, d.Error
}

// 列出所有用户（gorm v2中的计数函数 Count()只支持 *int64 作为参数）
func ListUser(username string, offset, limit int) ([]*UserModel, int64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count int64

	where := fmt.Sprintf("username like '%%%s%%'", username)
	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Self.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}

// 与纯文本密码进行比较。如果与加密的相同，则返回true（在“User”结构中）
func (u *UserModel) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// 密码加密
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// 字段验证（）
func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
