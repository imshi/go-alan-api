// 自定义错误状态码

package errno

var (
	// 系统错误
	// 格式：错误类型（1）- 服务模块（00）- 具体的错误码（0x）
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// 用户操作错误
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}
)
