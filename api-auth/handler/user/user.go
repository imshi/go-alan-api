package user

import "api-auth/model"

// 如果消息体有 JSON 参数需要传递，针对每一个 API 接口定义独立的 go struct 来接收，比如 CreateRequest 和 CreateResponse，并将这些结构体统一放在一个 Go 文件中，以方便后期维护和修改
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount int64             `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
